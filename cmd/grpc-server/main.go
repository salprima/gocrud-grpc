package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/salprima/gocrud-grpc/internal/protoapi"
	"github.com/salprima/gocrud-grpc/internal/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {

	args := os.Args[1:]
	var configname string = "default-config"
	if len(args) > 0 {
		configname = args[0] + "-config"
	}
	log.Printf("loading config file %s.yml \n", configname)

	viper.SetConfigName(configname)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s.yml \n", err))
	}

}

func main() {
	server := grpc.NewServer()

	protoapi.RegisterUserApiServer(server, &service.UserSvc{})

	port := ":" + viper.GetString("app.grpc.port")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", port, err)
	}

	panic(server.Serve(listener))
}
