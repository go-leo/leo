package assembler

import (
	"context"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/leo/v3/example/api/demo"
	"github.com/go-leo/leo/v3/example/internal/demo/command"
	"github.com/go-leo/leo/v3/example/internal/demo/model"
	"github.com/go-leo/leo/v3/example/internal/demo/query"
	"github.com/go-leo/leo/v3/metadatax"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DemoAssembler struct {
}

func (assembler *DemoAssembler) FromCreateUserRequest(ctx context.Context, request *demo.CreateUserRequest) (*command.CreateUserArgs, context.Context, error) {
	return &command.CreateUserArgs{
		User: &model.User{
			UserId: 0,
			Name:   request.GetUser().GetName(),
			Age:    request.GetUser().GetAge(),
			Salary: request.GetUser().GetSalary(),
			Token:  request.GetUser().GetToken(),
			Avatar: request.GetUser().GetAvatar(),
		},
	}, ctx, nil
}

func (assembler *DemoAssembler) ToCreateUserResponse(ctx context.Context, request *demo.CreateUserRequest, metadata metadatax.Metadata) (*demo.CreateUserResponse, error) {
	return &demo.CreateUserResponse{UserId: convx.ToUint64(metadata.Get("id"))}, nil
}

func (assembler *DemoAssembler) FromDeleteUserRequest(ctx context.Context, request *demo.DeleteUsersRequest) (*command.DeleteUserArgs, context.Context, error) {
	return &command.DeleteUserArgs{UserId: request.GetUserId()}, ctx, nil
}

func (assembler *DemoAssembler) ToDeleteUserResponse(ctx context.Context, request *demo.DeleteUsersRequest, metadata metadatax.Metadata) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}

func (assembler *DemoAssembler) FromUpdateUserRequest(ctx context.Context, request *demo.UpdateUserRequest) (*command.UpdateUserArgs, context.Context, error) {
	return &command.UpdateUserArgs{
		User: &model.User{
			UserId: request.GetUserId(),
			Name:   request.GetUser().GetName(),
			Age:    request.GetUser().GetAge(),
			Salary: request.GetUser().GetSalary(),
			Token:  request.GetUser().GetToken(),
			Avatar: request.GetUser().GetAvatar(),
		},
	}, ctx, nil
}

func (assembler *DemoAssembler) ToUpdateUserResponse(ctx context.Context, request *demo.UpdateUserRequest, metadata metadatax.Metadata) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}

func (assembler *DemoAssembler) FromGetUserRequest(ctx context.Context, request *demo.GetUserRequest) (*query.GetUserArgs, context.Context, error) {
	return &query.GetUserArgs{
		UserId: request.GetUserId(),
	}, ctx, nil
}

func (assembler *DemoAssembler) ToGetUserResponse(ctx context.Context, request *demo.GetUserRequest, res *query.GetUserRes) (*demo.GetUserResponse, error) {
	return &demo.GetUserResponse{
		User: &demo.User{
			UserId: res.User.UserId,
			Name:   res.User.Name,
			Age:    res.User.Age,
			Salary: res.User.Salary,
			Token:  res.User.Token,
			Avatar: res.User.Avatar,
		},
	}, nil
}

func (assembler *DemoAssembler) FromGetUsersRequest(ctx context.Context, request *demo.GetUsersRequest) (*query.GetUsersArgs, context.Context, error) {
	return &query.GetUsersArgs{
		PageNo:   request.GetPageNo(),
		PageSize: request.GetPageSize(),
	}, ctx, nil
}

func (assembler *DemoAssembler) ToGetUsersResponse(ctx context.Context, request *demo.GetUsersRequest, res *query.GetUsersRes) (*demo.GetUsersResponse, error) {
	list := make([]*demo.User, 0, len(res.List))
	for _, user := range res.List {
		list = append(list, &demo.User{
			UserId: user.UserId,
			Name:   user.Name,
			Age:    user.Age,
			Salary: user.Salary,
			Token:  user.Token,
			Avatar: user.Avatar,
		})
	}
	return &demo.GetUsersResponse{
		Users: list,
	}, nil
}

func (assembler *DemoAssembler) FromUploadUserAvatarRequest(ctx context.Context, request *demo.UploadUserAvatarRequest) (*command.UploadUserAvatarArgs, context.Context, error) {
	return &command.UploadUserAvatarArgs{
		UserId: request.GetUserId(),
		Avatar: request.GetAvatar().GetData(),
	}, ctx, nil
}

func (assembler *DemoAssembler) ToUploadUserAvatarResponse(ctx context.Context, request *demo.UploadUserAvatarRequest, metadata metadatax.Metadata) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}

func (assembler *DemoAssembler) FromGetUserAvatarRequest(ctx context.Context, request *demo.GetUserAvatarRequest) (*query.GetUserAvatarArgs, context.Context, error) {
	return &query.GetUserAvatarArgs{
		UserId: request.GetUserId(),
	}, ctx, nil
}

func (assembler *DemoAssembler) ToGetUserAvatarResponse(ctx context.Context, request *demo.GetUserAvatarRequest, res *query.GetUserAvatarRes) (*httpbody.HttpBody, error) {
	return &httpbody.HttpBody{
		ContentType: "image/jpg",
		Data:        res.Data,
		Extensions:  nil,
	}, nil
}

func (assembler *DemoAssembler) FromPushUsersRequest(ctx context.Context, request *http.HttpRequest) (*query.PushUsersArgs, context.Context, error) {
	return &query.PushUsersArgs{
		Data: request.GetBody(),
	}, ctx, nil
}

func (assembler *DemoAssembler) ToPushUsersResponse(ctx context.Context, request *http.HttpRequest, res *query.PushUsersRes) (*http.HttpResponse, error) {
	return &http.HttpResponse{
		Status:  200,
		Reason:  "OK",
		Headers: nil,
		Body:    res.Data,
	}, nil
}

//
//func (d DemoAssembler) FromAsyncGetUsersRequest(ctx context.Context, request *demo.AsyncGetUsersRequest) (*query.AsyncGetUsersArgs, context.Context, error) {
//	return &query.AsyncGetUsersArgs{
//		PageNo:   request.GetPageNo(),
//		PageSize: request.GetPageSize(),
//	}, ctx, nil
//}
//
//func (d DemoAssembler) ToAsyncGetUsersResponse(ctx context.Context, request *demo.AsyncGetUsersRequest, res *query.AsyncGetUsersRes) (*demo.AsyncGetUsersResponse, error) {
//	list := make([]*demo.AsyncGetUsersResponse_User, 0, len(res.List))
//	for _, user := range res.List {
//		list = append(list, &demo.AsyncGetUsersResponse_User{
//			Name:   user.Name,
//			Age:    user.Age,
//			Salary: user.Salary,
//			Token:  user.Token,
//		})
//	}
//	return &demo.AsyncGetUsersResponse{
//		Users: list,
//	}, nil
//}
//
//func (d DemoAssembler) FromAsyncDeleteUsersRequest(ctx context.Context, request *demo.AsyncDeleteUsersRequest) (*command.AsyncDeleteUsersArgs, context.Context, error) {
//	return &command.AsyncDeleteUsersArgs{
//		Names: request.GetNames(),
//	}, ctx, nil
//}
//
//func (d DemoAssembler) ToAsyncDeleteUsersResponse(ctx context.Context, request *demo.AsyncDeleteUsersRequest) (*emptypb.Empty, error) {
//	return new(emptypb.Empty), nil
//}
//
//func (d DemoAssembler) FromCreateUserRequest(ctx context.Context, request *demo.CreateUserRequest) (*command.CreateUserArgs, context.Context, error) {

//}
//
//func (d DemoAssembler) ToCreateUserResponse(ctx context.Context, request *demo.CreateUserRequest) (*emptypb.Empty, error) {
//	return &emptypb.Empty{}, nil
//}
//
//func (d DemoAssembler) FromGetUsersRequest(ctx context.Context, request *demo.GetUsersRequest) (*query.GetUsersArgs, context.Context, error) {

//}
//
//func (d DemoAssembler) ToGetUsersResponse(ctx context.Context, request *demo.GetUsersRequest, res *query.GetUsersRes) (*demo.GetUsersResponse, error) {

//}
//
//func (d DemoAssembler) FromDeleteUserRequest(ctx context.Context, request *demo.DeleteUsersRequest) (*command.DeleteUserArgs, context.Context, error) {

//}
//
//func (d DemoAssembler) ToDeleteUserResponse(ctx context.Context, request *demo.DeleteUsersRequest) (*emptypb.Empty, error) {
//	return new(emptypb.Empty), nil
//}
//
//func (d DemoAssembler) FromUpdateUserRequest(ctx context.Context, request *demo.UpdateUserRequest) (*command.UpdateUserArgs, context.Context, error) {
//	return &command.UpdateUserArgs{
//		User: &model.User{
//			Name:   request.GetName(),
//			Age:    request.GetAge(),
//			Salary: request.GetSalary(),
//			Token:  request.GetToken(),
//		},
//	}, ctx, nil
//}
//
//func (d DemoAssembler) ToUpdateUserResponse(ctx context.Context, request *demo.UpdateUserRequest) (*emptypb.Empty, error) {
//	return new(emptypb.Empty), nil
//}
//
//func (d DemoAssembler) FromGetUserRequest(ctx context.Context, request *demo.GetUserRequest) (*query.GetUserArgs, context.Context, error) {

//}
//
//func (d DemoAssembler) ToGetUserResponse(ctx context.Context, request *demo.GetUserRequest, res *query.GetUserRes) (*demo.GetUserResponse, error) {

//}

func NewDemoAssembler() demo.DemoAssembler {
	return &DemoAssembler{}
}
