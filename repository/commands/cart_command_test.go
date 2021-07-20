package commands

import (
	"booking-car/domain/models"
	"booking-car/pkg/postgresql"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestAddCart(t *testing.T) {
	db, mock := NewMock()
	con := postgresql.NewConnection(&postgresql.Config{})
	con.SetDb(db)
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewCartModel().SetUserId(gofakeit.UUID()).SetCarId(gofakeit.UUID()).SetCarBrand(gofakeit.Car().Brand).SetCarType(gofakeit.Car().Type).
		SetCarColor(gofakeit.HexColor()).SetProductionYear(strconv.Itoa(gofakeit.Car().Year)).SetPrice(gofakeit.Price(100000000, 150000000)).
		SetQuantity(int(gofakeit.Int8())).SetSubTotal(gofakeit.Price(100000000, 150000000)).SetCreatedAt(now).SetUpdatedAt(now)
	id := gofakeit.UUID()

	cmd := NewCartCommand(con, model)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `INSERT INTO carts(user_id,car_id,car_brand,car_type,car_color,production_year,price,quantity,sub_total,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UserId(), model.CarId(), model.CarBrand(), model.CarType(), model.CarColor(), model.ProductionYear(),
		model.Price(), model.Quantity(), model.SubTotal(), model.CreatedAt(), model.UpdatedAt()).WillReturnRows(rows)
	res, err := cmd.Add()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestAddCartError(t *testing.T) {
	db, mock := NewMock()
	con := postgresql.NewConnection(&postgresql.Config{})
	con.SetDb(db)
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewCartModel().SetCarBrand(gofakeit.Car().Brand).SetCarType(gofakeit.Car().Type).
		SetCarColor(gofakeit.HexColor()).SetProductionYear(strconv.Itoa(gofakeit.Car().Year)).SetPrice(gofakeit.Price(100000000, 150000000)).
		SetQuantity(int(gofakeit.Int8())).SetSubTotal(gofakeit.Price(100000000, 150000000)).SetCreatedAt(now).SetUpdatedAt(now)

	cmd := NewCartCommand(con, model)
	statement := `INSERT INTO carts(user_id,car_id,car_brand,car_type,car_color,production_year,price,quantity,sub_total,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UserId(), model.CarId(), model.CarBrand(), model.CarType(), model.CarColor(), model.ProductionYear(),
		model.Price(), model.Quantity(), model.SubTotal(), model.CreatedAt(), model.UpdatedAt()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Add()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestEditCart(t *testing.T) {
	db, mock := NewMock()
	con := postgresql.NewConnection(&postgresql.Config{})
	con.SetDb(db)
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewCartModel().SetCarId(gofakeit.UUID()).SetCarBrand(gofakeit.Car().Brand).SetCarType(gofakeit.Car().Type).
		SetCarColor(gofakeit.HexColor()).SetProductionYear(strconv.Itoa(gofakeit.Car().Year)).SetPrice(gofakeit.Price(100000000, 150000000)).
		SetQuantity(int(gofakeit.Int8())).SetSubTotal(gofakeit.Price(100000000, 150000000)).SetCreatedAt(now).SetUpdatedAt(now).SetId(id)

	cmd := NewCartCommand(con, model)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE carts SET car_id=$1,car_brand=$2,car_type=$3,car_color=$4,production_year=$5,price=$6,quantity=$7,sub_total=$8,updated_at=$9 WHERE id=$10 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.CarId(), model.CarBrand(), model.CarType(), model.CarColor(), model.ProductionYear(), model.Price(), model.Quantity(),
		model.SubTotal(), model.UpdatedAt(), model.Id()).WillReturnRows(rows)
	res, err := cmd.Edit()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestEditCartError(t *testing.T) {
	db, mock := NewMock()
	con := postgresql.NewConnection(&postgresql.Config{})
	con.SetDb(db)
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewCartModel().SetCarId(gofakeit.UUID()).SetCarBrand(gofakeit.Car().Brand).SetCarType(gofakeit.Car().Type).
		SetCarColor(gofakeit.HexColor()).SetProductionYear(strconv.Itoa(gofakeit.Car().Year)).SetPrice(gofakeit.Price(100000000, 150000000)).
		SetQuantity(int(gofakeit.Int8())).SetSubTotal(gofakeit.Price(100000000, 150000000)).SetCreatedAt(now).SetUpdatedAt(now)

	cmd := NewCartCommand(con, model)
	statement := `UPDATE carts SET car_id=$1,car_brand=$2,car_type=$3,car_color=$4,production_year=$5,price=$6,quantity=$7,sub_total=$8,updated_at=$9 WHERE id=$10 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.CarId(), model.CarBrand(), model.CarType(), model.CarColor(), model.ProductionYear(), model.Price(), model.Quantity(),
		model.SubTotal(), model.UpdatedAt(), model.Id()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Edit()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestDeleteCart(t *testing.T) {
	db, mock := NewMock()
	con := postgresql.NewConnection(&postgresql.Config{})
	con.SetDb(db)
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewCartModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)

	cmd := NewCartCommand(con, model)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE carts SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(), model.DeletedAt().Time, model.Id()).WillReturnRows(rows)
	res, err := cmd.Delete()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteCartError(t *testing.T) {
	db, mock := NewMock()
	con := postgresql.NewConnection(&postgresql.Config{})
	con.SetDb(db)
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewCartModel().SetUpdatedAt(now).SetDeletedAt(now)

	cmd := NewCartCommand(con, model)
	statement := `UPDATE carts SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(), model.DeletedAt().Time, model.Id()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Delete()

	assert.Error(t, err)
	assert.Empty(t, res)
}
