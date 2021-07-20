package queries

import (
	"booking-car/domain/models"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCarTypeQuery_Browse(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	carType := models.NewCarTypeModel().SetBrandId(gofakeit.UUID()).SetId(gofakeit.UUID()).SetName(gofakeit.CarType()).SetCreatedAt(now).SetUpdatedAt(now)
	repository := NewCarTypeQuery(db)

	rows := sqlmock.NewRows([]string{"id", "brand_id", "name", "created_at", "updated_at"}).AddRow(carType.Id(), carType.BrandId(), carType.Name(), carType.CreatedAt(), carType.UpdatedAt())
	statement := `SELECT id,brand_id,name,created_at,updated_at FROM car_types WHERE deleted_at IS NULL AND LOWER(name) LIKE $1 ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(search, 10, 0).WillReturnRows(rows)
	res, err := repository.Browse("", "created_at", "desc", 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarTypeQuery_BrowseNoResultNoError(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	repository := NewCarTypeQuery(db)

	statement := `SELECT id,brand_id,name,created_at,updated_at FROM car_types WHERE deleted_at IS NULL AND LOWER(name) LIKE $1 ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(search, 10, 0).WillReturnError(sql.ErrNoRows)
	res, err := repository.Browse("", "created_at", "desc", 10, 0)

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestCarTypeQuery_BrowseAll(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	brandId := gofakeit.UUID()
	carType := models.NewCarTypeModel().SetBrandId(gofakeit.UUID()).SetId(gofakeit.UUID()).SetName(gofakeit.CarType()).SetCreatedAt(now).SetUpdatedAt(now)
	repository := NewCarTypeQuery(db)

	rows := sqlmock.NewRows([]string{"id", "brand_id", "name", "created_at", "updated_at"}).AddRow(carType.Id(), carType.BrandId(), carType.Name(), carType.CreatedAt(), carType.UpdatedAt())
	statement := `SELECT id,brand_id,name,created_at,updated_at FROM car_types WHERE deleted_at IS NULL AND LOWER(name) LIKE $1 AND brand_id=$2`
	mock.ExpectQuery(statement).WithArgs(search, brandId).WillReturnRows(rows)
	res, err := repository.BrowseAll("", brandId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarTypeQuery_BrowseAllNeResultNoError(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	brandId := gofakeit.UUID()
	repository := NewCarTypeQuery(db)

	rows := sqlmock.NewRows([]string{"id", "brand_id", "name", "created_at", "updated_at"})
	statement := `SELECT id,brand_id,name,created_at,updated_at FROM car_types WHERE deleted_at IS NULL AND LOWER(name) LIKE $1 AND brand_id=$2`
	mock.ExpectQuery(statement).WithArgs(search, brandId).WillReturnRows(rows)
	res, err := repository.BrowseAll("", brandId)

	assert.NoError(t, err)
	assert.Empty(t, res)
}

func TestCarTypeQuery_ReadBy(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	carType := models.NewCarTypeModel().SetBrandId(gofakeit.UUID()).SetId(gofakeit.UUID()).SetName(gofakeit.CarType()).SetCreatedAt(now).SetUpdatedAt(now).SetId(id)
	repository := NewCarTypeQuery(db)

	rows := sqlmock.NewRows([]string{"id", "brand_id", "name", "created_at", "updated_at"}).AddRow(carType.Id(), carType.BrandId(), carType.Name(), carType.CreatedAt(), carType.UpdatedAt())
	statement := `SELECT id,brand_id,name,created_at,updated_at FROM car_types WHERE deleted_at IS NULL AND id=$1`
	mock.ExpectQuery(statement).WithArgs(id).WillReturnRows(rows)
	res, err := repository.ReadBy("id", "=", id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarTypeQuery_ReadByNoResult(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	id := gofakeit.UUID()
	repository := NewCarTypeQuery(db)

	statement := `SELECT id,brand_id,name,created_at,updated_at FROM car_types WHERE deleted_at IS NULL AND id=$1`
	mock.ExpectQuery(statement).WithArgs(id).WillReturnError(sql.ErrNoRows)
	res, err := repository.ReadBy("id", "=", id)

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestCarTypeQuery_Count(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	repository := NewCarTypeQuery(db)

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_types WHERE deleted_at IS NULL AND LOWER(name) LIKE $1`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("")

	assert.NoError(t, err)
	assert.NotZero(t, res)
}

func TestCarTypeQuery_CountZero(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 0
	repository := NewCarTypeQuery(db)

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_types WHERE deleted_at IS NULL AND LOWER(name) LIKE $1`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("")

	assert.NoError(t, err)
	assert.Zero(t, res)
}

func TestCarTypeQuery_CountBy(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	typeName := gofakeit.CarType()
	count := 1
	repository := NewCarTypeQuery(db)

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_types WHERE deleted_at IS NULL AND name=$1`
	mock.ExpectQuery(statement).WithArgs(typeName).WillReturnRows(rows)
	res, err := repository.CountBy("name", "=", "", typeName)

	assert.NoError(t, err)
	assert.NotZero(t, res)
}

func TestCarTypeQuery_CountByWithId(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	typeName := gofakeit.CarType()
	id := gofakeit.UUID()
	count := 1
	repository := NewCarTypeQuery(db)

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_types WHERE deleted_at IS NULL AND name=$1 AND id<>$2`
	mock.ExpectQuery(statement).WithArgs(typeName, id).WillReturnRows(rows)
	res, err := repository.CountBy("name", "=", id, typeName)

	assert.NoError(t, err)
	assert.NotZero(t, res)
}
