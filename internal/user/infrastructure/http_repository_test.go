package infrastructure

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_httpRepository_GetAll(t *testing.T) {
	repository := NewHttpRepository(http.DefaultClient)
	users, err := repository.GetAll(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, users, 10)
}

func Test_httpRepository_GetAll_HttpMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://fakerapi.it/api/v1/users?_quantity=10",
		httpmock.NewStringResponder(200, "{\n  \"status\": \"OK\",\n  \"code\": 200,\n  \"total\": 10,\n  \"data\": [\n    {\n      \"uuid\": \"43ec8f04-7568-3907-8efb-351efde29df1\",\n      \"firstname\": \"Maci\",\n      \"lastname\": \"Zemlak\",\n      \"username\": \"heathcote.carson\",\n      \"password\": \"tWa5KUBC9`cSJ&DBo\",\n      \"email\": \"norbert.hand@carroll.com\",\n      \"ip\": \"215.97.228.128\",\n      \"macAddress\": \"92:00:6C:65:43:72\",\n      \"website\": \"http:\\/\\/satterfield.com\",\n      \"image\": \"http:\\/\\/placeimg.com\\/640\\/480\\/people\"\n    },\n    {\n      \"uuid\": \"d5026588-5a3e-326f-926b-05d30c29a62e\",\n      \"firstname\": \"Destin\",\n      \"lastname\": \"Labadie\",\n      \"username\": \"vonrueden.estefania\",\n      \"password\": \"c-H(9!CWBgl^\",\n      \"email\": \"cfriesen@gmail.com\",\n      \"ip\": \"175.61.89.83\",\n      \"macAddress\": \"1A:F4:1C:79:84:0F\",\n      \"website\": \"http:\\/\\/schultz.com\",\n      \"image\": \"http:\\/\\/placeimg.com\\/640\\/480\\/people\"\n    }\n  ]\n}"))

	repository := NewHttpRepository(http.DefaultClient)
	users, err := repository.GetAll(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, httpmock.GetTotalCallCount(), 1)
}
