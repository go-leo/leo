package cmdx

// Middleware allows us to write something like decorators to Command.
// It can execute something before Execute or after.
type Middleware interface {
	// Decorate wraps the underlying Command, adding some functionality.
	Decorate(cmd Command) Command
}

// The MiddlewareFunc type is an adapter to allow the use of ordinary functions as Middleware.
// If f is a function with the appropriate signature, MiddlewareFunc(f) is a Middleware that calls f.
type MiddlewareFunc func(cmd Command) Command

// Decorate call f(cmd).
func (f MiddlewareFunc) Decorate(cmd Command) Command {
	return f(cmd)
}

// Chain decorates the given Command with all middlewares.
func Chain(cmd Command, middlewares ...Middleware) (chain Command) {
	chain = cmd
	for i := len(middlewares) - 1; i >= 0; i-- {
		chain = middlewares[i].Decorate(chain)
	}
	return chain
}
