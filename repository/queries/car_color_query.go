package queries

import (
	"booking-car/domain/models"
	"booking-car/domain/queries"
	"database/sql"
	"fmt"
	"strings"
)

type CarColorQuery struct {
	db *sql.DB
}

func NewCarColorQuery(db *sql.DB) queries.ICarColorQuery {
	return &CarColorQuery{db: db}
}

func (q CarColorQuery) Browse(search, orderBy, sort string, limit, offset int) (interface{}, error) {
	var res []*models.CarColors

	statement := models.CarColorSelectStatement + ` ` + models.CarColorDefaultWhereStatement + ` AND LOWER(name) LIKE $1 ORDER BY ` + orderBy + ` ` + sort + ` LIMIT $2 OFFSET $3`
	rows, err := q.db.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewCarColorModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.CarColors))
	}
	fmt.Println(res)

	return res, nil
}

func (q CarColorQuery) BrowseAll(search string) (interface{}, error) {
	var res []*models.CarColors

	statement := models.CarColorSelectStatement + ` ` + models.CarColorDefaultWhereStatement + ` AND LOWER(name) LIKE $1`
	rows, err := q.db.Query(statement, "%"+strings.ToLower(search)+"%")
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewCarColorModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.CarColors))
	}

	return res, nil
}

func (q CarColorQuery) ReadBy(column, operator string, value interface{}) (interface{}, error) {
	statement := models.CarColorSelectStatement + ` ` + models.CarColorDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`

	row := q.db.QueryRow(statement, value)
	res, err := models.NewCarColorModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q CarColorQuery) Count(search string) (res int, err error) {
	statement := models.CarColorCountSelectStatement + ` ` + models.CarColorDefaultWhereStatement + ` AND LOWER(name) LIKE $1`

	err = q.db.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q CarColorQuery) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	whereStatement := models.CarColorDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`
	whereParams := []interface{}{value}
	if id != "" {
		whereStatement += ` AND id<>$2`
		whereParams = append(whereParams, id)
	}
	statement := models.CarColorCountSelectStatement + ` ` + whereStatement

	err = q.db.QueryRow(statement, whereParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
