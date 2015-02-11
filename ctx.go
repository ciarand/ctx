package ctx

import (
	"net/http"

	"golang.org/x/net/context"
)

// For reference, these are the canonical original forms:
/*

type Middleware func(http.Handler) http.Handler

type Handler interface {
	ServeHTTP(context.Context, http.ResponseWriter, *http.Request)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)
*/

// ContextAwareHandlerFunc is the basic http.HandlerFunc but it also accepts
// a Context as its first argument.
type ContextAwareHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

// ContextAwareHandler is the basic http.Handler, but augmented to take
// a Context as its first argument.
type ContextAwareHandler interface {
	ServeHTTP(context.Context, http.ResponseWriter, *http.Request)
}

// ContextAwareMiddleware is just like a regular middleware, but it takes
// a Context and expects something else that can too.
type ContextAwareMiddleware func(context.Context, ContextAwareHandler) ContextAwareHandler
