package binding

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/formx"
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/netx/httpx/binding/internal/multipart"
	"errors"
	"net/http"
)

const defaultMemory = 32 << 20

func Form(req *http.Request, obj any, tag string) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := req.ParseMultipartForm(defaultMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}

	return formx.Unmarshal(req.Form, obj, tag)
}

func PostForm(req *http.Request, obj any, tag string) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := req.ParseMultipartForm(defaultMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	return formx.Unmarshal(req.PostForm, obj, tag)
}

func MultipartForm(req *http.Request, obj any, tag string) error {
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	err := formx.Unmarshal(req.Form, obj, tag)
	if err != nil {
		return err
	}
	return multipart.MappingByPtr(req, obj, tag)
}
