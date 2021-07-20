package queries

import (
	"booking-car/domain/models"
	"database/sql"
	"strings"
)

type CartQueryMock struct{
	db *sql.DB
}

func (q CartQueryMock) BrowseByUser(search, orderBy, sort, userId string, limit, offset int) (interface{}, error) {
	var res []*models.Carts
	statement := models.CartSelectStatement + ` ` + models.CartJoinStatement + ` ` + models.CartDefaultWhereStatement + ` AND c.user_id=$1 AND ` +
		`(LOWER(c.car_brand) LIKE $2 OR LOWER(c.car_type) LIKE $2 OR LOWER(c.car_color) LIKE $2) ` +
		`ORDER BY ` + orderBy + ` ` + sort + ` LIMIT $3 OFFSET $4`

	rows, err := q.db.Query(statement, userId, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewCartModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.Carts))
	}

	return res, nil
}

func (q CartQueryMock) BrowseAllByUser(userId string) (interface{}, error) {
	var res []*models.Carts
	statement := models.CartSelectStatement + ` ` + models.CartJoinStatement + ` ` + models.CartDefaultWhereStatement + ` AND c.user_id=$1`

	rows, err := q.db.Query(statement, userId)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewCartModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.Carts))
	}

	return res, nil
}

func (q CartQueryMock) ReadBy(column, operator, userId string, value interface{}) (interface{}, error) {
	statement := models.CartSelectStatement + ` ` + models.CartJoinStatement + ` ` + models.CartDefaultWhereStatement + ` AND c.user_id=$1 AND ` + column + `` + operator + `$2`

	row := q.db.QueryRow(statement, userId, value)
	res, err := models.NewCartModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q CartQueryMock) Count(search, userId string) (res int, err error) {
	statement := models.CartCountSelectStatement + ` ` + models.CartJoinStatement + ` ` + models.CartDefaultWhereStatement + ` AND c.user_id=$1 AND ` +
		`(LOWER(c.car_brand) LIKE $2 OR LOWER(c.car_type) LIKE $2 OR LOWER(c.car_color) LIKE $2)`

	err = q.db.QueryRow(statement, userId,"%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q CartQueryMock) CountBy(column, operator, userId string, value interface{}) (res int, err error) {
	statement := models.CartCountSelectStatement + ` ` + models.CartDefaultWhereStatement + ` AND c.user_id=$1 AND ` + column + `` + operator + `$2`

	err = q.db.QueryRow(statement, userId, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

