package queries

import (
	"booking-car/domain/models"
	"booking-car/domain/queries"
	"booking-car/pkg/postgresql"
	"fmt"
	"strings"
)

type TransactionQuery struct {
	db postgresql.IConnection
}

func NewTransactionQuery(db postgresql.IConnection) queries.ITransactionQuery {
	return &TransactionQuery{db: db}
}

func (q TransactionQuery) Browse(search, orderBy, sort, transactionType string, limit, offset int) (interface{}, error) {
	var res []*models.Transactions
	whereStatement := models.TransactionDefaultWhereStatement + ` AND t.transaction_number LIKE $1`
	params := []interface{}{"%" + strings.ToLower(search) + "%", limit, offset}
	if transactionType != "" {
		whereStatement += ` AND t.transaction_type=$4`
		params = append(params, transactionType)
	}
	statement := models.TransactionSelectListStatement + ` FROM transactions t ` + models.TransactionListJoinStatement + ` ` + whereStatement + ` ORDER BY ` + orderBy + ` ` + sort +
		` LIMIT $2 OFFSET $3`

	rows, err := q.db.GetDbInstance().Query(statement, params...)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewTransactionModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.Transactions))
	}

	return res, nil
}

func (q TransactionQuery) ReadBy(column, operator string, value interface{}) (interface{}, error) {
	statement := models.TransactionSelectListStatement + ` ` + models.TransactionSelectDetailStatement + ` FROM transactions t ` + models.TransactionListJoinStatement + ` ` +
		models.TransactionDetailJoinStatement + ` ` + models.TransactionDefaultWhereStatement + ` AND ` + column + `` + operator + `$1 ` + models.TransactionGroupByStatement
	fmt.Println(statement)
	row := q.db.GetDbInstance().QueryRow(statement, value)
	res, err := models.NewTransactionModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q TransactionQuery) Count(search, userId, transactionType string) (res int, err error) {
	whereStatement := models.TransactionDefaultWhereStatement + ` AND t.transaction_number LIKE $1`
	if transactionType != "" {
		whereStatement += ` AND t.transaction_type='`+transactionType+`'`
	}
	if userId != ""{
		whereStatement += ` AND t.user_id='`+userId+`'`
	}
	statement := models.TransactionSelectCountStatement + ` ` + models.TransactionListJoinStatement + ` ` + whereStatement

	err = q.db.GetDbInstance().QueryRow(statement, "%" + strings.ToLower(search) + "%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q TransactionQuery) CountBy(column, operator string, value interface{}) (res int, err error) {
	statement := models.TransactionSelectCountStatement + ` WHERE ` + column + `` + operator + `$1`

	err = q.db.GetDbInstance().QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q TransactionQuery) CountAll() (res int, err error) {
	statement := models.TransactionSelectCountStatement

	err = q.db.GetDbInstance().QueryRow(statement).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q TransactionQuery) BrowseByUserId(search, orderBy, sort, userId, transactionType string, limit, offset int) (interface{}, error) {
	var res []*models.Transactions
	whereStatement := models.TransactionDefaultWhereStatement + ` AND t.transaction_number LIKE $1 AND t.user_id=$2`
	params := []interface{}{"%" + strings.ToLower(search) + "%", userId, limit, offset}
	if transactionType != "" {
		whereStatement += ` AND t.transaction_type=$5`
		params = append(params, transactionType)
	}
	statement := models.TransactionSelectListStatement + ` FROM transactions t ` + models.TransactionListJoinStatement + ` ` + whereStatement + ` ORDER BY ` + orderBy + ` ` + sort +
		` LIMIT $3 OFFSET $4`

	rows, err := q.db.GetDbInstance().Query(statement, params...)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewTransactionModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.Transactions))
	}

	return res, nil
}
