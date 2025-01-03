// Code generated by protoc-gen-leo-cqrs. DO NOT EDIT.

package cqrs

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	command "github.com/go-leo/leo/v3/example/internal/cqrs/command"
	query "github.com/go-leo/leo/v3/example/internal/cqrs/query"
	metadatax "github.com/go-leo/leo/v3/metadatax"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func NewCQRSBus(
	createUser command.CreateUser,
	deleteUser command.DeleteUser,
	updateUser query.UpdateUser,
	findUser query.FindUser,
) (cqrs.Bus, error) {
	bus := cqrs.NewBus()
	if err := bus.RegisterCommand(createUser); err != nil {
		return nil, err
	}
	if err := bus.RegisterCommand(deleteUser); err != nil {
		return nil, err
	}
	if err := bus.RegisterQuery(updateUser); err != nil {
		return nil, err
	}
	if err := bus.RegisterQuery(findUser); err != nil {
		return nil, err
	}
	return bus, nil
}

// CQRSAssembler responsible for completing the transformation between domain model objects and DTOs
type CQRSAssembler interface {

	// FromCreateUserRequest convert request to command arguments
	FromCreateUserRequest(ctx context.Context, request *CreateUserRequest) (*command.CreateUserArgs, context.Context, error)

	// ToCreateUserResponse convert query result to response
	ToCreateUserResponse(ctx context.Context, request *CreateUserRequest, metadata metadatax.Metadata) (*emptypb.Empty, error)

	// FromDeleteUserRequest convert request to command arguments
	FromDeleteUserRequest(ctx context.Context, request *DeleteUserRequest) (*command.DeleteUserArgs, context.Context, error)

	// ToDeleteUserResponse convert query result to response
	ToDeleteUserResponse(ctx context.Context, request *DeleteUserRequest, metadata metadatax.Metadata) (*DeleteUserResponse, error)

	// FromUpdateUserRequest convert request to query arguments
	FromUpdateUserRequest(ctx context.Context, request *UpdateUserRequest) (*query.UpdateUserArgs, context.Context, error)

	// ToUpdateUserResponse convert query result to response
	ToUpdateUserResponse(ctx context.Context, request *UpdateUserRequest, res *query.UpdateUserRes) (*UpdateUserResponse, error)

	// FromFindUserRequest convert request to query arguments
	FromFindUserRequest(ctx context.Context, request *FindUserRequest) (*query.FindUserArgs, context.Context, error)

	// ToFindUserResponse convert query result to response
	ToFindUserResponse(ctx context.Context, request *FindUserRequest, res *query.FindUserRes) (*GetUserResponse, error)
}

// cQRSCqrsService implement the CQRSService with CQRS pattern
type cQRSCqrsService struct {
	bus       cqrs.Bus
	assembler CQRSAssembler
}

func (svc *cQRSCqrsService) CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error) {
	args, ctx, err := svc.assembler.FromCreateUserRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	metadata, err := svc.bus.Exec(ctx, args)
	if err != nil {
		return nil, err
	}
	return svc.assembler.ToCreateUserResponse(ctx, request, metadata)
}

func (svc *cQRSCqrsService) DeleteUser(ctx context.Context, request *DeleteUserRequest) (*DeleteUserResponse, error) {
	args, ctx, err := svc.assembler.FromDeleteUserRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	metadata, err := svc.bus.Exec(ctx, args)
	if err != nil {
		return nil, err
	}
	return svc.assembler.ToDeleteUserResponse(ctx, request, metadata)
}

func (svc *cQRSCqrsService) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*UpdateUserResponse, error) {
	args, ctx, err := svc.assembler.FromUpdateUserRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	res, err := svc.bus.Query(ctx, args)
	if err != nil {
		return nil, err
	}
	return svc.assembler.ToUpdateUserResponse(ctx, request, res.(*query.UpdateUserRes))
}

func (svc *cQRSCqrsService) FindUser(ctx context.Context, request *FindUserRequest) (*GetUserResponse, error) {
	args, ctx, err := svc.assembler.FromFindUserRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	res, err := svc.bus.Query(ctx, args)
	if err != nil {
		return nil, err
	}
	return svc.assembler.ToFindUserResponse(ctx, request, res.(*query.FindUserRes))
}

func NewCQRSCqrsService(bus cqrs.Bus, assembler CQRSAssembler) CQRSService {
	return &cQRSCqrsService{bus: bus, assembler: assembler}
}
