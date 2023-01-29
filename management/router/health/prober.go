package health

type Prober interface {
	// Check if the target of this Prober is healthy
	// If nil is returned, target is healthy, otherwise target is not healthy
	Check(target string) error
}
