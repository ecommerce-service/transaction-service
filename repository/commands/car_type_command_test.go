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

func TestAddCarType(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	carBrandId := gofakeit.UUID()
	model := models.NewCarTypeModel().SetName(gofakeit.Name()).SetBrandId(carBrandId).SetCreatedAt(now).SetUpdatedAt(now)
	id := gofakeit.UUID()
	cmd := NewCarTypeCommand(db, model)

	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `INSERT INTO car_types(name,brand_id,created_at,updated_at) VALUES($1,$2,$3,$4) RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.Name(), model.BrandId(), model.CreatedAt(), model.UpdatedAt()).WillReturnRows(rows)
	res, err := cmd.Add()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestAddCarTypeError(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	model := models.NewCarTypeModel().SetCreatedAt(now).SetUpdatedAt(now)
	cmd := NewCarTypeCommand(db, model)

	defer func() {
		db.Close()
	}()

	statement := `INSERT INTO car_types(name,brand_id,created_at,updated_at) VALUES($1,$2,$3,$4) RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.Name(), model.BrandId(), model.CreatedAt(), model.UpdatedAt()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Add()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestEditCarType(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	carBrandId := gofakeit.UUID()
	id := gofakeit.UUID()
	model := models.NewCarTypeModel().SetName(gofakeit.Name()).SetBrandId(carBrandId).SetUpdatedAt(now).SetId(id)
	cmd := NewCarTypeCommand(db, model)

	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE car_types SET name=$1,brand_id=$2,updated_at=$3 WHERE id=$4 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.Name(), model.BrandId(), model.UpdatedAt(), model.Id()).WillReturnRows(rows)
	res, err := cmd.Edit()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestEditCarTypeError(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	model := models.NewCarTypeModel().SetName(gofakeit.Name()).SetUpdatedAt(now)
	cmd := NewCarTypeCommand(db, model)

	defer func() {
		db.Close()
	}()

	statement := `UPDATE car_types SET name=$1,brand_id=$2,updated_at=$3 WHERE id=$4 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.Name(), model.BrandId(), model.UpdatedAt(), model.Id()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Edit()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestDeleteCarType(t *testing.T) {
	db, mock := NewMock()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewCarTypeModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)
	cmd := NewCarTypeCommand(db, model)

	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE car_types SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(), model.DeletedAt().Time, model.Id()).WillReturnRows(rows)
	res, err := cmd.Delete()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteCarTypeError(t *testing.T){
	db, mock := NewMock()

	now := time.Now().UTC()
	model := models.NewCarTypeModel().SetUpdatedAt(now).SetDeletedAt(now)
	cmd := NewCarTypeCommand(db, model)

	defer func() {
		db.Close()
	}()

	statement := `UPDATE car_types SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(), model.DeletedAt().Time, model.Id()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Delete()

	assert.Error(t, err)
	assert.Empty(t, res)
}
