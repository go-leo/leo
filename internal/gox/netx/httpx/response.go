package httpx

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/protobuf/proto"
)

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
type ResponseHelper struct {
	err        error
	resp       *http.Response
	statusCode int
	headers    http.Header
	trailers   http.Header
	cookies    []*http.Cookie
	bodyBytes  []byte
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func NewResponseHelper(resp *http.Response, err error) *ResponseHelper {
	respHelper := &ResponseHelper{resp: resp, err: err}
	if err != nil {
		return respHelper
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		respHelper.err = err
		return respHelper
	}
	err = resp.Body.Close()
	if err != nil {
		respHelper.err = err
		return respHelper
	}

	respHelper.resp = resp
	respHelper.statusCode = resp.StatusCode
	respHelper.headers = resp.Header
	respHelper.trailers = resp.Trailer
	respHelper.cookies = resp.Cookies()
	respHelper.bodyBytes = body
	return respHelper
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) Err() error {
	return helper.err
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) StatusCode() (int, error) {
	if helper.err != nil {
		return 0, helper.err
	}
	return helper.statusCode, nil
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) Headers() (http.Header, error) {
	if helper.err != nil {
		return nil, helper.err
	}
	return helper.headers, nil
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) Trailer() (http.Header, error) {
	if helper.err != nil {
		return nil, helper.err
	}
	return helper.trailers, nil
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) Cookies() ([]*http.Cookie, error) {
	if helper.err != nil {
		return nil, helper.err
	}
	return helper.cookies, nil
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) Body() (io.ReadCloser, error) {
	if helper.err != nil {
		return nil, helper.err
	}
	return io.NopCloser(bytes.NewReader(helper.bodyBytes)), nil
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) BytesBody() ([]byte, error) {
	if helper.err != nil {
		return nil, helper.err
	}
	return helper.bodyBytes, nil
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) TextBody() (string, error) {
	if helper.err != nil {
		return "", helper.err
	}
	return string(helper.bodyBytes), nil
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) ObjectBody(body any, unmarshal func([]byte, any) error) error {
	if helper.err != nil {
		return helper.err
	}
	err := unmarshal(helper.bodyBytes, body)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal body, body is %s, %w", helper.bodyBytes, err)
	}
	return err
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) JSONBody(body any) error {
	if helper.err != nil {
		return helper.err
	}
	return helper.ObjectBody(body, json.Unmarshal)
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) XMLBody(body any) error {
	if helper.err != nil {
		return helper.err
	}
	return helper.ObjectBody(body, xml.Unmarshal)
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) ProtobufBody(body proto.Message) error {
	if helper.err != nil {
		return helper.err
	}
	unmarshal := func(data []byte, v any) error {
		m := v.(proto.Message)
		return proto.Unmarshal(data, m)
	}
	return helper.ObjectBody(body, unmarshal)
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) GobBody(body proto.Message) error {
	if helper.err != nil {
		return helper.err
	}
	unmarshal := func(data []byte, v any) error {
		return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
	}
	return helper.ObjectBody(body, unmarshal)
}

// Deprecated: Do not use. use github.com/go-leo/netx/httpx/receiver instead.
func (helper *ResponseHelper) FileBody(file io.Writer) error {
	if helper.err != nil {
		return helper.err
	}
	_, err := file.Write(helper.bodyBytes)
	return err
}
