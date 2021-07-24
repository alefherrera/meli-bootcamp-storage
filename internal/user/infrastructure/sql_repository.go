package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"meli-bootcamp-storage/internal/models"
	"meli-bootcamp-storage/internal/user"
)

const (
	insertQuery    = "INSERT INTO users (id, firstname, lastname, username, password, email, ip, macAddress, website, image) values (?,?,?,?,?,?,?,?,?,?);"
	selectOneQuery = "SELECT * FROM users WHERE id = ?"
	updateQuery    = "UPDATE users SET firstname = ?, lastname = ?, username = ?, password = ?, email = ?, ip = ?, macAddress = ?, website = ?, image = ? WHERE id = ?"
	selectQuery    = "SELECT * FROM users"
	deleteQuery    = "DELETE from users where id = ?"
)

var _ user.Repository = (*sqlRepository)(nil)

type sqlRepository struct {
	db *sql.DB
}

func NewSqlRepository(db *sql.DB) *sqlRepository {
	return &sqlRepository{db: db}
}

func (receiver *sqlRepository) Store(ctx context.Context, model *models.User) error {
	result, err := receiver.db.ExecContext(ctx, insertQuery, model.Id, model.Firstname, model.Lastname, model.Username, model.Password, model.Email, model.IP, model.MacAddress, model.Website, model.Image)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("%d users created", rowsAffected)
	}

	return nil
}

func (receiver *sqlRepository) GetOne(ctx context.Context, id string) (*models.User, error) {
	result := new(models.User)
	row := receiver.db.QueryRowContext(ctx, selectOneQuery, id)

	err := row.Err()

	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&result.Id,
		&result.Firstname,
		&result.Lastname,
		&result.Username,
		&result.Password,
		&result.Email,
		&result.IP,
		&result.MacAddress,
		&result.Website,
		&result.Image,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (receiver *sqlRepository) Update(ctx context.Context, model *models.User) error {
	result, err := receiver.db.ExecContext(
		ctx,
		updateQuery,
		model.Firstname,
		model.Lastname,
		model.Username,
		model.Password,
		model.Email,
		model.IP,
		model.MacAddress,
		model.Website,
		model.Image,
		model.Id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("%d users updated", rowsAffected)
	}

	return nil
}

func (receiver *sqlRepository) GetAll(ctx context.Context) ([]models.User, error) {
	result := make([]models.User, 0)
	rows, err := receiver.db.QueryContext(ctx, selectQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(
			&user.Id,
			&user.Firstname,
			&user.Lastname,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.IP,
			&user.MacAddress,
			&user.Website,
			&user.Image,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, user)
	}

	return result, nil
}

func (receiver *sqlRepository) Delete(ctx context.Context, id string) error {
	result, err := receiver.db.ExecContext(ctx, deleteQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("%d users deleted", rowsAffected)
	}

	return nil
}
