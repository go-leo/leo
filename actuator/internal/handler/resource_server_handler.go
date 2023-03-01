package handler

import (
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/resource"
)

type ResourceServerHandler struct {
	ResourceServer resource.Server
}

func (h *ResourceServerHandler) Pattern() string {
	return "/actuator/resource/server"
}

func (h *ResourceServerHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	// TODO implement me
	panic("implement me")
}
