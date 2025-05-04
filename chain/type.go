package chain

import "net/http"

// MiddlewareFunc defines the signature of an HTTP middleware.
//
// A MiddlewareFunc is a higher-order function that wraps an
// [http.HandlerFunc], allowing the injection of behavior before and/or
// after the wrapped handler is executed.
//
// Example:
//
//	func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
//	    return func(w http.ResponseWriter, r *http.Request) {
//	        log.Printf("Request received: %s", r.URL.Path)
//	        next(w, r)
//	    }
//	}
type MiddlewareFunc func(next http.HandlerFunc) http.HandlerFunc
