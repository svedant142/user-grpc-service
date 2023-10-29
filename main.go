package main

import (
	"fmt"
	"net"
	userGrpc "user-service/business/user/grpc"
	userRepo "user-service/business/user/repository"
	userUC "user-service/business/user/usecase"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("user-grpc-service start")

	urepo := userRepo.NewUserRepo()
	uusecase := userUC.NewUserUC(urepo)
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("connection error", err)
		return
	}
	
	//grpc initialization
	server := grpc.NewServer()
	userGrpc.NewUserServerGrpc(server, uusecase)
	fmt.Println("Server Running at ", ":8080")

	err = server.Serve(list)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}	
}