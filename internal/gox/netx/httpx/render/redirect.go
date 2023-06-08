package render

import (
	"fmt"
	"net/http"
)

// Redirect redirects the http request to new location and writes redirect response.
func Redirect(
	w http.ResponseWriter,
	r *http.Request,
	location string,
	code int,
) error {
	if (code < http.StatusMultipleChoices || code > http.StatusPermanentRedirect) && code != http.StatusCreated {
		return fmt.Errorf("cannot redirect with status code %d", code)
	}
	http.Redirect(w, r, location, code)
	return nil
}
