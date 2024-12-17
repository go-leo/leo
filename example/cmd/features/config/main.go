package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/configx"
	"github.com/go-leo/leo/v3/configx/consulx"
	"github.com/go-leo/leo/v3/configx/environx"
	"github.com/go-leo/leo/v3/configx/filex"
	"github.com/go-leo/leo/v3/example/configs"
	"github.com/hashicorp/consul/api"
	"google.golang.org/protobuf/encoding/protojson"
	"os"
	"time"
)

var client *api.Client

var filename = "./grpc.json"

func init() {
	os.Setenv("leo_run_env", "debug")
	var err error
	client, err = api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	_, err = client.KV().Put(&api.KVPair{
		Key: "redis",
		Value: []byte(`redis:
    addr: localhost:6379
    db: 0
    network: tcp
    password: test`),
	}, nil)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, []byte(`{"grpc":{"addr":"localhost","port":8080}}`), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	config, err := configx.Load[*configs.Application](
		context.Background(),
		configx.WithResource(
			&environx.Resource{Prefix: "leo_"},
			&filex.Resource{Formatter: configx.Json{}, Filename: filename},
			&consulx.Resource{Formatter: configx.Yaml{}, Client: client, Key: "redis"},
		),
		configx.WithParser(configx.Env{}, configx.Json{}, configx.Yaml{}),
	)
	if err != nil {
		panic(err)
	}
	data, _ := protojson.Marshal(config)
	println(string(data))

	configC, errC, stop := configx.Watch[*configs.Application](
		context.Background(),
		configx.WithResource(
			&environx.Resource{Prefix: "leo_"},
			&filex.Resource{Formatter: configx.Json{}, Filename: filename},
			&consulx.Resource{Formatter: configx.Yaml{}, Client: client, Key: "redis"},
		),
		configx.WithParser(configx.Env{}, configx.Json{}, configx.Yaml{}),
	)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			os.Setenv("leo_run_env", time.Now().Format(time.RFC3339))
		}
	}()

	go func() {
		for {
			time.Sleep(3 * time.Second)
			err := os.WriteFile(filename, []byte(fmt.Sprintf(`{"grpc":{"addr":"localhost","port":%d}}`, randx.Int31n(65535))), 0644)
			if err != nil {
				panic(err)
			}
		}
	}()

	go func() {
		for {
			time.Sleep(3 * time.Second)
			_, err := client.KV().Put(&api.KVPair{
				Key: "redis",
				Value: []byte(fmt.Sprintf(`redis:
	addr: "localhost:%d"
	db: %d
	network: tcp
	password: test`, randx.Intn(65535), randx.Intn(200))),
			}, nil)
			if err != nil {
				panic(err)
			}
		}
	}()

	for {
		select {
		case config := <-configC:
			data, _ := protojson.Marshal(config)
			println(string(data))
		case err := <-errC:
			if errors.Is(err, configx.ErrStopWatch) {
				return
			}
			println(err)
		case <-time.After(time.Minute):
			stop()
		}
	}
}
