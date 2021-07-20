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

func TestCarColorQuery_Browse(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	color := models.NewCarColorModel().SetId(gofakeit.UUID()).SetName(gofakeit.Color()).SetHexCode(gofakeit.HexColor()).SetCreatedAt(now).SetUpdatedAt(now)
	repository := NewCarColorQuery(db)

	rows := sqlmock.NewRows([]string{"id", "name", "hex_code", "created_at", "updated_at"}).AddRow(color.Id(), color.Name(), color.HexCode(), color.CreatedAt(), color.UpdatedAt())
	statement := `SELECT id,name,hex_code,created_at,updated_at FROM car_colors WHERE deleted_at IS NULL AND LOWER(name) LIKE $1 ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(search, 10, 0).WillReturnRows(rows)
	res, err := repository.Browse("", "created_at", "desc", 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarColorQuery_BrowseNoResultNoError(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	repository := NewCarColorQuery(db)

	rows := sqlmock.NewRows([]string{"id", "name", "hex_code", "created_at", "updated_at"})
	statement := `SELECT id,name,hex_code,created_at,updated_at FROM car_colors WHERE deleted_at IS NULL AND LOWER(name) LIKE $1 ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(search, 10, 0).WillReturnRows(rows)
	res, err := repository.Browse("", "created_at", "desc", 10, 0)

	assert.NoError(t, err)
	assert.Empty(t, res)
}

func TestCarColorQuery_BrowseAll(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	color := models.NewCarColorModel().SetId(gofakeit.UUID()).SetName(gofakeit.Color()).SetHexCode(gofakeit.HexColor()).SetCreatedAt(now).SetUpdatedAt(now)
	repository := NewCarColorQuery(db)

	rows := sqlmock.NewRows([]string{"id", "name", "hex_code", "created_at", "updated_at"}).AddRow(color.Id(), color.Name(), color.HexCode(), color.CreatedAt(), color.UpdatedAt())
	statement := `SELECT id,name,hex_code,created_at,updated_at FROM car_colors WHERE deleted_at IS NULL AND LOWER(name) LIKE $1`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.BrowseAll("")

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarColorQuery_BrowseAllNoResultNoError(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	repository := NewCarColorQuery(db)

	rows := sqlmock.NewRows([]string{"id", "name", "hex_code", "created_at", "updated_at"})
	statement := `SELECT id,name,hex_code,created_at,updated_at FROM car_colors WHERE deleted_at IS NULL AND LOWER(name) LIKE $1`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.BrowseAll("")

	assert.NoError(t, err)
	assert.Empty(t, res)
}

func TestCarColorQuery_ReadBy(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	color := models.NewCarColorModel().SetId(gofakeit.UUID()).SetName(gofakeit.Color()).SetHexCode(gofakeit.HexColor()).SetCreatedAt(now).SetUpdatedAt(now)
	repository := NewCarColorQuery(db)


	rows := sqlmock.NewRows([]string{"id", "name", "hex_code", "created_at", "updated_at"}).AddRow(color.Id(), color.Name(), color.HexCode(), color.CreatedAt(), color.UpdatedAt())
	statement := `SELECT id,name,hex_code,created_at,updated_at FROM car_colors WHERE deleted_at IS NULL AND id=$1`
	mock.ExpectQuery(statement).WithArgs(id).WillReturnRows(rows)
	res, err := repository.ReadBy("id", "=", id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarColorQuery_ReadByNoResult(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	id := gofakeit.UUID()
	repository := NewCarColorQuery(db)

	statement := `SELECT id,name,hex_code,created_at,updated_at FROM car_colors WHERE deleted_at IS NULL AND id=$1`
	mock.ExpectQuery(statement).WithArgs(id).WillReturnError(sql.ErrNoRows)
	res, err := repository.ReadBy("id", "=", id)

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestCarColorQuery_Count(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	repository := NewCarColorQuery(db)


	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_colors WHERE deleted_at IS NULL AND LOWER(name) LIKE $1`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("")

	assert.NoError(t, err)
	assert.NotZero(t, res)
}

func TestCarColorQuery_CountZero(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	var count int
	repository := NewCarColorQuery(db)


	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_colors WHERE deleted_at IS NULL AND LOWER(name) LIKE $1`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("")

	assert.NoError(t, err)
	assert.Zero(t, res)
}

func TestCarColorQuery_CountByWithOutId(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	hexCode := gofakeit.HexColor()
	count := 1
	repository := NewCarColorQuery(db)


	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_colors WHERE deleted_at IS NULL AND hex_code=$1`
	mock.ExpectQuery(statement).WithArgs(hexCode).WillReturnRows(rows)
	res, err := repository.CountBy("hex_code", "=", "", hexCode)

	assert.NoError(t, err)
	assert.NotZero(t, res)
}

func TestCarColorQuery_CountByWithId(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	hexCode := gofakeit.HexColor()
	id := gofakeit.UUID()
	count := 1
	repository := NewCarColorQuery(db)


	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(id) FROM car_colors WHERE deleted_at IS NULL AND hex_code=$1 AND id<>$2`
	mock.ExpectQuery(statement).WithArgs(hexCode, id).WillReturnRows(rows)
	res, err := repository.CountBy("hex_code", "=", id, hexCode)

	assert.NoError(t, err)
	assert.NotZero(t, res)
}
