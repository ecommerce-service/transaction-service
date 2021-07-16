package commands

import "database/sql"

type ICarCommand interface {
	IBaseCommand

	EditStock(tx *sql.Tx) (err error)
}
