package health

type StatusAggregator interface {
	AggregateStatus(statuses ...Status) Status
}

type SliceStatusAggregator []Status

func (agg SliceStatusAggregator) AggregateStatus(statuses ...Status) Status {
	for _, orderStatus := range agg {
		for _, status := range statuses {
			if orderStatus.Code() == status.Code() {
				return status
			}
		}
	}
	return UnknownStatus("")
}

func DefaultStatusAggregator() StatusAggregator {
	return SliceStatusAggregator{
		DownStatus(""),
		OutOfServiceStatus(""),
		UpStatus(""),
		UnknownStatus(""),
	}
}
