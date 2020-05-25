package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/salprima/gocrud-grpc/internal/protoapi"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func init() {

	args := os.Args[1:]
	var configname string = "default-config"
	if len(args) > 0 {
		configname = args[0] + "-config"
	}
	fmt.Printf("loading config file %s.yml \n", configname)

	viper.SetConfigName(configname)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s.yml \n", err))
	}

}

func main() {

	port := ":" + viper.GetString("app.grpc.port")
	client, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	userApi := protoapi.NewUserApiClient(client)

	// Create users from sample
	log.Println("Creating sample users...")
	var allUsers []*protoapi.UserDto // temporarily load users in memory
	sampleUsers := sampleUsers()
	for _, userDto := range sampleUsers {
		user, err := userApi.CreateUser(context.Background(), userDto)
		if err != nil {
			log.Fatalf("Fail CreateUser: %v \n", err)
		}
		allUsers = append(allUsers, user)
		log.Println(user)
	}
	log.Println("Sample users created")
	fmt.Printf("%s\n%s\n%s\n", ">", ">", ">")

	// Get user by id
	log.Println("Get user by id...")
	ronaldo := allUsers[0]
	userID := &wrapperspb.StringValue{Value: ronaldo.Id}
	user, err := userApi.GetUserByID(context.Background(), userID)
	if err != nil {
		log.Fatalf("Fail GetUserByID: %v \n", err)
	}
	log.Println(user)
	log.Println("Get user by id DONE")
	fmt.Printf("%s\n%s\n%s\n", ">", ">", ">")

	// Get all users
	log.Println("Listing all users...")
	users, err := userApi.ListUsers(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Fail ListUsers: %v \n", err)
	}
	log.Println("ListUsers >>>>>")
	for _, u := range users.List {
		log.Println(u)
	}
	log.Println("Listing all users DONE")
	fmt.Printf("%s\n%s\n%s\n", ">", ">", ">")

	// Delete user by id
	log.Println("Delete user by id...")
	messi := allUsers[1]
	userIdToDelete := messi.Id //change this id according to your data
	deleted, err := userApi.DeleteUserByID(context.Background(), &wrapperspb.StringValue{Value: userIdToDelete})
	if err != nil {
		log.Fatalf("Fail DeleteUserByID: %v \n", err)
	}
	if deleted.GetValue() {
		log.Printf("UserID(%s) deleted \n", userIdToDelete)
	}
	log.Println("Delete user by id DONE")
	fmt.Printf("%s\n%s\n%s\n", ">", ">", ">")

	// Update user
	log.Println("Updating user...")
	delpiero := allUsers[2]
	userDtoUpdate := &protoapi.UserDto{
		Id:    delpiero.Id, //change this id according to your data
		Name:  "Alessandro Del Piero",
		Email: "adp10@juventus.com",
	}
	// don't confuse, this method return user before update
	// if you want to return updated result, you can call GetByUserID after
	oldUser, err := userApi.UpdateUser(context.Background(), userDtoUpdate)
	if err != nil {
		log.Fatalf("Fail UpdateUser: %v \n", err)
	}
	log.Printf("Before Update >>>>> %v \n", oldUser)
	log.Println("Updating user DONE")
	fmt.Printf("%s\n%s\n%s\n", ">", ">", ">")

	// Get user by email
	log.Println("Get user by email...")
	umail := &wrapperspb.StringValue{Value: "adp10@juventus.com"}
	updatedUser, err := userApi.GetUserByEmail(context.Background(), umail)
	if err != nil {
		log.Fatalf("Fail GetUserByEmail: %v \n", err)
	}
	log.Println(updatedUser)
	log.Println("Get user by email DONE")

}

func sampleUsers() []*protoapi.UserDto {

	var sampleUsers []*protoapi.UserDto

	ronaldo := &protoapi.UserDto{
		Name:  "Cristiano Ronaldo",
		Email: "ronaldo7@yopmail.com",
	}
	messi := &protoapi.UserDto{
		Name:  "Lionel Messi",
		Email: "messi10@yopmail.com",
	}
	delpiero := &protoapi.UserDto{
		Name:  "Del Piero",
		Email: "delpiero10@yopmail.com",
	}
	pogba := &protoapi.UserDto{
		Name:  "Paul Pogba",
		Email: "pogba6@yopmail.com",
	}
	nedved := &protoapi.UserDto{
		Name:  "Pavel Nedved",
		Email: "nedved11@yopmail.com",
	}

	sampleUsers = append(sampleUsers, ronaldo)
	sampleUsers = append(sampleUsers, messi)
	sampleUsers = append(sampleUsers, delpiero)
	sampleUsers = append(sampleUsers, pogba)
	sampleUsers = append(sampleUsers, nedved)

	return sampleUsers
}
