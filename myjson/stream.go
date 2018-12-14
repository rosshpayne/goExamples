// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"bytes"
	"errors"
	"io"
)

										// A Decoder reads and decodes JSON values from an input stream.
type Decoder struct {
	r     io.Reader
	buf   []byte
	d     decodeState							// defined in decode.go  Holds json value to be decoded
	scanp int 								// start of unread data in buf
	scan  scanner
	err   error

	tokenState int
	tokenStack []int
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may
// read data from r beyond the JSON values requested.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// UseNumber causes the Decoder to unmarshal a number into an interface{} as a
// Number instead of as a float64.
func (dec *Decoder) UseNumber() { dec.d.useNumber = true }

										// Decode reads the next JSON-encoded value from its
										// input and stores it in the value pointed to by v.
										//
										// See the documentation for Unmarshal for details about
										// the conversion of JSON into a Go value.
func (dec *Decoder) Decode(v interface{}) error {
	if dec.err != nil {
		return dec.err
	}

	if err := dec.tokenPrepareForDecode(); err != nil {      		// move from : state to value state
		return err
	}

	if !dec.tokenValueAllowed() {
		return &SyntaxError{msg: "not at beginning of value"}
	}
										// Read whole value into buffer.
	n, err := dec.readValue()						// use stream to read json value in buf
	if err != nil {
		return err
	}
	dec.d.init(dec.buf[dec.scanp : dec.scanp+n])				// initialise decodeState using json data/value supplied by readValue()
	dec.scanp += n								//  now move ahead unread portion of buf
										// Don't save err from unmarshal into dec.err:
										// the connection is still usable since we read a complete JSON
										// object from it before the error happened.
	err = dec.d.unmarshal(v)

	dec.tokenValueEnd()

	return err
}

										// Buffered returns a reader of the data remaining in the Decoder's
										// buffer. The reader is valid until the next call to Decode.
func (dec *Decoder) Buffered() io.Reader {
	return bytes.NewReader(dec.buf[dec.scanp:])
}
										// readValue reads a JSON value from dec.buf. Value is a array element or value part of json object
func (dec *Decoder) readValue() (int, error) {					// It returns the length of the encoding.
	dec.scan.reset()							// use scanner in scanner.go.  It scans a json value char by char.

	scanp := dec.scanp							// func var 
	var err error
Input:										// breaks referencing this label will start at then of following for-loop.
	for {									// outer for handles reads from stream to complete value scanning
		for i, c := range dec.buf[scanp:] {				// get byte from next unread portion of buf
			dec.scan.bytes++
			v := dec.scan.step(&dec.scan, c)			// run scanner function on current byte
			if v == scanEnd {
				scanp += i					//   completed scan of json value. abort loop
				break Input					
			}
										// scanEnd is delayed one byte. We might block trying to get that byte from src
										//   so instead invent a space byte.
			if (v == scanEndObject || v == scanEndArray) && dec.scan.step(&dec.scan, ' ') == scanEnd {
				scanp += i + 1
				break Input
			}
			if v == scanError {
				dec.err = dec.scan.err
				return 0, dec.scan.err
			}
		}
		scanp = len(dec.buf)

		// Did the last read have an error?
		// Delayed until now to allow buffer scan.
		if err != nil {
			if err == io.EOF {
				if dec.scan.step(&dec.scan, ' ') == scanEnd {
					break Input
				}
				if nonSpace(dec.buf) {
					err = io.ErrUnexpectedEOF
				}
			}
			dec.err = err
			return 0, err
		}

		n := scanp - dec.scanp
		err = dec.refill()					// update buf with data from stream
		scanp = dec.scanp + n
	}
	return scanp - dec.scanp, nil					// break Input starts here...
}

									// slice =  pointer to array index of start of slice + # of elements + cap of array
									//  append will update elements starting at position start+len in the array. If necessary it will 
									//  allocate new array to extend its size to accomodate append data.
func (dec *Decoder) refill() error {
									// Make room to read more into the buffer.
									// First slide down data already consumed.
	if dec.scanp > 0 {
		n := copy(dec.buf, dec.buf[dec.scanp:])			// copy unread slice into start of dec.buf. NB copy does no allocation it modifies array contents. 
									//   does not change len of dest slice. This is done in next stmt.
		dec.buf = dec.buf[:n]					// adjust buf len to # of bytes copied
		dec.scanp = 0
	}
									// Grow buffer if not large enough.
	const minRead = 512
	if cap(dec.buf)-len(dec.buf) < minRead {
		newBuf := make([]byte, len(dec.buf), 2*cap(dec.buf)+minRead)	// allocate new array sized to provide more space for following Read.  newBuf slice point to it.
		copy(newBuf, dec.buf)					// copy buf slice into new slice	
		dec.buf = newBuf					// buf slice gets its len, cap & array from newBuf slice.  Old dec.duf is GC
	}
									// Read. Delay error for next iteration (after scan).
	n, err := dec.r.Read(dec.buf[len(dec.buf):cap(dec.buf)]) 	// pass a new slice to Read, starting at portion of array that has not been populated ie. len(dec.buf)
	dec.buf = dec.buf[0 : len(dec.buf)+n]				//   ending at the capacity of array.  Read will load n bytes into this portion of array
									//   after returning from Read we then need to adjust size of dec.buf slice to accomodate the n bytes.

	return err
}

func nonSpace(b []byte) bool {
	for _, c := range b {
		if !isSpace(c) {
			return true
		}
	}
	return false
}

// An Encoder writes JSON values to an output stream.
type Encoder struct {
	w          io.Writer
	err        error
	escapeHTML bool

	indentBuf    *bytes.Buffer
	indentPrefix string
	indentValue  string
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w, escapeHTML: true}
}

// Encode writes the JSON encoding of v to the stream,
// followed by a newline character.
//
// See the documentation for Marshal for details about the
// conversion of Go values to JSON.
func (enc *Encoder) Encode(v interface{}) error {
	if enc.err != nil {
		return enc.err
	}
	e := newEncodeState()
	err := e.marshal(v, encOpts{escapeHTML: enc.escapeHTML})
	if err != nil {
		return err
	}

	// Terminate each value with a newline.
	// This makes the output look a little nicer
	// when debugging, and some kind of space
	// is required if the encoded value was a number,
	// so that the reader knows there aren't more
	// digits coming.
	e.WriteByte('\n')

	b := e.Bytes()
	if enc.indentPrefix != "" || enc.indentValue != "" {
		if enc.indentBuf == nil {
			enc.indentBuf = new(bytes.Buffer)
		}
		enc.indentBuf.Reset()
		err = Indent(enc.indentBuf, b, enc.indentPrefix, enc.indentValue)
		if err != nil {
			return err
		}
		b = enc.indentBuf.Bytes()
	}
	if _, err = enc.w.Write(b); err != nil {
		enc.err = err
	}
	encodeStatePool.Put(e)
	return err
}

// SetIndent instructs the encoder to format each subsequent encoded
// value as if indented by the package-level function Indent(dst, src, prefix, indent).
// Calling SetIndent("", "") disables indentation.
func (enc *Encoder) SetIndent(prefix, indent string) {
	enc.indentPrefix = prefix
	enc.indentValue = indent
}

// SetEscapeHTML specifies whether problematic HTML characters
// should be escaped inside JSON quoted strings.
// The default behavior is to escape &, <, and > to \u0026, \u003c, and \u003e
// to avoid certain safety problems that can arise when embedding JSON in HTML.
//
// In non-HTML settings where the escaping interferes with the readability
// of the output, SetEscapeHTML(false) disables this behavior.
func (enc *Encoder) SetEscapeHTML(on bool) {
	enc.escapeHTML = on
}

// RawMessage is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type RawMessage []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m RawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

var _ Marshaler = (*RawMessage)(nil)
var _ Unmarshaler = (*RawMessage)(nil)

// A Token holds a value of one of these types:
//
//	Delim, for the four JSON delimiters [ ] { }
//	bool, for JSON booleans
//	float64, for JSON numbers
//	Number, for JSON numbers
//	string, for JSON string literals
//	nil, for JSON null
//
type Token interface{}

const (
	tokenTopValue = iota
	tokenArrayStart
	tokenArrayValue
	tokenArrayComma
	tokenObjectStart
	tokenObjectKey
	tokenObjectColon
	tokenObjectValue
	tokenObjectComma
)
											// advance tokenstate from a separator state to a value state
func (dec *Decoder) tokenPrepareForDecode() error {
											// Note: Not calling peek before switch, to avoid
											// putting peek into the standard Decode path.
											// peek is only called when using the Token API.
	switch dec.tokenState {
	case tokenArrayComma:								// anticipated next scannable char
		c, err := dec.peek()							// scan until first non-space byte & return it. Update buf from stream.
		if err != nil {
			return err
		}
		if c != ',' {
			return &SyntaxError{"expected comma after array element", 0}
		}
		dec.scanp++								// increment scan pointer (index into dec.buf
		dec.tokenState = tokenArrayValue					// anticipated next significant char
	case tokenObjectColon:
		c, err := dec.peek()
		if err != nil {
			return err
		}
		if c != ':' {
			return &SyntaxError{"expected colon after object key", 0}
		}
		dec.scanp++
		dec.tokenState = tokenObjectValue
	}
	return nil
}

func (dec *Decoder) tokenValueAllowed() bool {

	switch dec.tokenState {
	case tokenTopValue, tokenArrayStart, tokenArrayValue, tokenObjectValue:
		return true
	}
	return false
}

func (dec *Decoder) tokenValueEnd() {
	switch dec.tokenState {
	case tokenArrayStart, tokenArrayValue:
		dec.tokenState = tokenArrayComma
	case tokenObjectValue:
		dec.tokenState = tokenObjectComma
	}
}

// A Delim is a JSON array or object delimiter, one of [ ] { or }.
type Delim rune

func (d Delim) String() string {
	return string(d)
}

// Token returns the next JSON token in the input stream.
// At the end of the input stream, Token returns nil, io.EOF.
//
// Token guarantees that the delimiters [ ] { } it returns are
// properly nested and matched: if Token encounters an unexpected
// delimiter in the input, it will return an error.
//
// The input stream consists of basic JSON values—bool, string,
// number, and null—along with delimiters [ ] { } of type Delim
// to mark the start and end of arrays and objects.
// Commas and colons are elided.
func (dec *Decoder) Token() (Token, error) {
	for {
		c, err := dec.peek()
		if err != nil {
			return nil, err
		}
		switch c {
		case '[':
			if !dec.tokenValueAllowed() {
				return dec.tokenError(c)
			}
			dec.scanp++
			dec.tokenStack = append(dec.tokenStack, dec.tokenState)
			dec.tokenState = tokenArrayStart
			return Delim('['), nil

		case ']':
			if dec.tokenState != tokenArrayStart && dec.tokenState != tokenArrayComma {
				return dec.tokenError(c)
			}
			dec.scanp++
			dec.tokenState = dec.tokenStack[len(dec.tokenStack)-1]
			dec.tokenStack = dec.tokenStack[:len(dec.tokenStack)-1]
			dec.tokenValueEnd()
			return Delim(']'), nil

		case '{':
			if !dec.tokenValueAllowed() {
				return dec.tokenError(c)
			}
			dec.scanp++
			dec.tokenStack = append(dec.tokenStack, dec.tokenState)
			dec.tokenState = tokenObjectStart
			return Delim('{'), nil

		case '}':
			if dec.tokenState != tokenObjectStart && dec.tokenState != tokenObjectComma {
				return dec.tokenError(c)
			}
			dec.scanp++
			dec.tokenState = dec.tokenStack[len(dec.tokenStack)-1]
			dec.tokenStack = dec.tokenStack[:len(dec.tokenStack)-1]
			dec.tokenValueEnd()
			return Delim('}'), nil

		case ':':
			if dec.tokenState != tokenObjectColon {
				return dec.tokenError(c)
			}
			dec.scanp++
			dec.tokenState = tokenObjectValue
			continue

		case ',':
			if dec.tokenState == tokenArrayComma {
				dec.scanp++
				dec.tokenState = tokenArrayValue
				continue
			}
			if dec.tokenState == tokenObjectComma {
				dec.scanp++
				dec.tokenState = tokenObjectKey
				continue
			}
			return dec.tokenError(c)

		case '"':
			if dec.tokenState == tokenObjectStart || dec.tokenState == tokenObjectKey {
				var x string
				old := dec.tokenState
				dec.tokenState = tokenTopValue
				err := dec.Decode(&x)
				dec.tokenState = old
				if err != nil {
					clearOffset(err)
					return nil, err
				}
				dec.tokenState = tokenObjectColon
				return x, nil
			}
			fallthrough

		default:
			if !dec.tokenValueAllowed() {
				return dec.tokenError(c)
			}
			var x interface{}
			if err := dec.Decode(&x); err != nil {
				clearOffset(err)
				return nil, err
			}
			return x, nil
		}
	}
}

func clearOffset(err error) {
	if s, ok := err.(*SyntaxError); ok {
		s.Offset = 0
	}
}

func (dec *Decoder) tokenError(c byte) (Token, error) {
	var context string
	switch dec.tokenState {
	case tokenTopValue:
		context = " looking for beginning of value"
	case tokenArrayStart, tokenArrayValue, tokenObjectValue:
		context = " looking for beginning of value"
	case tokenArrayComma:
		context = " after array element"
	case tokenObjectKey:
		context = " looking for beginning of object key string"
	case tokenObjectColon:
		context = " after object key"
	case tokenObjectComma:
		context = " after object key:value pair"
	}
	return nil, &SyntaxError{"invalid character " + quoteChar(c) + " " + context, 0}
}

// More reports whether there is another element in the
// current array or object being parsed.
func (dec *Decoder) More() bool {
	c, err := dec.peek()
	return err == nil && c != ']' && c != '}'
}

func (dec *Decoder) peek() (byte, error) {				// peep will find next significant byte in dec.buf, ignoring spaces as it does.
	var err error
	for {
		for i := dec.scanp; i < len(dec.buf); i++ {
			c := dec.buf[i]
			if isSpace(c) {
				continue
			}
			dec.scanp = i
			return c, nil
		}
		// buffer has been scanned, now report any error
		if err != nil {
			return 0, err
		}
		err = dec.refill()
	}
}

/*
TODO

// EncodeToken writes the given JSON token to the stream.
// It returns an error if the delimiters [ ] { } are not properly used.
//
// EncodeToken does not call Flush, because usually it is part of
// a larger operation such as Encode, and those will call Flush when finished.
// Callers that create an Encoder and then invoke EncodeToken directly,
// without using Encode, need to call Flush when finished to ensure that
// the JSON is written to the underlying writer.
func (e *Encoder) EncodeToken(t Token) error  {
	...
}

*/
