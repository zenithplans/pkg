// Package chain provides a lightweight and composable middleware chaining
// utility for HTTP handlers. It enables developers to structure their
// middleware in a modular, ordered, and reusable fashion.
//
// This package is inspired by [https://github.com/justinas/alice] pkg.
package chain

import "net/http"

// Chain holds a list of middleware to be applied sequentially.
// Middleware are executed in the order they are appended.
//
// A zero-value Chain is valid and represents an empty chain.
//
// Example:
//
//	c := chain.New(LoggingMiddleware, AuthMiddleware)
//	finalHandler := c.Then(MyHandler)
type Chain struct {
	middleware []MiddlewareFunc
}

// New creates a new Chain with the given middleware.
// Middleware are applied in the order provided, with the first middleware
// wrapping the second, and so on.
//
// Example:
//
//	c := chain.New(m1, m2, m3)
func New(middleware ...MiddlewareFunc) Chain {
	return Chain{middleware: middleware}
}

// Then wraps the provided final [http.HandlerFunc] with the chain of
// middleware. The middleware are applied in the order they were added.
//
// Example:
//
//	final := c.Then(myHandler)
//	http.Handle("/secure", final)
func (c Chain) Then(final http.HandlerFunc) http.HandlerFunc {
	for i := len(c.middleware) - 1; i >= 0; i-- {
		final = c.middleware[i](final)
	}
	return final
}

// Append returns a new Chain with the given middleware appended to the end.
// The original Chain remains unchanged.
//
// Example:
//
//	c := chain.New(m1).Append(m2, m3)
func (c Chain) Append(mw ...MiddlewareFunc) Chain {
	extendedChain := make([]MiddlewareFunc, len(c.middleware)+len(mw))

	// add original middleware to extendedChain
	copy(extendedChain, c.middleware)

	// add new middleware to extendedChain
	copy(extendedChain[len(c.middleware):], mw)

	return Chain{middleware: extendedChain}
}

// Extend returns a new Chain by concatenating another Chain to the current
// one. Middleware from the other chain are appended after the current
// chain's middleware.
//
// Example:
//
//	c1 := chain.New(m1)
//	c2 := chain.New(m2, m3)
//	c3 := c1.Extend(c2)
func (c Chain) Extend(other Chain) Chain {
	return c.Append(other.middleware...)
}
