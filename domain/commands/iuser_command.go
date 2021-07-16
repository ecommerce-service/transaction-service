package commands

type IUserCommand interface {
	IBaseCommand

	EditDeposit() (err error)
}
