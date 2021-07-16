package commands

type ICarCommand interface {
	IBaseCommand

	EditStock() (err error)
}
