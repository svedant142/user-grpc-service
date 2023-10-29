package usecase

import (
	"context"
	"testing"
	"user-service/domain/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//unit testing for the usecase layer

type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) GetUsersByID(ctx context.Context, userIDs []int64) ([]*entity.User, []int64, error) {
	return nil, nil, nil
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(mockUserRepo)
	mockRepo.On("GetUsersByID", mock.Anything, mock.Anything).Return(nil, nil, nil)

	userUC := NewUserUC(mockRepo)
	input := []int64{1,2,3}
	for _, id := range input {
		_, err := userUC.GetUserByID(context.Background(), id)
		assert.NoError(t, err)
	}
}

func TestGetUsersByIDs(t *testing.T) {
	mockRepo := new(mockUserRepo)
	mockRepo.On("GetUsersByID", mock.Anything, mock.Anything).Return(nil, nil)
	userUC := NewUserUC(mockRepo)
	input := [][]int64{{1,2,3},{4,5}}
	for _, id := range input {
		_, _, err := userUC.GetUsersByIDs(context.Background(), id)
		assert.NoError(t, err)

	}
}
