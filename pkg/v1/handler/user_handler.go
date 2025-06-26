// Package handler implements the main grpc calls
package handler

import (
	"context"

	"github.com/dawit_hopes/grpc_micro_service/internal/domain/models"
	pb "github.com/dawit_hopes/grpc_micro_service/internal/pb"
	interfaces "github.com/dawit_hopes/grpc_micro_service/pkg/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	usecase interfaces.UseCaseInterface
	pb.UnimplementedUserServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &UserHandler{usecase: usecase}
	pb.RegisterUserServiceServer(grpcServer, userGrpc)
}

func (srv *UserHandler) transformUserRPC(req *pb.CreateUserRequest) models.User {
	return models.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}
}

func (srv *UserHandler) transformUserModel(user models.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{Name: user.Name, Email: user.Email, Id: user.ID.Hex()}
}

func (srv *UserHandler) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.SuccessResponse, error) {
	data := srv.transformUserRPC(req)

	if data.Email == "" || data.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "all fields must be provided")
	}

	if err := srv.usecase.Create(data); err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{
		Message: "Successfuly created a user",
	}, nil
}

func (srv *UserHandler) Get(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	user, err := srv.usecase.Get(req.GetId())

	if err != nil {
		return &pb.UserProfileResponse{}, err
	}

	return srv.transformUserModel(user), nil
}

func (srv *UserHandler) Delete(ctx context.Context, req *pb.SingleUserRequest) (*pb.SuccessResponse, error) {
	if err := srv.usecase.Delete(req.GetId()); err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{
		Message: "User deleted Successfully",
	}, nil
}

func (srv *UserHandler) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.SuccessResponse, error) {

	if req.GetName() == "" || req.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "all fields must be provided")
	}
	newUser := models.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	if err := srv.usecase.Update(req.GetId(), newUser); err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{Message: "Successfully updated a user"}, nil
}
