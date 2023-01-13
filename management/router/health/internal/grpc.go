package internal

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
)

type GRPCProber struct {
	request     *hv1.HealthCheckRequest
	timeout     time.Duration
	dialOptions []grpc.DialOption
}

func NewGRPCProber(timeout time.Duration, tlsConfig *tls.Config) *GRPCProber {
	request := hv1.HealthCheckRequest{}
	var dialOptions []grpc.DialOption
	if tlsConfig != nil {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	} else {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	return &GRPCProber{
		request:     &request,
		timeout:     timeout,
		dialOptions: dialOptions,
	}
}

func (probe *GRPCProber) Check(target string) error {
	ctx, cancel := context.WithTimeout(context.Background(), probe.timeout)
	defer cancel()
	connection, err := grpc.DialContext(ctx, target, probe.dialOptions...)
	if err != nil {
		return err
	}
	defer connection.Close()
	client := hv1.NewHealthClient(connection)
	response, err := client.Check(ctx, probe.request)
	if err != nil {
		return err
	}
	if response.Status != hv1.HealthCheckResponse_SERVING {
		return fmt.Errorf("gRPC %s serving status: %s", target, response.Status)
	}
	return nil
}
