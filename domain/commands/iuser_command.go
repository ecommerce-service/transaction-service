package commands

type IUserCommand interface {
	IBaseCommand

	EditDeposit() (res string,err error)
}
