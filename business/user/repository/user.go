package repository

import (
	"context"
	"errors"
	"fmt"
	"user-service/domain/entity"
)

type IUserRepo interface {
	GetUsersByID(ctx context.Context, userIDs []int64) ([]*entity.User, []int64, error)
}

type UserRepo struct{
	//In the UserMap userID acts as key , while mapping the database
	UserMap map[uint64](*entity.User)
}

//mocking the database 
func MockDatabase() (map[uint64](*entity.User)) {
	userMap := make(map[uint64](*entity.User))

	// Adding 10 user records to the map
	for i := 1; i <= 10; i++ {
		userMap[uint64(i)] = &entity.User{
			ID:      int64(i),
			Fname:   fmt.Sprintf("User No. %d", i),
			City:     fmt.Sprintf("City No. %d", i),
			Phone:   int64(1000000000 + i),
			Height:  5.5 + float32(i)/10,
			Married: i%2 == 0,
		}
	}
	userMap[200000000000] = &entity.User{
		ID:      200000000000,
		Fname:   fmt.Sprintf("User No. %d", 200000000000),
		City:     fmt.Sprintf("City No. %d", 200000000000),
		Phone:   int64(1000000000 + 2),
		Height:  5.5 + float32(1)/10,
		Married: 200000000000%2 == 0,
	}
	return userMap
}

func NewUserRepo() IUserRepo {
	userMap := MockDatabase()
	return &UserRepo{UserMap: userMap}
}

func (urepo *UserRepo) GetUsersByID(ctx context.Context, userIDs []int64) ([]*entity.User, []int64, error) {
	var validUsers []*entity.User
	var invalidUserIds []int64
	if urepo.UserMap == nil {
		return nil, nil, errors.New("unable to connect with database records")
	}
	for _, id := range userIDs {
		if _, ok := urepo.UserMap[uint64(id)] ; !ok {
			invalidUserIds = append(invalidUserIds, id)
			continue
		}
		validUsers = append(validUsers, urepo.UserMap[uint64(id)])
	}
	return validUsers, invalidUserIds, nil
}