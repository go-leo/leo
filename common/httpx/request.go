package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/go-leo/leo/common/stringx"
)

type RequestBuilder struct {
	method  string
	uri     *url.URL
	query   url.Values
	headers http.Header
	body    io.Reader
	err     error
}

func (builder *RequestBuilder) Method(method string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	if stringx.IsBlank(method) {
		builder.err = errors.New("method is blank")
		return builder
	}
	builder.method = method
	return builder
}

func (builder *RequestBuilder) Get() *RequestBuilder {
	return builder.Method(http.MethodGet)
}

func (builder *RequestBuilder) Head() *RequestBuilder {
	return builder.Method(http.MethodHead)
}

func (builder *RequestBuilder) Post() *RequestBuilder {
	return builder.Method(http.MethodPost)
}

func (builder *RequestBuilder) Put() *RequestBuilder {
	return builder.Method(http.MethodPut)
}

func (builder *RequestBuilder) Patch() *RequestBuilder {
	return builder.Method(http.MethodPatch)
}

func (builder *RequestBuilder) Delete() *RequestBuilder {
	return builder.Method(http.MethodDelete)
}

func (builder *RequestBuilder) Connect() *RequestBuilder {
	return builder.Method(http.MethodConnect)
}

func (builder *RequestBuilder) Options() *RequestBuilder {
	return builder.Method(http.MethodOptions)
}

func (builder *RequestBuilder) Trace() *RequestBuilder {
	return builder.Method(http.MethodTrace)
}

func (builder *RequestBuilder) URL(uri *url.URL) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	builder.uri = uri
	return builder
}

func (builder *RequestBuilder) URLString(urlString string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	if strings.HasPrefix(urlString, "ws:") {
		urlString = "http:" + strings.TrimPrefix(urlString, "ws:")
	} else if strings.HasPrefix(urlString, "wss") {
		urlString = "http:" + strings.TrimPrefix(urlString, "https:")
	}
	uri, err := url.Parse(urlString)
	if err != nil {
		builder.err = err
		return builder
	}
	return builder.URL(uri)
}

func (builder *RequestBuilder) Query(name, value string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	if builder.query == nil {
		builder.query = make(url.Values)
	}
	builder.query.Set(name, value)
	return builder
}

func (builder *RequestBuilder) AddQuery(name, value string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	if builder.query == nil {
		builder.query = make(url.Values)
	}
	builder.query.Add(name, value)
	return builder
}

func (builder *RequestBuilder) RemoveQuery(name string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	if builder.query == nil {
		builder.query = make(url.Values)
	}
	builder.query.Del(name)
	return builder
}

func (builder *RequestBuilder) QueryString(q string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	query, err := url.ParseQuery(q)
	if err != nil {
		builder.err = err
		return builder
	}
	return builder.Queries(query)
}

func (builder *RequestBuilder) Queries(q url.Values) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	builder.query = q
	return builder
}

func (builder *RequestBuilder) Header(name, value string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	if builder.headers == nil {
		builder.headers = make(http.Header)
	}
	builder.headers.Set(name, value)
	return builder
}

func (builder *RequestBuilder) AddHeader(name, value string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	if builder.headers == nil {
		builder.headers = make(http.Header)
	}
	builder.headers.Add(name, value)
	return builder
}

func (builder *RequestBuilder) RemoveHeader(name string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	if builder.headers == nil {
		builder.headers = make(http.Header)
	}
	builder.headers.Del(name)
	return builder
}

func (builder *RequestBuilder) Headers(header http.Header) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	builder.headers = header.Clone()
	return builder
}

func (builder *RequestBuilder) Body(body io.Reader, contentType string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	builder.body = body
	builder.Header("Content-Type", contentType)
	return builder
}

func (builder *RequestBuilder) BytesBody(body []byte, contentType string) *RequestBuilder {
	reader := bytes.NewReader(body)
	return builder.Body(reader, contentType)
}

func (builder *RequestBuilder) TextBody(body string, contentType string) *RequestBuilder {
	reader := strings.NewReader(body)
	return builder.Body(reader, contentType)
}

func (builder *RequestBuilder) JSONBody(body any) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	data, err := json.Marshal(body)
	if err != nil {
		builder.err = err
		return builder
	}
	return builder.BytesBody(data, "application/json")
}

func (builder *RequestBuilder) XMLBody(body any) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	data, err := xml.Marshal(body)
	if err != nil {
		builder.err = err
		return builder
	}
	return builder.BytesBody(data, "application/xml")
}

func (builder *RequestBuilder) ProtobufBody(body proto.Message) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	data, err := proto.Marshal(body)
	if err != nil {
		builder.err = err
		return builder
	}
	return builder.BytesBody(data, "application/x-protobuf")
}

func (builder *RequestBuilder) FormBody(form url.Values) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	return builder.TextBody(form.Encode(), "application/x-www-form-urlencoded")
}

type FormData struct {
	FieldName string
	Value     string
	File      fs.File
}

func (builder *RequestBuilder) MultipartBody(formData ...*FormData) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	payload := new(bytes.Buffer)
	writer := multipart.NewWriter(payload)
	for _, datum := range formData {
		if datum.File != nil {
			stat, err := datum.File.Stat()
			if err != nil {
				builder.err = err
				return builder
			}
			mf, err := writer.CreateFormFile(datum.FieldName, stat.Name())
			if err != nil {
				builder.err = err
				return builder
			}
			if _, err = io.Copy(mf, datum.File); err != nil {
				builder.err = err
				return builder
			}
		} else {
			_ = writer.WriteField(datum.FieldName, datum.Value)
		}
	}
	if err := writer.Close(); err != nil {
		builder.err = err
		return builder
	}
	return builder.Body(payload, writer.FormDataContentType())
}

func (builder *RequestBuilder) Build(ctx context.Context) (*http.Request, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	if stringx.IsBlank(builder.method) {
		return nil, errors.New("method is blank")
	}
	if builder.uri == nil {
		return nil, errors.New("url is nil")
	}
	req, err := http.NewRequestWithContext(ctx, builder.method, builder.uri.String(), builder.body)
	if err != nil {
		return nil, err
	}
	for key, values := range builder.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	return req, nil
}
