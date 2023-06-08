package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/netx/httpx/render"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

type Handler struct {
	Logger log.Logger
}

func (h *Handler) Pattern() string {
	return "/log/level"
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	type payload struct {
		Level string `json:"level"`
	}
	switch r.Method {
	case http.MethodGet:
		_ = render.JSON(w, payload{Level: h.Logger.GetLevel().Name()}, render.PureJSON())
	case http.MethodPut:
		requestedLvl, err := decodePutRequest(r.Header.Get("Content-Type"), r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = render.JSON(w, errorResponse{Error: err.Error()}, render.PureJSON())
			return
		}
		h.Logger.SetLevel(requestedLvl)
		_ = render.JSON(w, payload{Level: h.Logger.GetLevel().Name()}, render.PureJSON())
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = render.JSON(w, errorResponse{Error: "Only GET and PUT are supported."}, render.PureJSON())

	}
}

// Decodes incoming PUT requests and returns the requested logging level.
func decodePutRequest(contentType string, r *http.Request) (log.Level, error) {
	if contentType == "application/x-www-form-urlencoded" {
		return decodePutURL(r)
	}
	return decodePutJSON(r.Body)
}

func decodePutURL(r *http.Request) (log.Level, error) {
	lvl := r.FormValue("level")
	if lvl == "" {
		return nil, errors.New("must specify logging level")
	}
	return log.ParseLevel(lvl)
}

func decodePutJSON(body io.Reader) (log.Level, error) {
	var pld struct {
		Level string `json:"level"`
	}
	if err := json.NewDecoder(body).Decode(&pld); err != nil {
		return nil, fmt.Errorf("malformed request body: %v", err)
	}
	if pld.Level == "" {
		return nil, errors.New("must specify logging level")
	}
	return log.ParseLevel(pld.Level)
}
