package user

import (
	"context"
	"meli-bootcamp-storage/internal/models"
)

type Repository interface {
	Store(ctx context.Context, model *models.User) error
	GetOne(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, model *models.User) error
	GetAll(ctx context.Context) ([]models.User, error)
	Delete(ctx context.Context, id string) error
}
