package handler

import (
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/stream"
)

type SteamRouterHandler struct {
	SteamRouter stream.Router
}

func (h *SteamRouterHandler) Pattern() string {
	return "/actuator/steam/router"
}

func (h *SteamRouterHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	// TODO implement me
	panic("implement me")
}
