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
	"google.golang.org/protobuf/types/known/wrapperspb"
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

	return s.toUserDto(&user), nil
}

// Get user by id
func (s *UserSvc) GetUserByID(ctx context.Context, id *wrappers.StringValue) (*protoapi.UserDto, error) {
	log.Printf("GetUserByID(%s) \n", id.GetValue())

	user, err := s.urepo.FindByID(id.GetValue())
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return s.toUserDto(&user), nil
}

// Get all users
func (s *UserSvc) ListUsers(ctx context.Context, e *empty.Empty) (*protoapi.UserDtoList, error) {
	log.Printf("ListUsers() \n")

	var udtos []*protoapi.UserDto
	users, err := s.urepo.FindAll()
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	for _, u := range users {
		udtos = append(udtos, s.toUserDto(&u))
	}

	userDtoList := &protoapi.UserDtoList{
		List: udtos,
	}

	return userDtoList, nil
}

//
func (s *UserSvc) UpdateUser(ctx context.Context, u *protoapi.UserDto) (*protoapi.UserDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

// Delete user by id
func (s *UserSvc) DeleteUserByID(ctx context.Context, id *wrappers.StringValue) (*wrappers.BoolValue, error) {
	log.Printf("DeleteUserByID(%s) \n", id.GetValue())

	deleted, err := s.urepo.DeleteByID(id.GetValue())
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return &wrapperspb.BoolValue{Value: deleted}, nil
}

// mapping user model to userdto
func (s *UserSvc) toUserDto(u *model.User) *protoapi.UserDto {
	userid, _ := hex.DecodeString(u.ID.Hex())
	udto := &protoapi.UserDto{
		Id:    string(userid),
		Name:  u.Name,
		Email: u.Email,
	}
	return udto
}
