package user_test

import (
	media "pop_culture/api/resource/Media"
	mediatype "pop_culture/api/resource/MediaType"
	user "pop_culture/api/resource/User"
	mockDB "pop_culture/mock/db"
	testUtil "pop_culture/util/test"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestRepository_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDb()
	testUtil.NoError(t, err)
	repo := user.NewRepository(db)

	name := "Name"
	email := "Email"
	password := "Password"
	id := uuid.New()

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `users` ").
		WithArgs(id, name, email, password, mockDB.AnyTime{}, mockDB.AnyTime{}, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	user := &user.User{ID: id,
		Name:      name,
		Email:     &email,
		Password:  password,
		Medias:    []media.Media{},
		Interests: []mediatype.TypeMedia{},
	}
	_, err = repo.Create(user)

	testUtil.NoError(t, err)

}

func TestRepository_Read(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDb()
	testUtil.NoError(t, err)

	repo := user.NewRepository(db)

	id := uuid.New()
	mockRows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(id, "Miguel1", "miguel1@test.com", "987654321")

	mock.ExpectQuery("^SELECT (.+) FROM `users` WHERE (.+)").
		WithArgs(id, 1).
		WillReturnRows(mockRows)

	user, err := repo.Read(id)
	testUtil.NoError(t, err)
	testUtil.Equal(t, "Miguel1", user.Name)

}

func TestRepository_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDb()
	testUtil.NoError(t, err)

	repo := user.NewRepository(db)

	email := "Email"
	id := uuid.New()

	_ = sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(id, "Miguel1", "miguel@test.com", "987654321")

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `users` SET").
		WithArgs("Name", email, mockDB.AnyTime{}, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	user := &user.User{
		ID:        id,
		Name:      "Name",
		Email:     &email,
		Password:  "Password",
		Medias:    []media.Media{},
		Interests: []mediatype.TypeMedia{},
	}

	rows, err := repo.Update(user)
	testUtil.NoError(t, err)
	testUtil.Equal(t, 1, rows)

}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDb()
	testUtil.NoError(t, err)

	repo := user.NewRepository(db)

	id := uuid.New()

	_ = sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(id, "Miguel1", "miguel@test.com", "987654321")

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `users` SET `deleted_at`").
		WithArgs(mockDB.AnyTime{}, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()
	rows, err := repo.Delete(id)
	testUtil.NoError(t, err)
	testUtil.Equal(t, 1, rows)

}
