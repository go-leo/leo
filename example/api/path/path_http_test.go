package path

import (
	"github.com/gorilla/mux"
	http1 "net/http"
	"testing"
)

func TestRouterURL(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", func(writer http1.ResponseWriter, request *http1.Request) {}).
		Name("article")
	url, err := r.Get("article").URL("category", "technology", "id", "42")
	if err != nil {
		panic(err)
	}
	t.Log(url)

	router := mux.NewRouter()
	router.NewRoute().Path("/v1/string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}").HandlerFunc(func(http1.ResponseWriter, *http1.Request) {}).Name("/leo.example.demo.v1.Path/String")
	router.NewRoute().Path("/v1/opt_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}").HandlerFunc(func(http1.ResponseWriter, *http1.Request) {}).Name("/leo.example.demo.v1.Path/OptString")

	path, err := router.Get("/leo.example.demo.v1.Path/String").URL("class", "10", "shelf", "20", "book", "30", "family", "40")
	if err != nil {
		panic(err)
	}
	t.Log(path)

}
