package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/salprima/gocrud-grpc/internal/protoapi"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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

	userDto := &protoapi.UserDto{
		Name:  "John Doe",
		Email: "johndoe@yopmail.com",
	}

	user, err := userApi.CreateUser(context.Background(), userDto)
	if err != nil {
		log.Fatalf("Fail CreateUser: %v", err)
	}
	log.Println(user)

	userID := &wrapperspb.StringValue{Value: user.Id}
	newUser, err := userApi.GetUserByID(context.Background(), userID)
	if err != nil {
		log.Fatalf("Fail GetUserByID: %v", err)
	}
	log.Println(newUser)
}
