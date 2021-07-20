package queries

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

func NewSqlMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCarBrandBrowse(t *testing.T) {
	db, sqlMock := NewSqlMock()
	defer func() {
		db.Close()
	}()
	search := "%%"
	now := time.Now().UTC()
	model := models.NewCarBrandModel().SetId(gofakeit.UUID()).SetName(gofakeit.Car().Brand).SetCreatedAt(now).SetUpdatedAt(now)
	repository := NewCarBrandQuery(db)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).AddRow(model.Id(), model.Name(), model.CreatedAt(), model.UpdatedAt())
	statement := `SELECT id,name,created_at,updated_at FROM car_brands WHERE deleted_at IS NULL AND LOWER(name) LIKE $1 ORDER BY updated_at ASC LIMIT $2 OFFSET $3`
	sqlMock.ExpectQuery(statement).WithArgs(search, 10, 0).WillReturnRows(rows)

	res, err := repository.Browse("", "updated_at", "ASC", 10, 0)
	data := res.([]*models.CarBrands)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Len(t, data, 1)
}

func TestCarBrandBrowseAll(t *testing.T) {
	db, sqlMock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	now := time.Now().UTC()
	model := models.NewCarBrandModel().SetId(gofakeit.UUID()).SetName(gofakeit.Car().Brand).SetCreatedAt(now).SetUpdatedAt(now)
	repository := NewCarBrandQuery(db)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).AddRow(model.Id(), model.Name(), model.CreatedAt(), model.UpdatedAt())
	statement := `SELECT id,name,created_at,updated_at FROM car_brands WHERE deleted_at IS NULL AND LOWER(name) LIKE $1`
	sqlMock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)

	res, err := repository.BrowseAll("")
	data := res.([]*models.CarBrands)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Len(t, data, 1)
}

func TestCarBrandReadBy(t *testing.T) {
	db, sqlMock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewCarBrandModel().SetId(id).SetName(gofakeit.Car().Brand).SetCreatedAt(now).SetUpdatedAt(now)
	repository := NewCarBrandQuery(db)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).AddRow(model.Id(), model.Name(), model.CreatedAt(), model.UpdatedAt())
	statement := `SELECT id,name,created_at,updated_at FROM car_brands WHERE deleted_at IS NULL AND id=$1`
	sqlMock.ExpectQuery(statement).WithArgs(id).WillReturnRows(rows)

	res, err := repository.ReadBy("id", "=", id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarBrandCount(t *testing.T) {
	db, sqlMock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	repository := NewCarBrandQuery(db)
	count := 1

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_brands WHERE deleted_at IS NULL AND LOWER(name) LIKE $1`
	sqlMock.ExpectQuery(statement).WithArgs("%%").WillReturnRows(rows)

	res, err := repository.Count("")

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarBrandCountByWithOutId(t *testing.T) {
	db, sqlMock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	repository := NewCarBrandQuery(db)
	count := 1
	value := gofakeit.Name()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_brands WHERE deleted_at IS NULL AND name=$1`
	sqlMock.ExpectQuery(statement).WithArgs(value).WillReturnRows(rows)

	res, err := repository.CountBy("name", "=", "", value)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarBrandCountByWithId(t *testing.T) {
	db, sqlMock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	repository := NewCarBrandQuery(db)
	count := 1
	value := gofakeit.Name()
	id := gofakeit.UUID()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_brands WHERE deleted_at IS NULL AND name=$1 AND id<>$2`
	sqlMock.ExpectQuery(statement).WithArgs(value, id).WillReturnRows(rows)

	res, err := repository.CountBy("name", "=", id, value)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarBrandCountByWithIdError(t *testing.T) {
	db, sqlMock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	repository := NewCarBrandQuery(db)
	count := 1
	value := gofakeit.Name()
	id := gofakeit.UUID()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_brands WHERE deleted_at IS NULL AND name=$1 AND id<>$2`
	sqlMock.ExpectQuery(statement).WithArgs(value, id).WillReturnRows(rows)

	res, err := repository.CountBy("name", "=", "", value)

	assert.Error(t, err)
	assert.Empty(t, res)
}
