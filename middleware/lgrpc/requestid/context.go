package requestid

import "context"

const key = "codeup.aliyun.com/qimao/leo/leo/requestid"

func FromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(key).(string)
	return val, ok
}

func NewContext(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, key, v)
}
