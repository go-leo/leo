// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package endpointsapis

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

type httpWorkspacesClient struct {
	listWorkspaces  endpoint.Endpoint
	getWorkspace    endpoint.Endpoint
	createWorkspace endpoint.Endpoint
	updateWorkspace endpoint.Endpoint
	deleteWorkspace endpoint.Endpoint
}

func (c *httpWorkspacesClient) ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	rep, err := c.listWorkspaces(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*ListWorkspacesResponse), nil
}

func (c *httpWorkspacesClient) GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error) {
	rep, err := c.getWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *httpWorkspacesClient) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	rep, err := c.createWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *httpWorkspacesClient) UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error) {
	rep, err := c.updateWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *httpWorkspacesClient) DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	rep, err := c.deleteWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewWorkspacesHTTPClient(
	instance string,
	mdw []endpoint.Middleware,
	opts ...http.ClientOption,
) interface {
	ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error)
	GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error)
	CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error)
	UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error)
	DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error)
} {
	router := mux.NewRouter()
	router.NewRoute().
		Name("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").
		Methods("GET").
		Path("/v1/projects/{project}/locations/{location}/workspaces")
	router.NewRoute().
		Name("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").
		Methods("GET").
		Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}")
	router.NewRoute().
		Name("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").
		Methods("POST").
		Path("/v1/projects/{project}/locations/{location}/workspaces")
	router.NewRoute().
		Name("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").
		Methods("PATCH").
		Path("/v1/projects/{project}/locations/{location}/Workspaces/{Workspac}")
	router.NewRoute().
		Name("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").
		Methods("DELETE").
		Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}")
	return &httpWorkspacesClient{
		listWorkspaces: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*ListWorkspacesRequest)
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
					if len(namedPathValues) != 4 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3])
					path, err := router.Get("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").URLPath(pairs...)
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
		getWorkspace: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*GetWorkspaceRequest)
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
					if len(namedPathValues) != 6 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3], "workspac", namedPathValues[5])
					path, err := router.Get("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").URLPath(pairs...)
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
		createWorkspace: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*CreateWorkspaceRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "POST"
					var url string
					var body io.Reader
					if req.Workspace != nil {
						data, err := protojson.Marshal(req.Workspace)
						if err != nil {
							return nil, err
						}
						body = bytes.NewBuffer(data)
					}
					var pairs []string
					namedPathParameter := req.Parent
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 4 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3])
					path, err := router.Get("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").URLPath(pairs...)
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
		updateWorkspace: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*UpdateWorkspaceRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "PATCH"
					var url string
					var body io.Reader
					if req.Workspace != nil {
						data, err := protojson.Marshal(req.Workspace)
						if err != nil {
							return nil, err
						}
						body = bytes.NewBuffer(data)
					}
					var pairs []string
					namedPathParameter := req.Name
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 6 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3], "Workspac", namedPathValues[5])
					path, err := router.Get("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").URLPath(pairs...)
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
		deleteWorkspace: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*DeleteWorkspaceRequest)
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
					if len(namedPathValues) != 6 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3], "workspac", namedPathValues[5])
					path, err := router.Get("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").URLPath(pairs...)
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
