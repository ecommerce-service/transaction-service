package queries

import (
	"booking-car/domain/models"
	"booking-car/domain/queries"
	"database/sql"
	"strings"
)

type CarTypeQuery struct {
	db *sql.DB
}

func NewCarTypeQuery(db *sql.DB) queries.ICarTypeQuery {
	return &CarTypeQuery{db: db}
}

func (q CarTypeQuery) Browse(search, orderBy, sort string, limit, offset int) (interface{}, error) {
	var res []*models.CarTypes
	statement := models.CarTypeSelectStatement + ` ` + models.CarTypeDefaultWhereStatement + ` AND LOWER(name) LIKE $1 ORDER BY ` + orderBy + ` ` + sort + ` LIMIT $2 OFFSET $3`

	rows, err := q.db.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		temp, err := models.NewCarType().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.CarTypes))
	}

	return res, nil
}

func (q CarTypeQuery) BrowseAll(search, brandId string) (interface{}, error) {
	var res []*models.CarTypes
	statement := models.CarTypeSelectStatement + ` ` + models.CarTypeDefaultWhereStatement + ` AND LOWER(name) LIKE $1 AND brand_id=$2`

	rows, err := q.db.Query(statement, "%"+strings.ToLower(search)+"%",brandId)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		temp, err := models.NewCarType().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.CarTypes))
	}

	return res, nil
}

func (q CarTypeQuery) ReadBy(column, operator string, value interface{}) (interface{}, error) {
	statement := models.CarTypeSelectStatement + ` ` + models.CarTypeDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`

	row := q.db.QueryRow(statement, value)
	res, err := models.NewCarType().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q CarTypeQuery) Count(search string) (res int, err error) {
	statement := models.CarTypeSelectCountStatement + ` ` + models.CarTypeDefaultWhereStatement + ` AND LOWER(name) LIKE $1`

	err = q.db.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q CarTypeQuery) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	whereStatement := models.CarTypeDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`
	whereParams := []interface{}{value}
	if id != "" {
		whereStatement += ` AND id<>$2`
		whereParams = append(whereParams, id)
	}
	statement := models.CarTypeSelectCountStatement + ` ` + whereStatement

	err = q.db.QueryRow(statement, whereParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
