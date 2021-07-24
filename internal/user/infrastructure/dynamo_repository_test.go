package infrastructure

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meli-bootcamp-storage/internal/models"
	"meli-bootcamp-storage/tests/util"
	"testing"
)

func Test_dynamoRepository_Store(t *testing.T) {
	dynamo, err := util.InitDynamo()
	assert.NoError(t, err)
	repository := NewDynamoRepository(dynamo, "Users")
	ctx := context.TODO()
	userId := uuid.New()
	user := models.User{
		Id: userId.String(),
	}
	err = repository.Store(ctx, &user)
	assert.NoError(t, err)
	getResult, err := repository.GetOne(ctx, uuid.New().String())
	assert.NoError(t, err)
	assert.Nil(t, getResult)
	getResult, err = repository.GetOne(ctx, userId.String())
	assert.NoError(t, err)
	assert.NotNil(t, getResult)
	assert.Equal(t, user.Id, getResult.Id)
}
