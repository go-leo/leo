package httpx

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"

	"google.golang.org/protobuf/proto"

	"github.com/go-leo/leo/common/iox"
)

type ResponseHelper struct {
	Response *http.Response
}

func (helper ResponseHelper) Body() io.ReadCloser {
	return helper.Response.Body
}

func (helper ResponseHelper) BytesBody() ([]byte, error) {
	defer iox.QuiteClose(helper.Response.Body)
	data, err := io.ReadAll(helper.Response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (helper ResponseHelper) TextBody() (string, error) {
	data, err := helper.BytesBody()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (helper ResponseHelper) FormBody() (url.Values, error) {
	data, err := helper.TextBody()
	if err != nil {
		return nil, err
	}
	return url.ParseQuery(data)
}

func (helper ResponseHelper) ObjectBody(body any, unmarshal func([]byte, any) error) error {
	data, err := helper.BytesBody()
	if err != nil {
		return err
	}
	return unmarshal(data, body)
}

func (helper ResponseHelper) JSONBody(body any) error {
	return helper.ObjectBody(body, json.Unmarshal)
}

func (helper ResponseHelper) XMLBody(body any) error {
	return helper.ObjectBody(body, xml.Unmarshal)
}

func (helper ResponseHelper) ProtobufBody(body proto.Message) error {
	unmarshal := func(data []byte, v any) error {
		m := v.(proto.Message)
		return proto.Unmarshal(data, m)
	}
	return helper.ObjectBody(body, unmarshal)
}

func (helper ResponseHelper) GobBody(body proto.Message) error {
	unmarshal := func(data []byte, v any) error {
		return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
	}
	return helper.ObjectBody(body, unmarshal)
}

func (helper ResponseHelper) FileBody(file io.Writer) error {
	defer iox.QuiteClose(helper.Response.Body)
	return iox.Copy(file, helper.Body())
}
