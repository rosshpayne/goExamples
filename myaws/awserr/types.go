package awserr

import "fmt"

// SprintError returns a string of the formatted error code.
//
// Both extra and origErr are optional.  If they are included their lines
// will be added, but if they are not included their lines will be ignored.
func SprintError(code, message, extra string, origErr error) string {		// origErr is the receiver for the concrete type that satisfies error interface.
	msg := fmt.Sprintf("%s: %s", code, message)				//  typically what is passed in is slice of errors that has the Error() method.
	if extra != "" {
		msg = fmt.Sprintf("%s\n\t%s", msg, extra)
	}
	if origErr != nil {
		msg = fmt.Sprintf("%s\ncaused by: %s", msg, origErr.Error())
	}
	return msg
}

// A baseError wraps the code and message which defines an error. It also
// can be used to wrap an original error object.
//
// Should be used as the root for errors satisfying the awserr.Error. Also	// NB awserr.Error is only an interface. to satisfy it we must build concrete types with
// for any error which does not fit into a specific error wrapper type. 	//  with the methods that satisfy the interface. baseError is this concrete type.
type baseError struct {								// must satisfy the awserr.Error interface - see New(..) and  NewBatchError(..)
	// Classification of error						//   any concrete type of this type will have all the methods in awserr.Error interface.
	code string								//   these methods are:  Error(), Code(), Message(), OrigErr()
										// note: type is unexported but its methods are.
	// Detailed information about error
	message string

	// Optional original error this error is based off of. Allows building
	// chained errors.
	errs []error								// baseError is setup for batch of base errors in this slice.
}										//  note: errs concrete type may satisfy more than the error interface i.e. Error`:w

// newBaseError returns an error object for the code, message, and errors.
//
// code is a short no whitespace phrase depicting the classification of
// the error that is being created.
//
// message is the free flow string containing detailed information about the
// error.
//
// origErrs is the error objects which will be nested under the new errors to
// be returned.
func newBaseError(code, message string, origErrs []error) *baseError {
	b := &baseError{							// allocated from heap as b escapes func scope.
		code:    code,							//   b's methods are exported so are executable from outside pkg.
		message: message,						//   because b fields are unexported they can only be populated & queried via its methods.
		errs:    origErrs,
	}

	return b
}

// Error returns the string representation of the error.
//
// See ErrorWithExtra for formatting.
//
// Satisfies the error interface.					// required to satisfy awserr.Error interface
func (b baseError) Error() string {
	size := len(b.errs)
	if size > 0 {
		return SprintError(b.code, b.message, "", errorList(b.errs))  // errorList converts a concrete type of []error to []error type with Error() method.
	}								      //   SperintError expects the third argument to satisify the error interface.

	return SprintError(b.code, b.message, "", nil)
}

// String returns the string representation of the error.
// Alias for Error to satisfy the stringer interface.
func (b baseError) String() string {
	return b.Error()
}

// Code returns the short phrase depicting the classification of the error.
func (b baseError) Code() string {
	return b.code
}

// Message returns the error details message.
func (b baseError) Message() string {
	return b.message
}

// OrigErr returns the original error if one was set. Nil is returned if no
// error was set. This only returns the first element in the list. If the full
// list is needed, use BatchedErrors.
func (b baseError) OrigErr() error {
	switch len(b.errs) {
	case 0:
		return nil
	case 1:
		return b.errs[0]
	default:									// default runs for batch of errors i.e. > 1
		if err, ok := b.errs[0].(Error); ok {					// if IV's concrete type satisifies Error interface then were dealing with a baseError type
			return NewBatchError(err.Code(), err.Message(), b.errs[1:])	//  therefore we must return a baseError which satisfies BatchErrors interface type 
											//  which is a superset of error interface
		}
		return NewBatchError("BatchedErrors", "multiple errors occurred", b.errs) // original error is of type error and there are multiple of them ie. batched.
											  //  return as a baseError containing the [] of errors.
	}
}

// OrigErrs returns the original errors if one was set. An empty slice is
// returned if no error was set.
func (b baseError) OrigErrs() []error {
	return b.errs
}

// So that the Error interface type can be included as an anonymous field
// in the requestError struct and not conflict with the error.Error() method.
type awsError Error

// A requestError wraps a request or service error.
//
// Composed of baseError for code, message, and original error.
type requestError struct {				// contains fields that contain concrete types that must satisfy either an interface or are Go base types.
	awsError					// contains a IV. It's concrete type must satisify Error interface i.e. Error(), Code(), Message(), OrigErrs() methods
							//   because its anonymous embedded type,these methods are promoted to the instance of requestError so 
							//   requestError satisfies the golang error interface aswellas aws Error interface.
	statusCode int					//   we say requestError has an aws Error aswell as hidden requestID & statusCode accessible only through methods.
	requestID  string
}

// newRequestError returns a wrapped error with additional information for
// request status code, and service requestID.
//
// Should be used to wrap all request which involve service requests. Even if
// the request failed without a service response, but had an HTTP status code
// that may be meaningful.
//
// Also wraps original errors via the baseError.
func newRequestError(err Error, statusCode int, requestID string) *requestError {
	return &requestError{
		awsError:   err,						// IV, concrete type would typically be a *baseError type
		statusCode: statusCode,
		requestID:  requestID,
	}
}


// Error returns the string representation of the error.
// Satisfies the error interface.
func (r requestError) Error() string {						// this  overrides the promoted Error() method in awsError to take into account
	extra := fmt.Sprintf("status code: %d, request id: %s",			//   statusCode & RequestID
		r.statusCode, r.requestID)
	return SprintError(r.Code(), r.Message(), extra, r.OrigErr())		// print values from promoted methods of awsError IV 
}

// String returns the string representation of the error.
// Alias for Error to satisfy the stringer interface.
func (r requestError) String() string {
	return r.Error()
}

// StatusCode returns the wrapped status code for the error
func (r requestError) StatusCode() int {
	return r.statusCode
}

// RequestID returns the wrapped requestID
func (r requestError) RequestID() string {
	return r.requestID
}

// OrigErrs returns the original errors if one was set. An empty slice is
// returned if no error was set.
func (r requestError) OrigErrs() []error {					// enables support for BatchedErrors interface.
	if b, ok := r.awsError.(BatchedErrors); ok {				// if awsError concrete type satisfies BatchedErrors interface
		return b.OrigErrs()
	}
	return []error{r.OrigErr()}
}

// An error list that satisfies the golang interface				// used to convert []error to support golang error interface.
type errorList []error

// Error returns the string representation of the error.
//
// Satisfies the error interface.
func (e errorList) Error() string {
	msg := ""
	// How do we want to handle the array size being zero
	if size := len(e); size > 0 {
		for i := 0; i < size; i++ {
			msg += fmt.Sprintf("%s", e[i].Error())
			// We check the next index to see if it is within the slice.
			// If it is, then we append a newline. We do this, because unit tests
			// could be broken with the additional '\n'
			if i+1 < size {
				msg += "\n"
			}
		}
	}
	return msg
}
