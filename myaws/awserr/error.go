// Package awserr represents API error interface accessors for the SDK.
package awserr

// An Error wraps lower level errors with code, message and an original error.
// The underlying concrete error type may also satisfy other interfaces which
// can be to used to obtain more specific information about the error.
//
// Calling Error() or String() will always include the full information about
// an error based on its underlying type.
//
// Example:
//
//     output, err := s3manage.Upload(svc, input, opts)
//     if err != nil {
//         if awsErr, ok := err.(awserr.Error); ok {
//             // Get error details
//             log.Println("Error:", awsErr.Code(), awsErr.Message())
//
//             // Prints out full error message, including original error if there was one.
//             log.Println("Error:", awsErr.Error())
//
//             // Get original error
//             if origErr := awsErr.OrigErr(); origErr != nil {
//                 // operate on original error.
//             }
//         } else {
//             fmt.Println(err.Error())
//         }
//     }
//
type Error interface {								// remember interface doesn't have fields it only has a list of methods that must be satisified
	// Satisfy the generic error interface.					//   ..some methods maybe presented in terms of other interfaces as embedded error.
	error									// interface satisfies interface error: has one method, Error() 
										// the concrete type that is this type must satisfy these methods.
	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErr() error								//  Error supports a single error IV. Alternatively, Error is extended in BatchedErrors
}										//    to support []error ie. multiple error IVs.

// BatchError is a batch of errors which also wraps lower level errors with
// code, message, and original errors. Calling Error() will include all errors
// that occurred in the batch.
//
// Deprecated: Replaced with BatchedErrors. Only defined for backwards
// compatibility.
type BatchError interface {							// an interface says nothing about concrete types only that they must satisfy the following methods
	// Satisfy the generic error interface.					//   an interface does not contain fields only methods or interface types.
	error									// satisfy methods in this interface..

	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErrs() []error
}

// BatchedErrors is a batch of errors which also wraps lower level errors with
// code, message, and original errors. Calling Error() will include all errors
// that occurred in the batch.
//
// Replaces BatchError
type BatchedErrors interface {
	// Satisfy the base Error interface.
	Error					

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErrs() []error
}

// New returns an Error object described by the code, message, and origErr.
//
// If origErr satisfies the Error interface it will not be wrapped within a new
// Error object and will instead be returned.
func New(code, message string, origErr error) Error {				// constructor func for concrete types that satisfy Error interface.
	var errs []error
	if origErr != nil {
		errs = append(errs, origErr)					// turn origErr into slice errs.
	}
	return newBaseError(code, message, errs)				// returns a concrete type that satisifies Error
}

// NewBatchError returns an BatchedErrors with a collection of errors as an
// array of errors.
func NewBatchError(code, message string, errs []error) BatchedErrors {
	return newBaseError(code, message, errs)				// returns a concrete type that satisfies BatchedErrors
}

// A RequestFailure is an interface to extract request failure information from
// an Error such as the request ID of the failed request returned by a service.
// RequestFailures may not always have a requestID value if the request failed
// prior to reaching the service such as a connection error.
//
// Example:
//
//     output, err := s3manage.Upload(svc, input, opts)
//     if err != nil {
//         if reqerr, ok := err.(RequestFailure); ok {
//             log.Println("Request failed", reqerr.Code(), reqerr.Message(), reqerr.RequestID())
//         } else {
//             log.Println("Error:", err.Error())
//         }
//     }
//
// Combined with awserr.Error:
//
//    output, err := s3manage.Upload(svc, input, opts)
//    if err != nil {
//        if awsErr, ok := err.(awserr.Error); ok {
//            // Generic AWS Error with Code, Message, and original error (if any)
//            fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
//
//            if reqErr, ok := err.(awserr.RequestFailure); ok {
//                // A service error occurred
//                fmt.Println(reqErr.StatusCode(), reqErr.RequestID())
//            }
//        } else {
//            fmt.Println(err.Error())
//        }
//    }
//
type RequestFailure interface {
	Error

	// The status code of the HTTP response.
	StatusCode() int

	// The request ID returned by the service for a request failure. This will
	// be empty if no request ID is available such as the request failed due
	// to a connection error.
	RequestID() string
}

// NewRequestFailure returns a new request error wrapper for the given Error
// provided.
func NewRequestFailure(err Error, statusCode int, reqID string) RequestFailure {
	return newRequestError(err, statusCode, reqID)
}
