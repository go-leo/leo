package chainx

// Chain decorates the given Command with all middlewares in the chain.
func Chain(cmd Command, middlewares ...Middleware) Command {
	for i := len(middlewares) - 1; i >= 0; i-- {
		cmd = middlewares[i].Middleware(cmd)
	}
	return cmd
}
