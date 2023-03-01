package health

import "net/http"

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

func (s status) Code() StatusCode {
	return s.code
}

func (s status) Description() string {
	return s.description
}

const (
	UnknownCode      StatusCode = "Unknown"
	UpCode           StatusCode = "Up"
	DownCode         StatusCode = "Down"
	OutOfServiceCode StatusCode = "OutOfService"
)

var (
	unknownStatus      = status{code: UnknownCode}
	upStatus           = status{code: UpCode}
	downStatus         = status{code: DownCode}
	outOfServiceStatus = status{code: OutOfServiceCode}
)

// UnknownStatus indicating that the component or subsystem is in an unknown state.
func UnknownStatus(description string) Status {
	st := unknownStatus
	st.description = description
	return st
}

// UpStatus indicating that the component or subsystem is functioning as expected.
func UpStatus(description string) Status {
	st := upStatus
	st.description = description
	return st
}

// DownStatus indicating that the component or subsystem has suffered an unexpected failure.
func DownStatus(description string) Status {
	st := downStatus
	st.description = description
	return st
}

// OutOfServiceStatus indicating that the component or subsystem has been taken out of service and should not be used.
func OutOfServiceStatus(description string) Status {
	st := outOfServiceStatus
	st.description = description
	return st
}

type HttpHealthStatusMapper interface {
	MapStatus(status Status) int
}

type httpHealthStatusMapper map[StatusCode]int

func (mapper httpHealthStatusMapper) MapStatus(status Status) int {
	return mapper[status.Code()]
}

func DefaultHttpHealthStatusMapper() HttpHealthStatusMapper {
	return httpHealthStatusMapper{
		UnknownCode:      http.StatusOK,
		UpCode:           http.StatusOK,
		DownCode:         http.StatusServiceUnavailable,
		OutOfServiceCode: http.StatusServiceUnavailable,
	}
}
