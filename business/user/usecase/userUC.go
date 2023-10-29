package usecase

import (
	"context"
	userRepo "user-service/business/user/repository"
	"user-service/domain/entity"
)

type IUserUC interface {	
	GetUserByID(ctx context.Context, userID int64) (*entity.User, error)
	GetUsersByIDs(ctx context.Context, userIDs []int64) ([]*entity.User, []int64, error)
}

type UserUC struct {
	repo userRepo.IUserRepo
}

func NewUserUC(userrepo userRepo.IUserRepo) IUserUC {
	return &UserUC{repo: userrepo}
}

func (uUc *UserUC) GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	users, invalidIDs, err := uUc.repo.GetUsersByID(ctx, []int64{userID})
	if err != nil {
		return nil, err
	}
	if len(users)==0 || len(invalidIDs) == 1 {
		return nil, nil
	}
	return users[0], nil
}

func (uUc *UserUC) GetUsersByIDs(ctx context.Context, userIDs []int64) ([]*entity.User, []int64, error) {
	return uUc.repo.GetUsersByID(ctx, userIDs)
}