package commands

import (
	"booking-car/domain/models"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestAdd(t *testing.T) {
	now := time.Now().UTC()
	db, mock := NewMock()
	model := models.NewCarBrandModel().SetName("avanza").SetCreatedAt(now).SetUpdatedAt(now)
	cmd := NewCarBrandCommand(db, model)

	defer func() {
		db.Close()
	}()

	var id string
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)

	statement := `INSERT INTO car_brands (name,created_at,updated_at) VALUES($1,$2,$3) RETURNING id`
	query := mock.ExpectQuery(statement)
	query.WithArgs(model.Name(), model.CreatedAt(), model.UpdatedAt()).WillReturnRows(rows)
	_, err := cmd.Add()
	assert.New(t).NoError(err)
}
