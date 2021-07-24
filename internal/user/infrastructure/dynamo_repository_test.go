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
	err = repository.Delete(ctx, userId.String())
	assert.NoError(t, err)
}

func Test_dynamoRepository_Update(t *testing.T) {
	dynamo, err := util.InitDynamo()
	assert.NoError(t, err)
	repository := NewDynamoRepository(dynamo, "Users")
	ctx := context.TODO()
	userId := uuid.New()
	user := models.User{
		Id:       userId.String(),
		Username: "Name",
	}
	err = repository.Store(ctx, &user)
	assert.NoError(t, err)
	modifiedUser := models.User{
		Id:       user.Id,
		Username: "NewName",
	}
	err = repository.Update(ctx, &modifiedUser)
	assert.NoError(t, err)
	assert.Equal(t, "NewName", modifiedUser.Username)
	getResult, err := repository.GetOne(ctx, user.Id)
	assert.NoError(t, err)
	assert.Equal(t, modifiedUser.Username, getResult.Username)
	assert.NotEqual(t, user.Username, getResult.Username)
	err = repository.Delete(ctx, user.Id)
	assert.NoError(t, err)
}

func Test_dynamoRepository_GetAll(t *testing.T) {
	dynamo, err := util.InitDynamo()
	assert.NoError(t, err)
	repository := NewDynamoRepository(dynamo, "Users")
	ctx := context.TODO()

	for i := 0; i < 15; i++ {
		user := models.User{
			Id: uuid.New().String(),
		}
		err = repository.Store(ctx, &user)
		assert.NoError(t, err)
	}

	all, err := repository.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, all, 15)

	for _, user := range all {
		err := repository.Delete(ctx, user.Id)
		assert.NoError(t, err)
	}
}

func Test_dynamoRepository_Delete(t *testing.T) {
	dynamo, err := util.InitDynamo()
	assert.NoError(t, err)
	repository := NewDynamoRepository(dynamo, "Users")
	ctx := context.TODO()

	user1 := models.User{
		Id: uuid.New().String(),
	}
	err = repository.Store(ctx, &user1)
	assert.NoError(t, err)

	user2 := models.User{
		Id: uuid.New().String(),
	}
	err = repository.Store(ctx, &user2)
	assert.NoError(t, err)

	all, err := repository.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, all, 2)

	err = repository.Delete(ctx, user2.Id)
	assert.NoError(t, err)

	result, err := repository.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, result, 1)

	err = repository.Delete(ctx, user1.Id)
	assert.NoError(t, err)
}
