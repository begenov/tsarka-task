package postgres

import (
	"context"
	"database/sql"
	"testing"

	"github.com/begenov/tsarka-task/internal/domain"

	"github.com/begenov/tsarka-task/pkg/util"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestUserRepo_CreateUser(t *testing.T) {
	db := setupTestDB(t)

	repo := NewUserRepo(db)

	user := domain.User{
		FirstName: util.RandomString(6),
		LastName:  util.RandomString(6),
	}

	id, err := repo.CreateUser(context.Background(), user)

	require.NoError(t, err)
	require.Greater(t, id, 0)

	err = repo.DeleteUser(context.Background(), id)
	require.NoError(t, err)
}

func TestUserRepo_GetUser(t *testing.T) {
	db := setupTestDB(t)

	repo := NewUserRepo(db)

	user := domain.User{
		FirstName: util.RandomString(6),
		LastName:  util.RandomString(6),
	}

	id, err := repo.CreateUser(context.Background(), user)
	require.NoError(t, err)

	resultUser, err := repo.GetUser(context.Background(), id)

	require.NoError(t, err)
	require.Equal(t, user.FirstName, resultUser.FirstName)
	require.Equal(t, user.LastName, resultUser.LastName)

	err = repo.DeleteUser(context.Background(), id)
	require.NoError(t, err)
}

func TestUserRepo_UpdateUser(t *testing.T) {
	db := setupTestDB(t)

	repo := NewUserRepo(db)

	user := domain.User{
		FirstName: util.RandomString(6),
		LastName:  util.RandomString(6),
	}

	id, err := repo.CreateUser(context.Background(), user)
	require.NoError(t, err)

	user.FirstName = util.RandomString(6)
	user.LastName = util.RandomString(6)
	user.ID = id

	updatedUser, err := repo.UpdateUser(context.Background(), user)

	require.NoError(t, err)
	require.Equal(t, user.FirstName, updatedUser.FirstName)
	require.Equal(t, user.LastName, updatedUser.LastName)

	err = repo.DeleteUser(context.Background(), id)
	require.NoError(t, err)
}

func TestUserRepo_DeleteUser(t *testing.T) {
	db := setupTestDB(t)

	repo := NewUserRepo(db)

	user := domain.User{
		FirstName: util.RandomString(6),
		LastName:  util.RandomString(6),
	}

	id, err := repo.CreateUser(context.Background(), user)
	require.NoError(t, err)

	err = repo.DeleteUser(context.Background(), id)

	require.NoError(t, err)

	_, err = repo.GetUser(context.Background(), id)

	require.Error(t, err)
}

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=root password=root dbname=users sslmode=disable")
	require.NoError(t, err)
	return db
}
