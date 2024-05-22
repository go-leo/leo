package assembler

import (
	"context"
	demo "github.com/go-leo/leo/v3/example/api/cqrs"
	"github.com/go-leo/leo/v3/example/internal/demo/command"
	"github.com/go-leo/leo/v3/example/internal/demo/model"
	"github.com/go-leo/leo/v3/example/internal/demo/query"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DemoServiceAssembler struct {
}

func (d DemoServiceAssembler) FromAsyncGetUsersRequest(ctx context.Context, request *demo.AsyncGetUsersRequest) (*query.AsyncGetUsersArgs, context.Context, error) {
	return &query.AsyncGetUsersArgs{
		PageNo:   request.GetPageNo(),
		PageSize: request.GetPageSize(),
	}, ctx, nil
}

func (d DemoServiceAssembler) ToAsyncGetUsersResponse(ctx context.Context, request *demo.AsyncGetUsersRequest, res *query.AsyncGetUsersRes) (*demo.AsyncGetUsersResponse, error) {
	list := make([]*demo.AsyncGetUsersResponse_User, 0, len(res.List))
	for _, user := range res.List {
		list = append(list, &demo.AsyncGetUsersResponse_User{
			Name:   user.Name,
			Age:    user.Age,
			Salary: user.Salary,
			Token:  user.Token,
		})
	}
	return &demo.AsyncGetUsersResponse{
		Users: list,
	}, nil
}

func (d DemoServiceAssembler) FromAsyncDeleteUsersRequest(ctx context.Context, request *demo.AsyncDeleteUsersRequest) (*command.AsyncDeleteUsersArgs, context.Context, error) {
	return &command.AsyncDeleteUsersArgs{
		Names: request.GetNames(),
	}, ctx, nil
}

func (d DemoServiceAssembler) ToAsyncDeleteUsersResponse(ctx context.Context, request *demo.AsyncDeleteUsersRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}

func (d DemoServiceAssembler) FromCreateUserRequest(ctx context.Context, request *demo.CreateUserRequest) (*command.CreateUserArgs, context.Context, error) {
	return &command.CreateUserArgs{
		User: &model.User{
			Name:   request.GetName(),
			Age:    request.GetAge(),
			Salary: request.GetSalary(),
			Token:  request.GetToken(),
		},
	}, ctx, nil
}

func (d DemoServiceAssembler) ToCreateUserResponse(ctx context.Context, request *demo.CreateUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (d DemoServiceAssembler) FromGetUsersRequest(ctx context.Context, request *demo.GetUsersRequest) (*query.GetUsersArgs, context.Context, error) {
	return &query.GetUsersArgs{
		PageNo:   request.GetPageNo(),
		PageSize: request.GetPageSize(),
	}, ctx, nil
}

func (d DemoServiceAssembler) ToGetUsersResponse(ctx context.Context, request *demo.GetUsersRequest, res *query.GetUsersRes) (*demo.GetUsersResponse, error) {
	list := make([]*demo.GetUsersResponse_User, 0, len(res.List))
	for _, user := range res.List {
		list = append(list, &demo.GetUsersResponse_User{
			Name:   user.Name,
			Age:    user.Age,
			Salary: user.Salary,
			Token:  user.Token,
		})
	}
	return &demo.GetUsersResponse{
		Users: list,
	}, nil
}

func (d DemoServiceAssembler) FromDeleteUserRequest(ctx context.Context, request *demo.DeleteUsersRequest) (*command.DeleteUserArgs, context.Context, error) {
	return &command.DeleteUserArgs{
		Name: request.GetName(),
	}, ctx, nil
}

func (d DemoServiceAssembler) ToDeleteUserResponse(ctx context.Context, request *demo.DeleteUsersRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}

func (d DemoServiceAssembler) FromUpdateUserRequest(ctx context.Context, request *demo.UpdateUserRequest) (*command.UpdateUserArgs, context.Context, error) {
	return &command.UpdateUserArgs{
		User: &model.User{
			Name:   request.GetName(),
			Age:    request.GetAge(),
			Salary: request.GetSalary(),
			Token:  request.GetToken(),
		},
	}, ctx, nil
}

func (d DemoServiceAssembler) ToUpdateUserResponse(ctx context.Context, request *demo.UpdateUserRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}

func (d DemoServiceAssembler) FromGetUserRequest(ctx context.Context, request *demo.GetUserRequest) (*query.GetUserArgs, context.Context, error) {
	return &query.GetUserArgs{
		User: &model.User{
			Name:   request.GetName(),
			Age:    request.GetAge(),
			Salary: request.GetSalary(),
			Token:  request.GetToken(),
		},
	}, ctx, nil
}

func (d DemoServiceAssembler) ToGetUserResponse(ctx context.Context, request *demo.GetUserRequest, res *query.GetUserRes) (*demo.GetUserResponse, error) {
	return &demo.GetUserResponse{
		Name:   res.User.Name,
		Age:    res.User.Age,
		Salary: res.User.Salary,
		Token:  res.User.Token,
	}, nil
}

func NewDemoServiceAssembler() cq.DemoServiceAssembler {
	return &DemoServiceAssembler{}
}
