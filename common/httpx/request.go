package httpx

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"

	"github.com/go-leo/leo/common/stringx"
)

type RequestBuilder struct {
	method  string
	uri     *url.URL
	query   url.Values
	headers http.Header
	body    io.Reader
	cookies []*http.Cookie
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
	if lenGetter, ok := body.(interface{ Len() int }); ok {
		builder.Header("Content-Length", strconv.Itoa(lenGetter.Len()))
	}
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

func (builder *RequestBuilder) FormBody(form url.Values) *RequestBuilder {
	return builder.TextBody(form.Encode(), "application/x-www-form-urlencoded")
}

func (builder *RequestBuilder) ObjectBody(body any, marshal func(any) ([]byte, error), contentType string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	data, err := marshal(body)
	if err != nil {
		builder.err = err
		return builder
	}
	return builder.BytesBody(data, contentType)
}

func (builder *RequestBuilder) JSONBody(body any) *RequestBuilder {
	return builder.ObjectBody(body, json.Marshal, "application/json")
}

func (builder *RequestBuilder) XMLBody(body any) *RequestBuilder {
	return builder.ObjectBody(body, xml.Marshal, "application/xml")
}

func (builder *RequestBuilder) ProtobufBody(body proto.Message) *RequestBuilder {
	marshal := func(v any) ([]byte, error) {
		message, _ := v.(proto.Message)
		return proto.Marshal(message)
	}
	return builder.ObjectBody(body, marshal, "application/x-protobuf")
}

func (builder *RequestBuilder) GobBody(body any) *RequestBuilder {
	marshal := func(v any) ([]byte, error) {
		var b bytes.Buffer
		if err := gob.NewEncoder(&b).Encode(v); err != nil {
			return nil, err
		}
		return b.Bytes(), nil
	}
	return builder.ObjectBody(body, marshal, "application/x-gob")
}

type FormData struct {
	FieldName string
	Value     string
	File      io.Reader
	Filename  string
}

func (builder *RequestBuilder) MultipartBody(formData ...*FormData) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	payload := new(bytes.Buffer)
	writer := multipart.NewWriter(payload)
	for _, form := range formData {
		if form.File != nil {
			mf, err := writer.CreateFormFile(form.FieldName, filepath.Base(form.Filename))
			if err != nil {
				builder.err = err
				return builder
			}
			if _, err = io.Copy(mf, form.File); err != nil {
				builder.err = err
				return builder
			}
		} else {
			_ = writer.WriteField(form.FieldName, form.Value)
		}
	}
	if err := writer.Close(); err != nil {
		builder.err = err
		return builder
	}
	return builder.Body(payload, writer.FormDataContentType())
}

func (builder *RequestBuilder) BasicAuth(username, password string) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	token := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return builder.CustomAuth("Basic", token)
}

func (builder *RequestBuilder) BearerAuth(token string) *RequestBuilder {
	return builder.CustomAuth("Bearer", token)
}

func (builder *RequestBuilder) CustomAuth(scheme, token string) *RequestBuilder {
	return builder.APIKey("Authorization", scheme+" "+token)
}

func (builder *RequestBuilder) APIKey(key string, value string) *RequestBuilder {
	return builder.Header(key, value)
}

func (builder *RequestBuilder) Cookie(cookie *http.Cookie) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	index := slices.IndexFunc(builder.cookies, func(c *http.Cookie) bool {
		return c.Name == cookie.Name
	})
	if index >= 0 {
		builder.cookies[index] = cookie
		return builder
	}
	return builder.AddCookie(cookie)
}

func (builder *RequestBuilder) AddCookie(cookie *http.Cookie) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	builder.cookies = append(builder.cookies, cookie)
	return builder
}

func (builder *RequestBuilder) RemoveCookie(cookie *http.Cookie) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	index := slices.IndexFunc(builder.cookies, func(c *http.Cookie) bool {
		return c.Name == cookie.Name
	})
	if index == -1 {
		return builder
	}
	builder.cookies = slices.Delete(builder.cookies, index, index+1)
	return builder
}

func (builder *RequestBuilder) Cookies(cookies ...*http.Cookie) *RequestBuilder {
	if builder.err != nil {
		return builder
	}
	builder.cookies = make([]*http.Cookie, 0, len(cookies))
	for _, cookie := range cookies {
		if cookie == nil {
			continue
		}
		builder.cookies = append(builder.cookies, cookie)
	}
	return builder
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
	for _, cookie := range builder.cookies {
		req.AddCookie(cookie)
	}
	return req, nil
}
