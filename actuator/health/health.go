package health

// Health is an component that contributes data to results returned from the HealthEndpoint
type Health interface {
	// Status return the status of the health.
	Status() Status
	// SetStatus set the status of the health.
	SetStatus(status Status) Health
	// Details return the details of the health.
	Details() map[string]any
	// SetDetails set the details of the health.
	SetDetails(details map[string]any) Health
	// With add a key-value pair to details of the health.
	With(key string, val any) Health
}

type health struct {
	status  Status
	details map[string]any
}

func (h *health) Status() Status {
	return h.status
}

func (h *health) SetStatus(status Status) Health {
	h.status = status
	return h
}

func (h *health) Details() map[string]any {
	return h.details
}

func (h *health) SetDetails(details map[string]any) Health {
	h.details = details
	return h
}

func (h *health) With(key string, val any) Health {
	if h.details == nil {
		h.details = make(map[string]any)
	}
	h.details[key] = val
	return h
}

func NewHealth(status Status, details map[string]any) Health {
	return &health{status: status, details: details}
}

func UnknownHealth() Health {
	return NewHealth(unknownStatus, nil)
}

func UpHealth() Health {
	return NewHealth(upStatus, nil)
}

func DownHealth() Health {
	return NewHealth(downStatus, nil)
}

func DownHealthWithError(err error) Health {
	return NewHealth(downStatus, map[string]any{"error": err})
}

func OutOfServiceHealth() Health {
	return NewHealth(outOfServiceStatus, nil)
}
