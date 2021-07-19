package infrastructure

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"meli-bootcamp-storage/internal/models"
	"meli-bootcamp-storage/tests/util"
	"testing"
)

func Test_sqlRepository_Store(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)
	repository := NewSqlRepository(db)
	ctx := context.TODO()
	userId := uuid.New()
	user := models.User{
		UUID: userId,
	}
	err = repository.Store(ctx, &user)
	assert.NoError(t, err)
	getResult, err := repository.GetOne(ctx, uuid.New())
	assert.NoError(t, err)
	assert.Nil(t, getResult)
	getResult, err = repository.GetOne(ctx, userId)
	assert.NoError(t, err)
	assert.NotNil(t, getResult)
	assert.Equal(t, user.UUID, getResult.UUID)
}

func Test_sqlRepository_Store_Mock(t *testing.T) {
	db, mock, err := util.InitDbMock()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))
	columns := []string{"uuid", "firstname", "lastname", "username", "password", "email", "ip", "macAddress", "website", "image"}
	rows := sqlmock.NewRows(columns)
	userId := uuid.New()
	userId2 := uuid.New()
	rows.AddRow(userId, "", "", "", "", "", "", "", "", "")
	rows2 := sqlmock.NewRows(columns)
	mock.ExpectQuery("SELECT .* FROM users").WithArgs(userId2).WillReturnRows(rows2)
	mock.ExpectQuery("SELECT .* FROM users").WithArgs(userId).WillReturnRows(rows)
	repository := NewSqlRepository(db)
	ctx := context.TODO()
	user := models.User{
		UUID: userId,
	}
	err = repository.Store(ctx, &user)
	assert.NoError(t, err)
	getResult, err := repository.GetOne(ctx, userId2)
	assert.NoError(t, err)
	assert.Nil(t, getResult)
	getResult, err = repository.GetOne(ctx, userId)
	assert.NoError(t, err)
	assert.NotNil(t, getResult)
	assert.Equal(t, user.UUID, getResult.UUID)
}

func Test_sqlRepository_Update(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)
	repository := NewSqlRepository(db)
	ctx := context.TODO()
	userId := uuid.New()
	user := models.User{
		UUID:     userId,
		Username: "Name",
	}
	err = repository.Store(ctx, &user)
	assert.NoError(t, err)
	modifiedUser := models.User{
		UUID:     user.UUID,
		Username: "NewName",
	}
	err = repository.Update(ctx, &modifiedUser)
	assert.NoError(t, err)
	assert.Equal(t, "NewName", modifiedUser.Username)
	getResult, err := repository.GetOne(ctx, user.UUID)
	assert.NoError(t, err)
	assert.Equal(t, modifiedUser.Username, getResult.Username)
	assert.NotEqual(t, user.Username, getResult.Username)
}

func Test_sqlRepository_GetAll(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)
	repository := NewSqlRepository(db)
	ctx := context.TODO()

	for i := 0; i < 15; i++ {
		user := models.User{
			UUID: uuid.New(),
		}
		err = repository.Store(ctx, &user)
		assert.NoError(t, err)
	}

	all, err := repository.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, all, 15)
}

func Test_sqlRepository_Delete(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)
	repository := NewSqlRepository(db)
	ctx := context.TODO()

	user1 := models.User{
		UUID: uuid.New(),
	}
	err = repository.Store(ctx, &user1)
	assert.NoError(t, err)

	user2 := models.User{
		UUID: uuid.New(),
	}
	err = repository.Store(ctx, &user2)
	assert.NoError(t, err)

	all, err := repository.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, all, 2)

	err = repository.Delete(ctx, user2.UUID)
	assert.NoError(t, err)

	result, err := repository.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
}
