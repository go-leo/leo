package consulx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-playground/form/v4"
	"google.golang.org/grpc"
	"net/url"
	"time"
)

func DefaultClientCreator(ctx context.Context, rawURL *url.URL, color string, dialOptions ...grpc.DialOption) (etcdv3.Client, error) {
	q := args{}
	decoder := form.NewDecoder()
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) { return time.ParseDuration(vals[0]) }, time.Duration(0))
	if err := decoder.Decode(&q, rawURL.Query()); err != nil {
		return nil, fmt.Errorf("malformed url parameters, %v", err)
	}

	username := rawURL.User.Username()
	password, _ := rawURL.User.Password()

	options := etcdv3.ClientOptions{
		Cert:          q.Cert,
		Key:           q.Key,
		CACert:        q.CACert,
		DialTimeout:   q.DialTimeout,
		DialKeepAlive: q.DialKeepAlive,
		DialOptions:   dialOptions,
		Username:      username,
		Password:      password,
	}
	return etcdv3.NewClient(ctx, q.Machines, options)
}

type args struct {
	Machines      []string      `form:"machines"`
	Cert          string        `form:"cert"`
	Key           string        `form:"key"`
	CACert        string        `form:"ca_cert"`
	DialTimeout   time.Duration `form:"dial_timeout"`
	DialKeepAlive time.Duration `form:"dial_keepalive"`
}
