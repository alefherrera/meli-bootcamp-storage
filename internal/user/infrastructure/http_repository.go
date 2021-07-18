package infrastructure

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"meli-bootcamp-storage/internal/models"
	"meli-bootcamp-storage/internal/user"
	"net/http"
)

var _ user.Repository = (*httpRepository)(nil)

type httpRepository struct {
	client *http.Client
}

type responseBody struct {
	Status string        `json:"status"`
	Code   int           `json:"code"`
	Total  int           `json:"total"`
	Data   []models.User `json:"data"`
}

func NewHttpRepository(client *http.Client) *httpRepository {
	return &httpRepository{client: client}
}

func (receiver *httpRepository) Store(ctx context.Context, model *models.User) error {
	panic("implement me")
}

func (receiver *httpRepository) GetOne(ctx context.Context) (*models.User, error) {
	panic("implement me")
}

func (receiver *httpRepository) Update(ctx context.Context, model *models.User) error {
	panic("implement me")
}

func (receiver *httpRepository) GetAll(ctx context.Context) ([]models.User, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://fakerapi.it/api/v1/users?_quantity=10", nil)

	if err != nil {
		return nil, err
	}

	response, err := receiver.client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Error getting users: %s", string(bytes)))
	}

	var result responseBody
	err = json.Unmarshal(bytes, &result)

	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (receiver *httpRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}
