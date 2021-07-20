package queries

import (
	"booking-car/domain/models"
	"booking-car/domain/queries"
	"database/sql"
	"strings"
)

type CarBrandQuery struct {
	db *sql.DB
}

func NewCarBrandQuery(db *sql.DB) queries.ICarBrandQuery {
	return &CarBrandQuery{db: db}
}

func (q CarBrandQuery) Browse(search, orderBy, sort string, limit, offset int) (interface{}, error) {
	var res []*models.CarBrands

	statement := models.BrandSelectStatement + ` ` + models.BrandDefaultWhereStatement + ` AND LOWER(name) LIKE $1 ORDER BY ` + orderBy + ` ` + sort + ` LIMIT $2 OFFSET $3`
	rows, err := q.db.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewCarBrandModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.CarBrands))
	}

	return res, nil
}

func (q CarBrandQuery) BrowseAll(search string) (interface{}, error) {
	var res []*models.CarBrands

	statement := models.BrandSelectStatement + ` ` + models.BrandDefaultWhereStatement + ` AND LOWER(name) LIKE $1`
	rows, err := q.db.Query(statement, "%"+strings.ToLower(search)+"%")
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewCarBrandModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.CarBrands))
	}

	return res, nil
}

func (q CarBrandQuery) ReadBy(column, operator string, value interface{}) (interface{}, error) {
	statement := models.BrandSelectStatement + ` ` + models.BrandDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`

	row := q.db.QueryRow(statement, value)
	res, err := models.NewCarBrandModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q CarBrandQuery) Count(search string) (res int, err error) {
	statement := models.BrandSelectCountStatement + ` ` + models.BrandDefaultWhereStatement + ` AND LOWER(name) LIKE $1`

	err = q.db.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q CarBrandQuery) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	whereStatement := models.BrandDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`
	whereParams := []interface{}{value}
	if id != "" {
		whereStatement += ` AND id<>$2`
		whereParams = append(whereParams, id)
	}
	statement := models.BrandSelectCountStatement + ` ` + whereStatement

	err = q.db.QueryRow(statement, whereParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
