package queries

import (
	"booking-car/domain/models"
	"booking-car/domain/queries"
	"booking-car/pkg/postgresql"
	"strings"
)

type CartQuery struct {
	db postgresql.IConnection
}

func NewCartQuery(db postgresql.IConnection) queries.ICartQuery {
	return &CartQuery{db: db}
}

func (q CartQuery) Browse(search, orderBy, sort, userId string, limit, offset int) (interface{}, error) {
	var res []*models.Carts
	statement := models.CartSelectStatement + ` ` + models.CartJoinStatement + ` ` + models.CartDefaultWhereStatement + ` AND c.user_id=$1 AND ` +
		`(LOWER(c.car_brand) LIKE $2 OR LOWER(c.car_type) LIKE $2 OR LOWER(c.car_color) LIKE $2) ` +
		`ORDER BY ` + orderBy + ` ` + sort + ` LIMIT $3 OFFSET $4`

	rows, err := q.db.GetDbInstance().Query(statement, userId, "%"+strings.ToLower(search)+"%", limit, offset)
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

func (q CartQuery) BrowseAllByUser(userId string) (interface{}, error) {
	var res []*models.Carts
	statement := models.CartSelectStatement + ` ` + models.CartJoinStatement + ` ` + models.CartDefaultWhereStatement + ` AND c.user_id=$1`

	rows, err := q.db.GetDbInstance().Query(statement, userId)
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

func (q CartQuery) ReadBy(column, operator, userId string, value interface{}) (interface{}, error) {
	statement := models.CartSelectStatement + ` ` + models.CartJoinStatement + ` ` + models.CartDefaultWhereStatement + ` AND c.user_id=$1 AND ` + column + `` + operator + `$2`

	row := q.db.GetDbInstance().QueryRow(statement, userId, value)
	res, err := models.NewCartModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q CartQuery) Count(search, userId string) (res int, err error) {
	statement := models.CartCountSelectStatement + ` ` + models.CartJoinStatement + ` ` + models.CartDefaultWhereStatement + ` AND c.user_id=$1 AND ` +
		`(LOWER(c.car_brand) LIKE $2 OR LOWER(c.car_type) LIKE $2 OR LOWER(c.car_color) LIKE $2)`

	err = q.db.GetDbInstance().QueryRow(statement, "%"+strings.ToLower(search)+"%", userId).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q CartQuery) CountBy(column, operator, userId string, value interface{}) (res int, err error) {
	statement := models.CartCountSelectStatement + ` ` + models.CartDefaultWhereStatement + ` AND c.user_id=$1 AND ` + column + `` + operator + `$2`

	err = q.db.GetDbInstance().QueryRow(statement, userId, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
