package path

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	http1 "net/http"
	"testing"
	"time"
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

func TestTimestamp(t *testing.T) {
	timestamp := timestamppb.New(time.Now())
	data, _ := protojson.Marshal(timestamp)
	t.Log(string(data))

	var newTimestamp timestamppb.Timestamp
	_ = protojson.Unmarshal(data, &newTimestamp)
	t.Log(newTimestamp.AsTime())
}

func TestDuration(t *testing.T) {
	duration := durationpb.New(time.Hour)
	data, _ := protojson.Marshal(duration)
	t.Log(string(data))

	var newDuration durationpb.Duration
	_ = protojson.Unmarshal(data, &newDuration)
	t.Log(newDuration.AsDuration())
}

func TestNil(t *testing.T) {
	data, _ := json.Marshal(nil)
	t.Log(string(data))
}
