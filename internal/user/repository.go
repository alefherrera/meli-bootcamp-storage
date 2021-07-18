package user

import (
	"context"
	"github.com/google/uuid"
	"meli-bootcamp-storage/internal/models"
)

type Repository interface {
	Store(ctx context.Context, model *models.User) error
	GetOne(ctx context.Context) (*models.User, error)
	Update(ctx context.Context, model *models.User) error
	GetAll(ctx context.Context) ([]models.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
