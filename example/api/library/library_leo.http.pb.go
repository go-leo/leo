// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package library

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	mux "github.com/gorilla/mux"
	protojson "google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http1 "net/http"
	strings "strings"
)

type httpLibraryServiceClient struct {
	createShelf  endpoint.Endpoint
	getShelf     endpoint.Endpoint
	listShelves  endpoint.Endpoint
	deleteShelf  endpoint.Endpoint
	mergeShelves endpoint.Endpoint
	createBook   endpoint.Endpoint
	getBook      endpoint.Endpoint
	listBooks    endpoint.Endpoint
	deleteBook   endpoint.Endpoint
	updateBook   endpoint.Endpoint
	moveBook     endpoint.Endpoint
}

func (c *httpLibraryServiceClient) CreateShelf(ctx context.Context, request *CreateShelfRequest) (*Shelf, error) {
	rep, err := c.createShelf(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Shelf), nil
}

func (c *httpLibraryServiceClient) GetShelf(ctx context.Context, request *GetShelfRequest) (*Shelf, error) {
	rep, err := c.getShelf(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Shelf), nil
}

func (c *httpLibraryServiceClient) ListShelves(ctx context.Context, request *ListShelvesRequest) (*ListShelvesResponse, error) {
	rep, err := c.listShelves(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*ListShelvesResponse), nil
}

func (c *httpLibraryServiceClient) DeleteShelf(ctx context.Context, request *DeleteShelfRequest) (*emptypb.Empty, error) {
	rep, err := c.deleteShelf(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *httpLibraryServiceClient) MergeShelves(ctx context.Context, request *MergeShelvesRequest) (*Shelf, error) {
	rep, err := c.mergeShelves(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Shelf), nil
}

func (c *httpLibraryServiceClient) CreateBook(ctx context.Context, request *CreateBookRequest) (*Book, error) {
	rep, err := c.createBook(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}

func (c *httpLibraryServiceClient) GetBook(ctx context.Context, request *GetBookRequest) (*Book, error) {
	rep, err := c.getBook(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}

func (c *httpLibraryServiceClient) ListBooks(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error) {
	rep, err := c.listBooks(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*ListBooksResponse), nil
}

func (c *httpLibraryServiceClient) DeleteBook(ctx context.Context, request *DeleteBookRequest) (*emptypb.Empty, error) {
	rep, err := c.deleteBook(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *httpLibraryServiceClient) UpdateBook(ctx context.Context, request *UpdateBookRequest) (*Book, error) {
	rep, err := c.updateBook(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}

func (c *httpLibraryServiceClient) MoveBook(ctx context.Context, request *MoveBookRequest) (*Book, error) {
	rep, err := c.moveBook(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}

func NewLibraryServiceHTTPClient(
	instance string,
	mdw []endpoint.Middleware,
	opts ...http.ClientOption,
) interface {
	CreateShelf(ctx context.Context, request *CreateShelfRequest) (*Shelf, error)
	GetShelf(ctx context.Context, request *GetShelfRequest) (*Shelf, error)
	ListShelves(ctx context.Context, request *ListShelvesRequest) (*ListShelvesResponse, error)
	DeleteShelf(ctx context.Context, request *DeleteShelfRequest) (*emptypb.Empty, error)
	MergeShelves(ctx context.Context, request *MergeShelvesRequest) (*Shelf, error)
	CreateBook(ctx context.Context, request *CreateBookRequest) (*Book, error)
	GetBook(ctx context.Context, request *GetBookRequest) (*Book, error)
	ListBooks(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error)
	DeleteBook(ctx context.Context, request *DeleteBookRequest) (*emptypb.Empty, error)
	UpdateBook(ctx context.Context, request *UpdateBookRequest) (*Book, error)
	MoveBook(ctx context.Context, request *MoveBookRequest) (*Book, error)
} {
	router := mux.NewRouter()
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/CreateShelf").
		Methods("POST").
		Path("/v1/shelves")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/GetShelf").
		Methods("GET").
		Path("/v1/shelves/{shelf}")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/ListShelves").
		Methods("GET").
		Path("/v1/shelves")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/DeleteShelf").
		Methods("DELETE").
		Path("/v1/shelves/{shelf}")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/MergeShelves").
		Methods("POST").
		Path("/v1/shelves/{shelf}:merge")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/CreateBook").
		Methods("POST").
		Path("/v1/shelves/{shelf}/books")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/GetBook").
		Methods("GET").
		Path("/v1/shelves/{shelf}/books/{book}")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/ListBooks").
		Methods("GET").
		Path("/v1/shelves/{shelf}/books")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/DeleteBook").
		Methods("DELETE").
		Path("/v1/shelves/{shelf}/books/{book}")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/UpdateBook").
		Methods("PATCH").
		Path("/v1/shelves/{shelf}/books/{book}")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/MoveBook").
		Methods("POST").
		Path("/v1/shelves/{shelf}/books/{book}:move")
	return &httpLibraryServiceClient{
		createShelf: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*CreateShelfRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "POST"
					var url string
					var body io.Reader
					if req.Shelf != nil {
						data, err := protojson.Marshal(req.Shelf)
						if err != nil {
							return nil, err
						}
						body = bytes.NewBuffer(data)
					}
					var pairs []string
					path, err := router.Get("/google.example.library.v1.LibraryService/CreateShelf").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
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
		getShelf: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*GetShelfRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "GET"
					var url string
					var body io.Reader
					var pairs []string
					namedPathParameter := req.Name
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 2 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1])
					path, err := router.Get("/google.example.library.v1.LibraryService/GetShelf").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
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
		listShelves: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*ListShelvesRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "GET"
					var url string
					var body io.Reader
					var pairs []string
					path, err := router.Get("/google.example.library.v1.LibraryService/ListShelves").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := r.URL.Query()
					// page_sizePageSize int32
					// page_tokenPageToken string
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
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
		deleteShelf: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*DeleteShelfRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "DELETE"
					var url string
					var body io.Reader
					var pairs []string
					namedPathParameter := req.Name
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 2 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1])
					path, err := router.Get("/google.example.library.v1.LibraryService/DeleteShelf").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
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
		mergeShelves: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*MergeShelvesRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "POST"
					var url string
					var body io.Reader
					if req != nil {
						data, err := protojson.Marshal(req)
						if err != nil {
							return nil, err
						}
						body = bytes.NewBuffer(data)
					}
					var pairs []string
					namedPathParameter := req.Name
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 2 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1])
					path, err := router.Get("/google.example.library.v1.LibraryService/MergeShelves").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
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
		createBook: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*CreateBookRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "POST"
					var url string
					var body io.Reader
					if req.Book != nil {
						data, err := protojson.Marshal(req.Book)
						if err != nil {
							return nil, err
						}
						body = bytes.NewBuffer(data)
					}
					var pairs []string
					namedPathParameter := req.Parent
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 2 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1])
					path, err := router.Get("/google.example.library.v1.LibraryService/CreateBook").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
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
		getBook: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*GetBookRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "GET"
					var url string
					var body io.Reader
					var pairs []string
					namedPathParameter := req.Name
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 4 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1], "book", namedPathValues[3])
					path, err := router.Get("/google.example.library.v1.LibraryService/GetBook").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
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
		listBooks: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*ListBooksRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "GET"
					var url string
					var body io.Reader
					var pairs []string
					namedPathParameter := req.Parent
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 2 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1])
					path, err := router.Get("/google.example.library.v1.LibraryService/ListBooks").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := r.URL.Query()
					// page_sizePageSize int32
					// page_tokenPageToken string
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
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
		deleteBook: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*DeleteBookRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "DELETE"
					var url string
					var body io.Reader
					var pairs []string
					namedPathParameter := req.Name
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 4 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1], "book", namedPathValues[3])
					path, err := router.Get("/google.example.library.v1.LibraryService/DeleteBook").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
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
		updateBook: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*UpdateBookRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "PATCH"
					var url string
					var body io.Reader
					if req.Book != nil {
						data, err := protojson.Marshal(req.Book)
						if err != nil {
							return nil, err
						}
						body = bytes.NewBuffer(data)
					}
					var pairs []string
					if req.Book == nil {
						return nil, fmt.Errorf("%s is nil", "req.Book")
					}
					namedPathParameter := req.Book.Name
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 4 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1], "book", namedPathValues[3])
					path, err := router.Get("/google.example.library.v1.LibraryService/UpdateBook").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := r.URL.Query()
					// update_maskUpdateMask message
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
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
		moveBook: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*MoveBookRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "POST"
					var url string
					var body io.Reader
					if req != nil {
						data, err := protojson.Marshal(req)
						if err != nil {
							return nil, err
						}
						body = bytes.NewBuffer(data)
					}
					var pairs []string
					namedPathParameter := req.Name
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 4 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1], "book", namedPathValues[3])
					path, err := router.Get("/google.example.library.v1.LibraryService/MoveBook").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
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