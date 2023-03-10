package health

import "net/http"

type HttpHealthStatusMapper interface {
	MapStatus(status Status) int
}

type httpHealthStatusMapper map[StatusCode]int

func (mapper httpHealthStatusMapper) MapStatus(status Status) int {
	return mapper[status.Code()]
}

func DefaultHttpHealthStatusMapper() HttpHealthStatusMapper {
	return httpHealthStatusMapper{
		UnknownCode():      http.StatusOK,
		UpCode():           http.StatusOK,
		DownCode():         http.StatusServiceUnavailable,
		OutOfServiceCode(): http.StatusServiceUnavailable,
	}
}
