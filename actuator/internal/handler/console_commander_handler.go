package handler

import (
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/console"
)

type ConsoleCommanderHandler struct {
	ConsoleCommander console.Commander
}

func (h *ConsoleCommanderHandler) Pattern() string {
	return "/actuator/console/commander"
}

func (h *ConsoleCommanderHandler) Handle(http.ResponseWriter, *http.Request) {
	// TODO implement me
	panic("implement me")
}
