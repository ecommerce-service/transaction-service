package queries

import (
	"booking-car/domain/models"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestCarQuery_Browse(t *testing.T) {
	db, mock := NewSqlMock()
	now := time.Now().UTC()
	search := "%%"
	brand := models.NewCarBrandModel().SetId(gofakeit.UUID()).SetName(gofakeit.Car().Brand)
	carType := models.NewCarTypeModel().SetId(gofakeit.UUID()).SetName(gofakeit.CarType())
	carColor := models.NewCarColorModel().SetId(gofakeit.UUID()).SetName(gofakeit.Color()).SetHexCode(gofakeit.HexColor())
	car := models.NewCarModel().SetId(gofakeit.UUID()).SetProductionYear(strconv.Itoa(gofakeit.Year())).SetPrice(gofakeit.Price(100000000, 130000000)).
		SetStock(int(gofakeit.Int8())).SetCreatedAt(now).SetUpdatedAt(now)
	car.CarBrands = brand
	car.CarTypes = carType
	car.CarColors = carColor
	repository := NewCarQuery(db)
	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"c.id", "c.production_year", "c.price", "c.stock", "c.created_at", "c.updated_at", "b.id", "b.name", "ct.id", "ct.name", "cc.id",
		"cc.name", "cc.hex_code"}).AddRow(car.Id(), car.ProductionYear(), car.Price(), car.Stock(), car.CreatedAt(), car.UpdatedAt(), car.CarBrands.Id(), car.CarBrands.Name(),
		car.CarTypes.Id(), car.CarTypes.Name(), car.CarColors.Id(), car.CarColors.Name(), car.CarColors.HexCode())
	statement := `SELECT c.id,c.production_year,c.price,c.stock,c.created_at,c.updated_at,b.id,b.name,ct.id,ct.name,cc.id,cc.name,cc.hex_code FROM cars c ` +
		`INNER JOIN car_types ct ON ct.id = c.car_type_id AND ct.deleted_at IS NULL INNER JOIN car_brands b ON b.id = ct.brand_id AND b.deleted_at IS NULL INNER JOIN ` +
		`car_colors cc ON cc.id = c.car_color_id AND cc.deleted_at IS NULL WHERE c.deleted_at IS NULL AND (production_year LIKE $1 OR LOWER(b.name) LIKE $1) ` +
		`ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(search, 10, 0).WillReturnRows(rows)
	res, err := repository.Browse("", "created_at", "desc", 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarQuery_BrowseNoResultNoError(t *testing.T) {
	db, mock := NewSqlMock()
	now := time.Now().UTC()
	search := "%%"
	brand := models.NewCarBrandModel().SetId(gofakeit.UUID()).SetName(gofakeit.Car().Brand)
	carType := models.NewCarTypeModel().SetId(gofakeit.UUID()).SetName(gofakeit.CarType())
	carColor := models.NewCarColorModel().SetId(gofakeit.UUID()).SetName(gofakeit.Color()).SetHexCode(gofakeit.HexColor())
	car := models.NewCarModel().SetId(gofakeit.UUID()).SetProductionYear(strconv.Itoa(gofakeit.Year())).SetPrice(gofakeit.Price(100000000, 130000000)).
		SetStock(int(gofakeit.Int8())).SetCreatedAt(now).SetUpdatedAt(now)
	car.CarBrands = brand
	car.CarTypes = carType
	car.CarColors = carColor
	repository := NewCarQuery(db)
	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"c.id", "c.production_year", "c.price", "c.stock", "c.created_at", "c.updated_at", "b.id", "b.name", "ct.id", "ct.name", "cc.id",
		"cc.name", "cc.hex_code"})
	statement := `SELECT c.id,c.production_year,c.price,c.stock,c.created_at,c.updated_at,b.id,b.name,ct.id,ct.name,cc.id,cc.name,cc.hex_code FROM cars c ` +
		`INNER JOIN car_types ct ON ct.id = c.car_type_id AND ct.deleted_at IS NULL INNER JOIN car_brands b ON b.id = ct.brand_id AND b.deleted_at IS NULL INNER JOIN ` +
		`car_colors cc ON cc.id = c.car_color_id AND cc.deleted_at IS NULL WHERE c.deleted_at IS NULL AND (production_year LIKE $1 OR LOWER(b.name) LIKE $1) ` +
		`ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(search, 10, 0).WillReturnRows(rows)
	res, err := repository.Browse("", "created_at", "desc", 10, 0)

	assert.NoError(t, err)
	assert.Empty(t, res)
}

func TestCarQuery_ReadBy(t *testing.T) {
	db, mock := NewSqlMock()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	brand := models.NewCarBrandModel().SetId(gofakeit.UUID()).SetName(gofakeit.Car().Brand)
	carType := models.NewCarTypeModel().SetId(gofakeit.UUID()).SetName(gofakeit.CarType())
	carColor := models.NewCarColorModel().SetId(gofakeit.UUID()).SetName(gofakeit.Color()).SetHexCode(gofakeit.HexColor())
	car := models.NewCarModel().SetId(gofakeit.UUID()).SetProductionYear(strconv.Itoa(gofakeit.Year())).SetPrice(gofakeit.Price(100000000, 130000000)).
		SetStock(int(gofakeit.Int8())).SetCreatedAt(now).SetUpdatedAt(now).SetId(id)
	car.CarBrands = brand
	car.CarTypes = carType
	car.CarColors = carColor
	repository := NewCarQuery(db)
	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows([]string{"c.id", "c.production_year", "c.price", "c.stock", "c.created_at", "c.updated_at", "b.id", "b.name", "ct.id", "ct.name", "cc.id",
		"cc.name", "cc.hex_code"}).AddRow(car.Id(), car.ProductionYear(), car.Price(), car.Stock(), car.CreatedAt(), car.UpdatedAt(), car.CarBrands.Id(), car.CarBrands.Name(),
		car.CarTypes.Id(), car.CarTypes.Name(), car.CarColors.Id(), car.CarColors.Name(), car.CarColors.HexCode())
	statement := `SELECT c.id,c.production_year,c.price,c.stock,c.created_at,c.updated_at,b.id,b.name,ct.id,ct.name,cc.id,cc.name,cc.hex_code FROM cars c INNER JOIN car_types ct ON ct.id = c.car_type_id AND ct.deleted_at IS NULL INNER JOIN car_brands b ON b.id = ct.brand_id AND b.deleted_at IS NULL INNER JOIN car_colors cc ON cc.id = c.car_color_id AND cc.deleted_at IS NULL WHERE c.deleted_at IS NULL AND c.id=$1`
	mock.ExpectQuery(statement).WithArgs(id).WillReturnRows(rows)
	res, err := repository.ReadBy("c.id", "=", id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCarQuery_ReadByNoResult(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	id := gofakeit.UUID()
	repository := NewCarQuery(db)

	statement := `SELECT c.id,c.production_year,c.price,c.stock,c.created_at,c.updated_at,b.id,b.name,ct.id,ct.name,cc.id,cc.name,cc.hex_code FROM cars c ` +
		`INNER JOIN car_types ct ON ct.id = c.car_type_id AND ct.deleted_at IS NULL INNER JOIN car_brands b ON b.id = ct.brand_id AND b.deleted_at IS NULL ` +
		`INNER JOIN car_colors cc ON cc.id = c.car_color_id AND cc.deleted_at IS NULL WHERE c.deleted_at IS NULL AND c.id=$1`
	mock.ExpectQuery(statement).WithArgs(id).WillReturnError(sql.ErrNoRows)
	_, err := repository.ReadBy("id", "=", id)

	assert.Error(t, err)
}

func TestCarQuery_Count(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	repository := NewCarQuery(db)
	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(c.id) FROM cars c INNER JOIN car_types ct ON ct.id = c.car_type_id AND ct.deleted_at IS NULL INNER JOIN car_brands b ON b.id = ct.brand_id AND b.deleted_at IS NULL INNER JOIN car_colors cc ON cc.id = c.car_color_id AND cc.deleted_at IS NULL WHERE c.deleted_at IS NULL AND (production_year LIKE $1 OR LOWER(b.name) LIKE $1)`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("")

	assert.NoError(t, err)
	assert.NotZero(t, res)
}

func TestCarQuery_CountZero(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 0
	repository := NewCarQuery(db)
	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(c.id) FROM cars c INNER JOIN car_types ct ON ct.id = c.car_type_id AND ct.deleted_at IS NULL INNER JOIN car_brands b ON b.id = ct.brand_id AND b.deleted_at IS NULL INNER JOIN car_colors cc ON cc.id = c.car_color_id AND cc.deleted_at IS NULL WHERE c.deleted_at IS NULL AND (production_year LIKE $1 OR LOWER(b.name) LIKE $1)`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("")

	assert.NoError(t, err)
	assert.Zero(t, res)
}

func TestCarQuery_CountBy(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	productionYear := strconv.Itoa(gofakeit.Year())
	count := 1
	repository := NewCarQuery(db)
	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(c.id) FROM cars c WHERE c.deleted_at IS NULL AND c.production_year=$1`
	mock.ExpectQuery(statement).WithArgs(productionYear).WillReturnRows(rows)
	res, err := repository.CountBy("c.production_year","=","",productionYear)

	assert.NoError(t, err)
	assert.NotZero(t, res)
}

func TestCarQuery_CountByWithId(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	productionYear := strconv.Itoa(gofakeit.Year())
	id := gofakeit.UUID()
	count := 1
	repository := NewCarQuery(db)
	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(c.id) FROM cars c WHERE c.deleted_at IS NULL AND c.production_year=$1 AND c.id<>$2`
	mock.ExpectQuery(statement).WithArgs(productionYear,id).WillReturnRows(rows)
	res, err := repository.CountBy("c.production_year","=",id,productionYear)

	assert.NoError(t, err)
	assert.NotZero(t, res)
}

