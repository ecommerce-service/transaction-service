package usecases

import (
	"booking-car/domain/requests"
	"booking-car/domain/view_models"
)

type IUserUseCase interface {
	GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.UserVm, pagination view_models.PaginationVm, err error)

	GetByID(ID string) (res view_models.UserVm, err error)

	Edit(req *requests.UserEditRequest, id string) (res string, err error)

	Add(req *requests.UserAddRequest) (res string, err error)

	Delete(ID string) (err error)

	Count(search string) (res int, err error)

	CountBy(column, operator, id string, value interface{}) (res int, err error)

	CheckDuplication(email, username, phoneNumber, id string) (bool, error)
}
