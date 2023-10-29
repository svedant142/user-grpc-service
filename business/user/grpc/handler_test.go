package grpc

import (
	"context"
	"log"
	"net"
	"testing"
	"time"
	"user-service/business/user/grpc/proto"
	userRepo "user-service/business/user/repository"
	userUC "user-service/business/user/usecase"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

// Integration Tests for both the endpoints with valid and invalid cases handled as shown in testcase

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	urepo := userRepo.NewUserRepo()
	uusecase := userUC.NewUserUC(urepo)
	NewUserServerGrpc(s, uusecase)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(string, time.Duration) (net.Conn, error) {
	return lis.Dial()
}

func TestIntegrationGetUser(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	var NormalTestCases [2]*proto.GetUserRequest
	client := proto.NewUserHandlerClient(conn)
	NormalTestCases[0] = &proto.GetUserRequest{
		ID: 1, 
	}
	NormalTestCases[1] = &proto.GetUserRequest{
		ID: 100, 
	}
	for _, getUserRequest := range NormalTestCases {
		resp, err := client.GetUser(context.Background(), getUserRequest)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	}

	var ErroredTestCases [2]*proto.GetUserRequest
	ErroredTestCases[0] = &proto.GetUserRequest{
		ID: 0, 
	}
	ErroredTestCases[1] = &proto.GetUserRequest{
		ID: -1, 
	}
	for _, getUserRequest := range ErroredTestCases {
		_, err := client.GetUser(context.Background(), getUserRequest)
		assert.Error(t, err)
	}
}

func TestIntegrationGetUsersByIDs(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	var NormalTestCases [3]*proto.GetUserListRequest
	client := proto.NewUserHandlerClient(conn)
	NormalTestCases[0] = &proto.GetUserListRequest{
		IDs: []int64{1,2,3,4,5},
	}
	NormalTestCases[1] = &proto.GetUserListRequest{
		IDs: []int64{100,1,200}, 
	}
	NormalTestCases[2] = &proto.GetUserListRequest{
		IDs: []int64{100,1000,200}, 
	}
	for _, getUserRequest := range NormalTestCases {
		resp, err := client.GetUsersByIDs(context.Background(), getUserRequest)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	}

	var ErroredTestCases [2]*proto.GetUserListRequest
	ErroredTestCases[0] = &proto.GetUserListRequest{
		IDs: []int64{},  
	}
	ErroredTestCases[1] = &proto.GetUserListRequest{
		IDs: []int64{0},  
	}
	ErroredTestCases[1] = &proto.GetUserListRequest{
		IDs: []int64{1,-1,2},  
	}
	for _, getUserRequest := range ErroredTestCases {
		_, err := client.GetUsersByIDs(context.Background(), getUserRequest)
		assert.Error(t, err)
	}
}