package handler

import (
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/view"
)

type ViewControllerHandler struct {
	ViewController view.Controller
}

func (h *ViewControllerHandler) Pattern() string {
	return "/actuator/view/controller"
}

func (h *ViewControllerHandler) Handle(writer http.ResponseWriter, request *http.Request) {

}
