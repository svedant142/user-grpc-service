package response

import (
	"user-service/business/user/grpc/proto"
	"user-service/domain/entity"
)

func ToUserResponse(user *entity.User) *proto.GetUserResponse {
	if user == nil {
		return &proto.GetUserResponse{
			Usermessage: &proto.GetUserResponse_Message{
				Message: "user not found",
			},
			Status: true,
		}
	} 
		return &proto.GetUserResponse{
			Usermessage: &proto.GetUserResponse_User{
				User: &proto.User{
					ID:      int64(user.ID),
					Fname:   user.Fname,
					City:    user.City,
					Phone:   user.Phone,
					Height:  float32(user.Height),
					Married: user.Married,
				},
			},
			Status: true,
		}
}

func ToUserListResponse(users []*entity.User, invalidIDs []int64) *proto.GetUserListResponse {
	successResponse := &proto.SuccessListResponse{
		Users:       []*proto.User{},
		InvalidIDs: invalidIDs,
		Status:     true,
	}
	for _, u := range users {
		if u == nil {
			return &proto.GetUserListResponse{
				Response: &proto.GetUserListResponse_ErrorResponse{
					ErrorResponse: &proto.ErrorResponse{
						Error: "user not found, empty user",
						Status: false,
					},
				},
			}
		}
		protoUser := &proto.User{
				ID:      int64(u.ID),
				Fname:   u.Fname,
				City:    u.City,
				Phone:   u.Phone,
				Height:  u.Height,
				Married: u.Married,
		}
		successResponse.Users = append(successResponse.Users, protoUser)
	}
	return &proto.GetUserListResponse{
		Response: &proto.GetUserListResponse_SuccessListResponse{
				SuccessListResponse: successResponse,
		},
	}
}