package health

type StatusCode string

// Status express state of a component or subsystem.
type Status interface {
	// Code return the code of the Status.
	Code() StatusCode
	// Description return the description of the Status.
	Description() string
}

type status struct {
	code        StatusCode
	description string
}

func (s status) Code() StatusCode    { return s.code }
func (s status) Description() string { return s.description }

func UnknownCode() StatusCode { return "Unknown" }

func UpCode() StatusCode { return "Up" }

func DownCode() StatusCode { return "Down" }

func OutOfServiceCode() StatusCode { return "OutOfService" }

func NewStatus(code StatusCode, description string) Status {
	return status{code: code, description: description}
}

// UnknownStatus indicating that the component or subsystem is in an unknown state.
func UnknownStatus(description string) Status {
	return status{code: UnknownCode(), description: description}
}

// UpStatus indicating that the component or subsystem is functioning as expected.
func UpStatus(description string) Status {
	return status{code: UpCode(), description: description}
}

// DownStatus indicating that the component or subsystem has suffered an unexpected failure.
func DownStatus(description string) Status {
	return status{code: DownCode(), description: description}
}

// OutOfServiceStatus indicating that the component or subsystem has been taken out of service and should not be used.
func OutOfServiceStatus(description string) Status {
	return status{code: OutOfServiceCode(), description: description}
}
