package grpc

import (
	"context"
	"log"
	"user-service/business/user/grpc/proto"
	userUsecase "user-service/business/user/usecase"
	"user-service/domain/dto/request"
	"user-service/domain/dto/response"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedUserHandlerServer
	usecase userUsecase.IUserUC
}
func NewUserServerGrpc(gserver *grpc.Server, userUCase userUsecase.IUserUC) {
	userServer := &server{
		usecase: userUCase,
	}
	proto.RegisterUserHandlerServer(gserver, userServer)
	reflection.Register(gserver)
}


func (s *server) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserResponse, error){
	userID := in.GetID()
	if err := request.ValidateUserID(userID); err != nil {
		return nil, err
	}
	user, err := s.usecase.GetUserByID(ctx, userID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	res := response.ToUserResponse(user)
	return res, nil
}

func (s *server) GetUsersByIDs(ctx context.Context, in *proto.GetUserListRequest) (*proto.GetUserListResponse, error){
	userIDs := in.GetIDs()
	if err := request.ValidateUserIDs(userIDs); err != nil {
		return nil, err 
	}
	users, invalidIDs, err := s.usecase.GetUsersByIDs(ctx, userIDs)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	res := response.ToUserListResponse(users, invalidIDs)
	return res, nil
}

