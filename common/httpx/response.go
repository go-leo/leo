package httpx

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/go-leo/iox"
)

// Deprecated: Do not use. use github.com/go-leo/netx instead.
type ResponseHelper struct {
	Response *http.Response
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) StatusCode() int {
	return helper.Response.StatusCode
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) Headers() http.Header {
	return helper.Response.Header
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) LastModified() string {
	return helper.Response.Header.Get("Last-Modified")
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) Etag() string {
	return helper.Response.Header.Get("Etag")
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) CacheControl() string {
	return helper.Response.Header.Get("Cache-Control")
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) Trailer() http.Header {
	return helper.Response.Trailer
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) Cookies() []*http.Cookie {
	return helper.Response.Cookies()
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) Body() io.ReadCloser {
	return helper.Response.Body
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) BytesBody() ([]byte, error) {
	defer iox.QuiteClose(helper.Response.Body)
	b := new(bytes.Buffer)
	_, err := io.Copy(b, helper.Response.Body)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) TextBody() (string, error) {
	defer iox.QuiteClose(helper.Response.Body)
	b := new(strings.Builder)
	_, err := io.Copy(b, helper.Response.Body)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) ObjectBody(body any, unmarshal func([]byte, any) error) error {
	data, err := helper.BytesBody()
	if err != nil {
		return err
	}
	return unmarshal(data, body)
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) JSONBody(body any) error {
	return helper.ObjectBody(body, json.Unmarshal)
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) XMLBody(body any) error {
	return helper.ObjectBody(body, xml.Unmarshal)
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) ProtobufBody(body proto.Message) error {
	unmarshal := func(data []byte, v any) error {
		m := v.(proto.Message)
		return proto.Unmarshal(data, m)
	}
	return helper.ObjectBody(body, unmarshal)
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) GobBody(body proto.Message) error {
	unmarshal := func(data []byte, v any) error {
		return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
	}
	return helper.ObjectBody(body, unmarshal)
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (helper ResponseHelper) FileBody(file io.Writer) error {
	defer iox.QuiteClose(helper.Response.Body)
	return iox.Copy(file, helper.Body())
}
