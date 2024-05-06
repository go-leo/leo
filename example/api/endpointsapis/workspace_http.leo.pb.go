// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package endpointsapis

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	mux "github.com/gorilla/mux"
	protojson "google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	io "io"
	http1 "net/http"
	strconv "strconv"
)

func NewWorkspacesHTTPServer(
	endpoints interface {
		ListWorkspaces() endpoint.Endpoint
		GetWorkspace() endpoint.Endpoint
		CreateWorkspace() endpoint.Endpoint
		UpdateWorkspace() endpoint.Endpoint
		DeleteWorkspace() endpoint.Endpoint
	},
	mdw []endpoint.Middleware,
	opts ...http.ServerOption,
) http1.Handler {
	r := mux.NewRouter()
	r.Name("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").
		Methods("GET").
		Path("/v1/projects/{project}/locations/{location}/workspaces").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.ListWorkspaces(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &ListWorkspacesRequest{}
				vars := mux.Vars(r)
				req.Parent = fmt.Sprintf("projects/%s/locations/%s", vars["project"], vars["location"])
				queries := r.URL.Query()
				if v, err := strconv.ParseInt(queries.Get("page_size"), 10, 32); err != nil {
					return nil, err
				} else {
					req.PageSize = int32(v)
				}
				req.PageToken = queries.Get("page_token")
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*ListWorkspacesResponse)
				_ = resp
				w.WriteHeader(http1.StatusOK)
				data, err := protojson.Marshal(resp)
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
	r.Name("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").
		Methods("GET").
		Path("/v1/projects/{project}/locations/{location}/workspaces/{workspace}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.GetWorkspace(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &GetWorkspaceRequest{}
				vars := mux.Vars(r)
				req.Name = fmt.Sprintf("projects/%s/locations/%s/workspaces/%s", vars["project"], vars["location"], vars["workspace"])
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Workspace)
				_ = resp
				w.WriteHeader(http1.StatusOK)
				data, err := protojson.Marshal(resp)
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
	r.Name("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").
		Methods("POST").
		Path("/v1/projects/{project}/locations/{location}/workspaces").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.CreateWorkspace(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &CreateWorkspaceRequest{}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				if err := protojson.Unmarshal(body, req.Workspace); err != nil {
					return nil, err
				}
				vars := mux.Vars(r)
				req.Parent = fmt.Sprintf("projects/%s/locations/%s", vars["project"], vars["location"])
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Workspace)
				_ = resp
				w.WriteHeader(http1.StatusOK)
				data, err := protojson.Marshal(resp)
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
	r.Name("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").
		Methods("PATCH").
		Path("/v1/projects/{project}/locations/{location}/Workspaces/{Workspace}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.UpdateWorkspace(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &UpdateWorkspaceRequest{}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				if err := protojson.Unmarshal(body, req.Workspace); err != nil {
					return nil, err
				}
				vars := mux.Vars(r)
				req.Name = fmt.Sprintf("projects/%s/locations/%s/Workspaces/%s", vars["project"], vars["location"], vars["Workspace"])
				queries := r.URL.Query()
				mask, err := fieldmaskpb.New(req.Workspace, queries["update_mask"]...)
				if err != nil {
					return nil, err
				}
				req.UpdateMask = mask
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Workspace)
				_ = resp
				w.WriteHeader(http1.StatusOK)
				data, err := protojson.Marshal(resp)
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
	r.Name("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").
		Methods("DELETE").
		Path("/v1/projects/{project}/locations/{location}/workspaces/{workspace}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.DeleteWorkspace(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &DeleteWorkspaceRequest{}
				vars := mux.Vars(r)
				req.Name = fmt.Sprintf("projects/%s/locations/%s/workspaces/%s", vars["project"], vars["location"], vars["workspace"])
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				_ = resp
				w.WriteHeader(http1.StatusOK)
				data, err := protojson.Marshal(resp)
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
	return r
}

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
	r := mux.NewRouter()
	r.Name("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").
		Methods("GET").
		Path("/v1/projects/{project}/locations/{location}/workspaces")
	r.Name("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").
		Methods("GET").
		Path("/v1/projects/{project}/locations/{location}/workspaces/{workspace}")
	r.Name("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").
		Methods("POST").
		Path("/v1/projects/{project}/locations/{location}/workspaces")
	r.Name("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").
		Methods("PATCH").
		Path("/v1/projects/{project}/locations/{location}/Workspaces/{Workspace}")
	r.Name("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").
		Methods("DELETE").
		Path("/v1/projects/{project}/locations/{location}/workspaces/{workspace}")
	return &httpWorkspacesClient{
		listWorkspaces: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req := obj.(*ListWorkspacesRequest)
					var body io.Reader
					r, err := http1.NewRequestWithContext(ctx, "GET", "", body)
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
					req := obj.(*GetWorkspaceRequest)
					var body io.Reader
					r, err := http1.NewRequestWithContext(ctx, "GET", "", body)
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
					req := obj.(*CreateWorkspaceRequest)
					var body io.Reader
					data, err := protojson.Marshal(req.Workspace)
					if err != nil {
						return nil, err
					}
					body := bytes.NewBuffer(data)
					r, err := http1.NewRequestWithContext(ctx, "POST", "", body)
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
					req := obj.(*UpdateWorkspaceRequest)
					var body io.Reader
					data, err := protojson.Marshal(req.Workspace)
					if err != nil {
						return nil, err
					}
					body := bytes.NewBuffer(data)
					r, err := http1.NewRequestWithContext(ctx, "PATCH", "", body)
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
					req := obj.(*DeleteWorkspaceRequest)
					var body io.Reader
					r, err := http1.NewRequestWithContext(ctx, "DELETE", "", body)
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
