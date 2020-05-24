package service

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/salprima/gocrud-grpc/internal/protoapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserService
type UserSvc struct {
}

//
func (s *UserSvc) CreateUser(ctx context.Context, u *protoapi.UserDto) (*empty.Empty, error) {
	log.Printf("CreateUser(%v) \n", u)
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

//
func (s *UserSvc) GetUserByID(ctx context.Context, id *wrappers.StringValue) (*protoapi.UserDto, error) {
	log.Printf("GetUserByID(%s) \n", id)
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}

//
func (s *UserSvc) ListUsers(ctx context.Context, e *empty.Empty) (*protoapi.UserDtoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}

//
func (s *UserSvc) UpdateUser(ctx context.Context, u *protoapi.UserDto) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

//
func (s *UserSvc) DeleteUserByID(ctx context.Context, id *wrappers.StringValue) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserByID not implemented")
}
