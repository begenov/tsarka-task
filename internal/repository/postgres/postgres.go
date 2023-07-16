package postgres

import (
	"context"
	"database/sql"

	"github.com/begenov/tsarka-task/internal/domain"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user domain.User) (int, error) {
	stmt := `INSERT INTO "user" (first_name, last_name) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRowContext(ctx, stmt, user.FirstName, user.LastName).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepo) GetUser(ctx context.Context, id int) (domain.User, error) {

	stmt := `SELECT * FROM "user" WHERE id = $1`
	row := r.db.QueryRow(stmt, id)
	var user domain.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	stmt := `UPDATE "user" SET first_name = $1, last_name = $2 WHERE id = $3`
	_, err := r.db.Exec(stmt, user.FirstName, user.LastName, user.ID)
	if err != nil {
		return domain.User{}, err
	}

	stmt = `SELECT first_name, last_name FROM "user" WHERE id = $1`
	row := r.db.QueryRow(stmt, user.ID)

	updatedUser := domain.User{
		ID: user.ID,
	}

	err = row.Scan(&updatedUser.FirstName, &updatedUser.LastName)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, id int) error {
	stmt := `DELETE FROM "user" WHERE id = $1`
	_, err := r.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
