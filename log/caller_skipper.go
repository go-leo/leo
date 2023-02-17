package log

type CallerSkipper interface {
	// SkipCaller returns a Logger that will offset the call stack by the specified number of frames
	// when logging call site information.
	SkipCaller(depth int) Logger
}
