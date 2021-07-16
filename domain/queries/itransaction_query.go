package queries

type ITransactionQuery interface {
	Browse(search, orderBy, sort, transactionType string, limit, offset int) (interface{}, error)

	ReadBy(column, operator string, value interface{}) (interface{}, error)

	Count(search, userID, transactionType string) (res int, err error)

	CountBy(column, operator string, value interface{}) (res int, err error)

	CountAll() (res int, err error)

	BrowseByUserId(search, orderBy, sort, userId, transactionType string, limit, offset int) (interface{}, error)
}
