package commands

type ICartCommand interface {
	IBaseCommand

	EditQuantity() (res string,err error)

	DeleteAllByUserId() (err error)
}
