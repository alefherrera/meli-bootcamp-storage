package infrastructure

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"meli-bootcamp-storage/internal/models"
	"meli-bootcamp-storage/internal/user"
)

var _ user.Repository = (*dynamoRepository)(nil)

type dynamoRepository struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

func NewDynamoRepository(dynamo *dynamodb.DynamoDB, table string) *dynamoRepository {
	return &dynamoRepository{
		dynamo: dynamo,
		table:  table,
	}
}

func (receiver *dynamoRepository) Store(ctx context.Context, model *models.User) error {
	av, err := dynamodbattribute.MarshalMap(model)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(receiver.table),
	}

	_, err = receiver.dynamo.PutItemWithContext(ctx, input)

	if err != nil {
		return err
	}

	return nil
}

func (receiver *dynamoRepository) GetOne(ctx context.Context, id string) (*models.User, error) {
	result, err := receiver.dynamo.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(receiver.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	return itemToUser(result.Item)
}

func (receiver *dynamoRepository) Update(ctx context.Context, model *models.User) error {
	return receiver.Store(ctx, model)
}

func (receiver *dynamoRepository) GetAll(ctx context.Context) ([]models.User, error) {
	result, err := receiver.dynamo.ScanWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(receiver.table),
	})

	if err != nil {
		return nil, err
	}

	output := make([]models.User, len(result.Items))

	for i, item := range result.Items {

		user, err := itemToUser(item)

		if err != nil {
			return nil, err
		}

		output[i] = *user
	}

	return output, nil
}

func (receiver *dynamoRepository) Delete(ctx context.Context, id string) error {
	_, err := receiver.dynamo.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(receiver.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func itemToUser(input map[string]*dynamodb.AttributeValue) (*models.User, error) {
	var item models.User
	err := dynamodbattribute.UnmarshalMap(input, &item)

	if err != nil {
		return nil, err
	}

	return &item, nil
}
