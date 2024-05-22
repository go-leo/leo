// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cqrs

import (
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	http1 "net/http"
)

func NewCQRSHTTPServer(
	endpoints interface {
	},
	opts []http.ServerOption,
	mdw ...endpoint.Middleware,
) http1.Handler {
	router := mux.NewRouter()
	return router
}

type cQRSHTTPClient struct {
}

func NewCQRSHTTPClient(
	scheme string,
	instance string,
	opts []http.ClientOption,
	mdw ...endpoint.Middleware,
) interface {
} {
	return &cQRSHTTPClient{}
}
