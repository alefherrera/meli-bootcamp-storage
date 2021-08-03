package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"meli-bootcamp-storage/internal/models"
	"meli-bootcamp-storage/internal/user"
	"meli-bootcamp-storage/internal/user/infrastructure"
)

func main() {
	ctx := context.TODO()
	db, err := initDb()
	if err != nil {
		panic(err)
	}

	repository := infrastructure.NewSqlRepository(db)

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

func initDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:mariadb@tcp(127.0.0.1:3306)/mydb")

	if err == nil {
		return db, db.Ping()
	}

	return db, err
}
