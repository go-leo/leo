package health

type ServingStatus int32

const (
	UNKNOWN     ServingStatus = 0
	SERVING     ServingStatus = 1
	NOT_SERVING ServingStatus = 2
)

var ServingStatusName = map[ServingStatus]string{
	UNKNOWN:     "UNKNOWN",
	SERVING:     "SERVING",
	NOT_SERVING: "NOT_SERVING",
}

func (ss ServingStatus) String() string {
	return ServingStatusName[ss]
}

type HealthCheckResponse struct {
	Status ServingStatus `json:"status,omitempty"`
}

func (x *HealthCheckResponse) GetStatus() ServingStatus {
	if x != nil {
		return x.Status
	}
	return UNKNOWN
}
