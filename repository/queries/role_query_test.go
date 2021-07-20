package queries

import (
	"booking-car/domain/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoleQuery_BrowseAll(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	color := models.NewRoleModel().SetId(gofakeit.Int64()).SetName(gofakeit.Name())
	repository := NewRoleQuery(db)

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(color.Id(), color.Name())
	statement := `SELECT id,name FROM roles WHERE LOWER(name) LIKE $1 ORDER BY id ASC`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.BrowseAll("")

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
