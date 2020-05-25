# Go CRUD gRPC

[CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) example to learn how [gRPC](https://grpc.io/) works in [Go](https://golang.org/) with [MongoDB](https://www.mongodb.com/)

## Prerequisite
- **Go**, follow this official [installation](https://golang.org/doc/install) guide
- **MongoDB**, follow this [instruction](https://docs.mongodb.com/manual/installation/) from it's official page
- **gRPC**, follow this [quick start](https://grpc.io/docs/quickstart/go/) guide for Go specific

## How to run gRPC server
1. clone this repository `git clone git@github.com:salprima/gocrud-grpc.git`
2. execute `go mod tidy` to add missing and remove unused modules dependencies
3. adjust [`configs/default-config.yml`](configs/default-config.yml) according to your local environment
4. start gRPC server by executing `make run-grpc-server` 
5. if everything okay, you will see something like the following in your terminal
```sh
$ make run-grpc-server
go run cmd/grpc-server/main.go
2020/05/25 19:13:26 loading config file default-config.yml 
2020/05/25 19:13:26 Starting up GRPC server...
2020/05/25 19:13:26 Creating connection to database...
2020/05/25 19:13:26 Connected to database...
```

## How to run gRPC client
1. follow steps 1-3 on **How to run gRPC server** above
2. make sure gRPC server up and running 
3. run gRPC client by executing `make run-grpc-client`
4. if all good, you will see log in your terminal simulating all gRPC methods as describe in [`api/protobuf/user.proto`](api/protobuf/user.proto)

## What to expect?
- understand how gRPC client & server implementation in Go
- know how MongoDB works within Go using [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)
- a little example of how [Viper](https://github.com/spf13/viper) works to deal with configuration file
- **TIPS**: you can run gRPC client using another language you loved, see available quick start guide [here](https://grpc.io/docs/quickstart/)

## License
This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/), see the [LICENSE](LICENSE) file for details.
