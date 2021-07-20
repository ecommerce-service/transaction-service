package commands

import (
	"booking-car/domain/models"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddColor(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	model := models.NewCarColorModel().SetName(gofakeit.Name()).SetHexCode(gofakeit.HexColor()).SetCreatedAt(now).SetUpdatedAt(now)
	id := gofakeit.UUID()
	cmd := NewCarColorCommand(db, model)

	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `INSERT INTO car_colors (name,hex_code,created_at,updated_at) VALUES($1,$2,$3,$4) RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.Name(), model.HexCode(), model.CreatedAt(), model.UpdatedAt()).WillReturnRows(rows)
	res, err := cmd.Add()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestAddColorError(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	model := models.NewCarColorModel().SetCreatedAt(now).SetUpdatedAt(now)
	cmd := NewCarColorCommand(db, model)

	defer func() {
		db.Close()
	}()

	statement := `INSERT INTO car_colors (name,hex_code,created_at,updated_at) VALUES($1,$2,$3,$4) RETURNING id`
	mock.ExpectQuery(statement).WithArgs("", "", model.CreatedAt(), model.UpdatedAt()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Add()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestEditColor(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewCarColorModel().SetName(gofakeit.Name()).SetHexCode(gofakeit.HexColor()).SetUpdatedAt(now).SetId(id)
	cmd := NewCarColorCommand(db, model)

	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE car_colors SET name=$1,hex_code=$2,updated_at=$3 WHERE id=$4 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.Name(), model.HexCode(), model.UpdatedAt(), model.Id()).WillReturnRows(rows)
	res, err := cmd.Edit()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestEditColorError(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	model := models.NewCarColorModel().SetHexCode(gofakeit.HexColor()).SetCreatedAt(now).SetUpdatedAt(now)
	cmd := NewCarColorCommand(db, model)

	defer func() {
		db.Close()
	}()

	statement := `UPDATE car_colors SET name=$1,hex_code=$2,updated_at=$3 WHERE id=$4 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.Name(), model.HexCode(), model.UpdatedAt(), model.Id()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Edit()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestDeleteColor(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewCarColorModel().SetDeletedAt(now).SetUpdatedAt(now).SetId(id)
	cmd := NewCarColorCommand(db, model)

	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE car_colors SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(),model.DeletedAt().Time, model.Id()).WillReturnRows(rows)
	res, err := cmd.Delete()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteColorError(t *testing.T){
	db, mock := NewMock()

	now := time.Now().UTC()
	model := models.NewCarColorModel().SetDeletedAt(now).SetUpdatedAt(now)
	cmd := NewCarColorCommand(db, model)

	defer func() {
		db.Close()
	}()

	statement := `UPDATE car_colors SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(),model.DeletedAt().Time, model.Id()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Delete()

	assert.Error(t, err)
	assert.Empty(t, res)
}
