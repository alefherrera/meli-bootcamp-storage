package util

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	txdb.Register("txdb", "mysql", "root:mariadb@/mydb")
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", "identifier")

	if err == nil {
		return db, db.Ping()
	}

	return db, err
}

func InitDbMock() (*sql.DB, sqlmock.Sqlmock, error) {
	return sqlmock.New()
}
