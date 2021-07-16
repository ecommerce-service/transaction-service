package queries

import (
	"booking-car/domain/models"
	"booking-car/domain/queries"
	"database/sql"
	"strings"
)

type CarQuery struct {
	db *sql.DB
}

func NewCarQuery(db *sql.DB) queries.IBaseQuery {
	return &CarQuery{db: db}
}

func (q CarQuery) Browse(search, orderBy, sort string, limit, offset int) (interface{}, error) {
	var res []*models.Cars
	statement := models.CarSelectStatement + ` ` + models.CarJoinStatement + ` ` + models.CarDefaultWhereStatement + ` AND (production_year LIKE $1 OR LOWER(b.name) LIKE $1) ` +
		`ORDER BY ` + orderBy + ` ` + sort + ` LIMIT $2 OFFSET $3`

	rows, err := q.db.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		temp, err := models.NewCarModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.Cars))
	}

	return res, nil
}

func (q CarQuery) ReadBy(column, operator string, value interface{}) (interface{}, error) {
	statement := models.CarSelectStatement + ` ` + models.CarJoinStatement + ` ` + models.CarDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`

	row := q.db.QueryRow(statement, value)
	res, err := models.NewCarModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q CarQuery) Count(search string) (res int, err error) {
	statement := models.CarCountSelectStatement + ` ` + models.CarJoinStatement + ` ` + models.CarDefaultWhereStatement + ` AND (production_year LIKE $1 OR LOWER(b.name) LIKE $1)`

	err = q.db.QueryRow(statement,"%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res,err
	}

	return res,nil
}

func (q CarQuery) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	whereStatement := models.CarDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`
	whereParams := []interface{}{value}
	if id != "" {
		whereStatement += ` AND c.id<>$2`
		whereParams = append(whereParams, id)
	}
	statement := models.CarCountSelectStatement + ` ` + whereStatement

	err = q.db.QueryRow(statement, whereParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
