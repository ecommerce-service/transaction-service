package queries

import (
	"booking-car/domain/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestCartQuery_BrowseByUser(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	userId := gofakeit.UUID()
	search := "%%"
	model := models.NewCartModel().SetId(gofakeit.UUID()).SetCarId(gofakeit.UUID()).SetCarBrand(gofakeit.Car().Brand).SetCarType(gofakeit.CarType()).SetCarColor(gofakeit.Color()).
		SetProductionYear(strconv.Itoa(gofakeit.Year())).SetPrice(gofakeit.Price(100000000, 130000000)).SetQuantity(int(gofakeit.Int8())).
		SetSubTotal(gofakeit.Price(100000000, 130000000)).SetCreatedAt(now).SetUpdatedAt(now)
	repository := CartQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"id", "c.car_id", "c.car_brand", "c.car_type", "c.car_color", "c.production_year", "c.price", "c.quantity", "c.sub_total", "c.created_at",
		"c.updated_at"}).AddRow(model.Id(), model.CarId(), model.CarBrand(), model.CarType(), model.CarColor(), model.ProductionYear(), model.Price(), model.Quantity(), model.SubTotal(),
		model.CreatedAt(), model.UpdatedAt())
	statement := `SELECT c.id,c.car_id,c.car_brand,c.car_type,c.car_color,c.production_year,c.price,c.quantity,c.sub_total,c.created_at,c.updated_at FROM carts c ` +
		`INNER JOIN users u ON u.id = c.user_id AND u.deleted_at IS NULL WHERE c.deleted_at IS NULL AND c.user_id=$1 AND ` +
		`(LOWER(c.car_brand) LIKE $2 OR LOWER(c.car_type) LIKE $2 OR LOWER(c.car_color) LIKE $2) ORDER BY created_at desc LIMIT $3 OFFSET $4`
	mock.ExpectQuery(statement).WithArgs(userId, search, 10, 0).WillReturnRows(rows)
	res, err := repository.BrowseByUser("", "created_at", "desc", userId, 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCartQuery_BrowseAllByUser(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	userId := gofakeit.UUID()
	model := models.NewCartModel().SetId(gofakeit.UUID()).SetCarId(gofakeit.UUID()).SetCarBrand(gofakeit.Car().Brand).SetCarType(gofakeit.CarType()).SetCarColor(gofakeit.Color()).
		SetProductionYear(strconv.Itoa(gofakeit.Year())).SetPrice(gofakeit.Price(100000000, 130000000)).SetQuantity(int(gofakeit.Int8())).
		SetSubTotal(gofakeit.Price(100000000, 130000000)).SetCreatedAt(now).SetUpdatedAt(now)
	repository := CartQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"id", "c.car_id", "c.car_brand", "c.car_type", "c.car_color", "c.production_year", "c.price", "c.quantity", "c.sub_total", "c.created_at",
		"c.updated_at"}).AddRow(model.Id(), model.CarId(), model.CarBrand(), model.CarType(), model.CarColor(), model.ProductionYear(), model.Price(), model.Quantity(), model.SubTotal(),
		model.CreatedAt(), model.UpdatedAt())
	statement := `SELECT c.id,c.car_id,c.car_brand,c.car_type,c.car_color,c.production_year,c.price,c.quantity,c.sub_total,c.created_at,c.updated_at FROM carts c ` +
		`INNER JOIN users u ON u.id = c.user_id AND u.deleted_at IS NULL WHERE c.deleted_at IS NULL AND c.user_id=$1`
	mock.ExpectQuery(statement).WithArgs(userId).WillReturnRows(rows)
	res, err := repository.BrowseAllByUser(userId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCartQuery_ReadBy(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	userId := gofakeit.UUID()
	id := gofakeit.UUID()
	model := models.NewCartModel().SetId(id).SetCarId(gofakeit.UUID()).SetCarBrand(gofakeit.Car().Brand).SetCarType(gofakeit.CarType()).SetCarColor(gofakeit.Color()).
		SetProductionYear(strconv.Itoa(gofakeit.Year())).SetPrice(gofakeit.Price(100000000, 130000000)).SetQuantity(int(gofakeit.Int8())).
		SetSubTotal(gofakeit.Price(100000000, 130000000)).SetCreatedAt(now).SetUpdatedAt(now)
	repository := CartQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"id", "c.car_id", "c.car_brand", "c.car_type", "c.car_color", "c.production_year", "c.price", "c.quantity", "c.sub_total", "c.created_at",
		"c.updated_at"}).AddRow(model.Id(), model.CarId(), model.CarBrand(), model.CarType(), model.CarColor(), model.ProductionYear(), model.Price(), model.Quantity(), model.SubTotal(),
		model.CreatedAt(), model.UpdatedAt())
	statement := `SELECT c.id,c.car_id,c.car_brand,c.car_type,c.car_color,c.production_year,c.price,c.quantity,c.sub_total,c.created_at,c.updated_at FROM carts c ` +
		`INNER JOIN users u ON u.id = c.user_id AND u.deleted_at IS NULL WHERE c.deleted_at IS NULL AND c.user_id=$1 AND c.id=$2`
	mock.ExpectQuery(statement).WithArgs(userId, id).WillReturnRows(rows)
	res, err := repository.ReadBy("c.id", "=", userId, id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCartQuery_Count(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	userId := gofakeit.UUID()
	repository := CartQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT count(c.id) FROM carts c INNER JOIN users u ON u.id = c.user_id AND u.deleted_at IS NULL WHERE c.deleted_at IS NULL AND c.user_id=$1 AND (LOWER(c.car_brand) LIKE $2 OR LOWER(c.car_type) LIKE $2 OR LOWER(c.car_color) LIKE $2)`
	mock.ExpectQuery(statement).WithArgs(userId, search).WillReturnRows(rows)
	res, err := repository.Count("", userId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCartQuery_CountBy(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	count := 1
	carId := gofakeit.UUID()
	userId := gofakeit.UUID()
	repository := CartQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT count(c.id) FROM carts c WHERE c.deleted_at IS NULL AND c.user_id=$1 AND c.car_id=$2`
	mock.ExpectQuery(statement).WithArgs(userId,carId).WillReturnRows(rows)
	res, err := repository.CountBy("c.car_id", "=", userId, carId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
