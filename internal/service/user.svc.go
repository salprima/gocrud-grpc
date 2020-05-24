package service

import (
	"context"
	"encoding/hex"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/salprima/gocrud-grpc/internal/model"
	"github.com/salprima/gocrud-grpc/internal/protoapi"
	"github.com/salprima/gocrud-grpc/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserService
type UserSvc struct {
	urepo *repository.UserRepo
}

// Instantiate new user Service
func NewUserSvc(urepo *repository.UserRepo) *UserSvc {
	return &UserSvc{
		urepo: urepo,
	}
}

// Create new user
func (s *UserSvc) CreateUser(ctx context.Context, dto *protoapi.UserDto) (*protoapi.UserDto, error) {
	log.Printf("CreateUser(%v) \n", dto)

	newUser := &model.User{
		Name:  dto.Name,
		Email: dto.Email,
	}

	user, err := s.urepo.Save(newUser)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	userid, err := hex.DecodeString(user.ID.Hex())
	userDto := &protoapi.UserDto{
		Id:    string(userid),
		Name:  user.Name,
		Email: user.Email,
	}

	return userDto, nil
}

// Get user by id
func (s *UserSvc) GetUserByID(ctx context.Context, id *wrappers.StringValue) (*protoapi.UserDto, error) {
	log.Printf("GetUserByID(%s) \n", id.GetValue())

	user, err := s.urepo.FindByID(id.GetValue())
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	userid, err := hex.DecodeString(user.ID.Hex())
	userDto := &protoapi.UserDto{
		Id:    string(userid),
		Name:  user.Name,
		Email: user.Email,
	}

	return userDto, nil
}

//
func (s *UserSvc) ListUsers(ctx context.Context, e *empty.Empty) (*protoapi.UserDtoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}

//
func (s *UserSvc) UpdateUser(ctx context.Context, u *protoapi.UserDto) (*protoapi.UserDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

//
func (s *UserSvc) DeleteUserByID(ctx context.Context, id *wrappers.StringValue) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserByID not implemented")
}
