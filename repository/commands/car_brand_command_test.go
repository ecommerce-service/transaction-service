package commands

import (
	"booking-car/domain/models"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

type CarBrand struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

var (
	fakeData CarBrand
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCarBrandAdd(t *testing.T) {
	db, mock := NewMock()
	now := time.Now().UTC()
	model := models.NewCarBrandModel().SetName(gofakeit.Name()).SetCreatedAt(now).SetUpdatedAt(now)
	id := gofakeit.UUID()
	cmd := NewCarBrandCommand(db, model)

	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `INSERT INTO car_brands (name,created_at,updated_at) VALUES($1,$2,$3) RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.Name(), model.CreatedAt(), model.UpdatedAt()).WillReturnRows(rows)
	res, err := cmd.Add()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestAddCarBrandError(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	model := models.NewCarBrandModel().SetUpdatedAt(now).SetCreatedAt(now)
	cmd := NewCarBrandCommand(db, model)

	defer func() {
		db.Close()
	}()

	statement := `INSERT INTO car_brands (name,created_at,updated_at) VALUES($1,$2,$3) RETURNING id`
	mock.ExpectQuery(statement).WithArgs("", model.CreatedAt(), model.UpdatedAt()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Add()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestEditCarBrand(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewCarBrandModel().SetName(gofakeit.Name()).SetCreatedAt(now).SetUpdatedAt(now).SetId(id)
	cmd := NewCarBrandCommand(db, model)

	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE car_brands SET name=$1,updated_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.Name(), model.UpdatedAt(), model.Id()).WillReturnRows(rows)
	res, err := cmd.Edit()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestEditCarBrandError(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	model := models.NewCarBrandModel().SetCreatedAt(now).SetUpdatedAt(now)
	cmd := NewCarBrandCommand(db, model)

	defer func() {
		db.Close()
	}()

	statement := `UPDATE car_brands SET name=$1,updated_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs("", model.UpdatedAt(), model.Id()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Edit()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestCarBrandDelete(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewCarBrandModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)
	cmd := NewCarBrandCommand(db, model)

	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE car_brands SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(), model.DeletedAt().Time, model.Id()).WillReturnRows(rows)
	res, err := cmd.Delete()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteCarBrandError(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	model := models.NewCarBrandModel().SetUpdatedAt(now).SetDeletedAt(now)
	cmd := NewCarBrandCommand(db, model)

	defer func() {
		db.Close()
	}()

	statement := `UPDATE car_brands SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(), model.DeletedAt().Time, model.Id()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Delete()

	assert.Error(t, err)
	assert.Empty(t, res)
}
