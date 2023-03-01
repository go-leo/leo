package handler

import (
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/rpc"
)

type RPCProviderHandler struct {
	RPCProvider rpc.Provider
}

func (h *RPCProviderHandler) Pattern() string {
	return "/actuator/rpc/provider"
}

func (h *RPCProviderHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	// TODO implement me
	panic("implement me")
}
