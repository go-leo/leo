package log

import (
	"context"
	"time"

	"github.com/go-leo/stringx"

	"github.com/go-leo/leo/v2/log"
)

type FieldBuilder struct {
	fields []log.F
}

func NewFieldBuilder() *FieldBuilder {
	return &FieldBuilder{}
}

func (f *FieldBuilder) System(system string) *FieldBuilder {
	if stringx.IsBlank(system) {
		return f
	}
	f.fields = append(f.fields, log.F{K: "system", V: system})
	return f
}

func (f *FieldBuilder) StartTime(startTime time.Time) *FieldBuilder {
	f.fields = append(f.fields, log.F{K: "start_time", V: startTime.Format(time.RFC3339)})
	return f
}

func (f *FieldBuilder) Deadline(ctx context.Context) *FieldBuilder {
	if d, ok := ctx.Deadline(); ok {
		f.fields = append(f.fields, log.F{K: "deadline", V: d.Format(time.RFC3339)})
	}
	return f
}

func (f *FieldBuilder) Latency(duration time.Duration) *FieldBuilder {
	f.fields = append(f.fields, log.F{K: "latency", V: duration.String()})
	return f
}

func (f *FieldBuilder) Method(method string) *FieldBuilder {
	if stringx.IsBlank(method) {
		return f
	}
	f.fields = append(f.fields, log.F{K: "method", V: method})
	return f
}

func (f *FieldBuilder) Path(path string) *FieldBuilder {
	if stringx.IsBlank(path) {
		return f
	}
	f.fields = append(f.fields, log.F{K: "path", V: path})
	return f
}

func (f *FieldBuilder) MetaData(md map[string][]string) *FieldBuilder {
	f.fields = append(f.fields, log.F{K: "meta_data", V: md})
	return f
}

func (f *FieldBuilder) Status(status string) *FieldBuilder {
	if stringx.IsBlank(status) {
		return f
	}
	f.fields = append(f.fields, log.F{K: "status", V: status})
	return f
}

func (f *FieldBuilder) PeerAddress(status string) *FieldBuilder {
	if stringx.IsBlank(status) {
		return f
	}
	f.fields = append(f.fields, log.F{K: "peer.address", V: status})
	return f
}

func (f *FieldBuilder) Error(err string) *FieldBuilder {
	if stringx.IsBlank(err) {
		return f
	}
	f.fields = append(f.fields, log.F{K: "error", V: err})
	return f
}

func (f *FieldBuilder) Request(req any) *FieldBuilder {
	f.fields = append(f.fields, log.F{K: "request", V: req})
	return f
}

func (f *FieldBuilder) Response(response any) *FieldBuilder {
	f.fields = append(f.fields, log.F{K: "response", V: response})
	return f
}

func (f *FieldBuilder) Build() []log.F {
	return f.fields
}
