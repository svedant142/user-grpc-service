package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

//unit testing for the repository layer

func TestGetGetUsersByID(t *testing.T) {
	userRepo := NewUserRepo()
	input := [][]int64{{1,2,3},{4,112,2},{111,212},{-1},{}}
	for _, req := range input {
		_, _, err := userRepo.GetUsersByID(context.Background(), req)
		assert.NoError(t, err)
	}
}