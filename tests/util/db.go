package util

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func init() {
	txdb.Register("txdb", "mysql", "root:mariadb@/mydb")
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", uuid.New().String())

	if err == nil {
		return db, db.Ping()
	}

	return db, err
}

func InitDbMock() (*sql.DB, sqlmock.Sqlmock, error) {
	return sqlmock.New()
}

func InitDynamo() (*dynamodb.DynamoDB, error) {
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
