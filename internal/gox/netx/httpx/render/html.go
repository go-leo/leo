package render

import (
	"html/template"
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/slicex"
)

// HTML executes template and writes its result with custom ContentType for response.
func HTML(
	w http.ResponseWriter,
	data any,
	template *template.Template,
	name ...string,
) error {
	writeContentType(w, []string{"text/html; charset=utf-8"})
	if slicex.IsEmpty(name) {
		return template.Execute(w, data)
	}
	return template.ExecuteTemplate(w, name[0], data)
}
