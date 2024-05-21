// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package path

import (
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	urlx "github.com/go-leo/gox/netx/urlx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	mux "github.com/gorilla/mux"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	io "io"
	http1 "net/http"
	url "net/url"
	strings "strings"
)

func NewNamedPathHTTPServer(
	endpoints interface {
		NamedPathString() endpoint.Endpoint
		NamedPathOptString() endpoint.Endpoint
		NamedPathWrapString() endpoint.Endpoint
		EmbedNamedPathString() endpoint.Endpoint
		EmbedNamedPathOptString() endpoint.Endpoint
		EmbedNamedPathWrapString() endpoint.Endpoint
	},
	mdw []endpoint.Middleware,
	opts ...http.ServerOption,
) http1.Handler {
	router := mux.NewRouter()
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/NamedPathString").
		Methods("GET").
		Path("/v1/string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.NamedPathString(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &NamedPathRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.String_ = fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family"))
				if varErr != nil {
					return nil, varErr
				}
				queries := r.URL.Query()
				var queryErr error
				req.OptString = proto.String(queries.Get("opt_string"))
				req.WrapString = wrapperspb.String(queries.Get("wrap_string"))
				if queryErr != nil {
					return nil, queryErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				_ = resp
				w.WriteHeader(http1.StatusOK)
				data, err := json.Marshal(resp)
				if err != nil {
					return err
				}
				if _, err := w.Write(data); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/NamedPathOptString").
		Methods("GET").
		Path("/v1/opt_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.NamedPathOptString(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &NamedPathRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.OptString = proto.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
				if varErr != nil {
					return nil, varErr
				}
				queries := r.URL.Query()
				var queryErr error
				req.String_ = queries.Get("string")
				req.WrapString = wrapperspb.String(queries.Get("wrap_string"))
				if queryErr != nil {
					return nil, queryErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				_ = resp
				w.WriteHeader(http1.StatusOK)
				data, err := json.Marshal(resp)
				if err != nil {
					return err
				}
				if _, err := w.Write(data); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/NamedPathWrapString").
		Methods("GET").
		Path("/v1/wrap_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.NamedPathWrapString(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &NamedPathRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.WrapString = wrapperspb.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
				if varErr != nil {
					return nil, varErr
				}
				queries := r.URL.Query()
				var queryErr error
				req.String_ = queries.Get("string")
				req.OptString = proto.String(queries.Get("opt_string"))
				if queryErr != nil {
					return nil, queryErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				_ = resp
				w.WriteHeader(http1.StatusOK)
				data, err := json.Marshal(resp)
				if err != nil {
					return err
				}
				if _, err := w.Write(data); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/EmbedNamedPathString").
		Methods("GET").
		Path("/v1/embed/string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.EmbedNamedPathString(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &EmbedNamedPathRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				if req.Embed == nil {
					req.Embed = &NamedPathRequest{}
				}
				req.Embed.String_ = fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				_ = resp
				w.WriteHeader(http1.StatusOK)
				data, err := json.Marshal(resp)
				if err != nil {
					return err
				}
				if _, err := w.Write(data); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString").
		Methods("GET").
		Path("/v1/embed/opt_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.EmbedNamedPathOptString(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &EmbedNamedPathRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				if req.Embed == nil {
					req.Embed = &NamedPathRequest{}
				}
				req.Embed.OptString = proto.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				_ = resp
				w.WriteHeader(http1.StatusOK)
				data, err := json.Marshal(resp)
				if err != nil {
					return err
				}
				if _, err := w.Write(data); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString").
		Methods("GET").
		Path("/v1/embed/wrap_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.EmbedNamedPathWrapString(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &EmbedNamedPathRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				if req.Embed == nil {
					req.Embed = &NamedPathRequest{}
				}
				req.Embed.WrapString = wrapperspb.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				_ = resp
				w.WriteHeader(http1.StatusOK)
				data, err := json.Marshal(resp)
				if err != nil {
					return err
				}
				if _, err := w.Write(data); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	return router
}

type namedPathHTTPClient struct {
	namedPathString          endpoint.Endpoint
	namedPathOptString       endpoint.Endpoint
	namedPathWrapString      endpoint.Endpoint
	embedNamedPathString     endpoint.Endpoint
	embedNamedPathOptString  endpoint.Endpoint
	embedNamedPathWrapString endpoint.Endpoint
}

func (c *namedPathHTTPClient) NamedPathString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error) {
	rep, err := c.namedPathString(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *namedPathHTTPClient) NamedPathOptString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error) {
	rep, err := c.namedPathOptString(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *namedPathHTTPClient) NamedPathWrapString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error) {
	rep, err := c.namedPathWrapString(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *namedPathHTTPClient) EmbedNamedPathString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error) {
	rep, err := c.embedNamedPathString(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *namedPathHTTPClient) EmbedNamedPathOptString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error) {
	rep, err := c.embedNamedPathOptString(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *namedPathHTTPClient) EmbedNamedPathWrapString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error) {
	rep, err := c.embedNamedPathWrapString(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewNamedPathHTTPClient(
	scheme string,
	instance string,
	mdw []endpoint.Middleware,
	opts ...http.ClientOption,
) interface {
	NamedPathString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
	NamedPathOptString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
	NamedPathWrapString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
	EmbedNamedPathString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
	EmbedNamedPathOptString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
	EmbedNamedPathWrapString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
} {
	router := mux.NewRouter()
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/NamedPathString").
		Methods("GET").
		Path("/v1/string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/NamedPathOptString").
		Methods("GET").
		Path("/v1/opt_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/NamedPathWrapString").
		Methods("GET").
		Path("/v1/wrap_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/EmbedNamedPathString").
		Methods("GET").
		Path("/v1/embed/string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString").
		Methods("GET").
		Path("/v1/embed/opt_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().
		Name("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString").
		Methods("GET").
		Path("/v1/embed/wrap_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	return &namedPathHTTPClient{
		namedPathString: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*NamedPathRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					namedPathParameter := req.GetString_()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 8 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
					path, err := router.Get("/leo.example.path.v1.NamedPath/NamedPathString").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					queries["opt_string"] = append(queries["opt_string"], req.GetOptString())
					queries["wrap_string"] = append(queries["wrap_string"], req.GetWrapString().GetValue())
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					return nil, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		namedPathOptString: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*NamedPathRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					namedPathParameter := req.GetOptString()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 8 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
					path, err := router.Get("/leo.example.path.v1.NamedPath/NamedPathOptString").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					queries["string"] = append(queries["string"], req.GetString_())
					queries["wrap_string"] = append(queries["wrap_string"], req.GetWrapString().GetValue())
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					return nil, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		namedPathWrapString: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*NamedPathRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					namedPathParameter := req.GetWrapString().GetValue()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 8 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
					path, err := router.Get("/leo.example.path.v1.NamedPath/NamedPathWrapString").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					queries["string"] = append(queries["string"], req.GetString_())
					queries["opt_string"] = append(queries["opt_string"], req.GetOptString())
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					return nil, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		embedNamedPathString: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*EmbedNamedPathRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					namedPathParameter := req.GetEmbed().GetString_()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 8 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
					path, err := router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathString").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					return nil, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		embedNamedPathOptString: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*EmbedNamedPathRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					namedPathParameter := req.GetEmbed().GetOptString()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 8 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
					path, err := router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					return nil, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		embedNamedPathWrapString: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*EmbedNamedPathRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					namedPathParameter := req.GetEmbed().GetWrapString().GetValue()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 8 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
					path, err := router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					return nil, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
	}
}
