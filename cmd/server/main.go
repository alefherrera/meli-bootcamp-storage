package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"meli-bootcamp-storage/internal/models"
	"meli-bootcamp-storage/internal/user"
	"meli-bootcamp-storage/internal/user/infrastructure"
)

func main() {
	dynamo, err := initDynamo()
	if err != nil {
		panic(err)
	}
	repository := infrastructure.NewDynamoRepository(dynamo, "Users")

	ctx := context.TODO()

	err = storeUsers(ctx, repository)
	if err != nil {
		panic(err)
	}

	err = printUsers(ctx, repository)
	if err != nil {
		panic(err)
	}
}

func printUsers(ctx context.Context, repository user.Repository) error {
	users, err := repository.GetAll(ctx)

	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))

	return nil
}

func storeUsers(ctx context.Context, repository user.Repository) error {
	err := repository.Store(ctx, &models.User{
		Id:        uuid.New().String(),
		Firstname: "Rick",
		Lastname:  "Sanchez",
	})
	if err != nil {
		return err
	}

	err = repository.Store(ctx, &models.User{
		Id:        uuid.New().String(),
		Firstname: "Morty",
		Lastname:  "Smith",
	})

	if err != nil {
		return err
	}

	return nil
}

func initDynamo() (*dynamodb.DynamoDB, error) {
	region := "us-west-2"
	endpoint := "http://localhost:8000"
	cred := credentials.NewStaticCredentials("local", "local", "")
	sess, err := session.NewSession(aws.NewConfig().WithEndpoint(endpoint).WithRegion(region).WithCredentials(cred))
	if err != nil {
		return nil, err
	}
	dynamo := dynamodb.New(sess)
	return dynamo, nil
}
