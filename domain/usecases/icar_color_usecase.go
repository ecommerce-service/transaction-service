package usecases

import (
	"booking-car/domain/requests"
	"booking-car/domain/view_models"
)

type ICarColorUseCase interface {
	GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CarColorVm, pagination view_models.PaginationVm, err error)

	GetAll(search string) (res []view_models.CarColorVm, err error)

	GetByID(id string) (res view_models.CarColorVm, err error)

	Edit(req *requests.CarColorRequest, id string) (res string, err error)

	Add(req *requests.CarColorRequest) (res string, err error)

	Delete(id string) (err error)

	Count(search string) (res int, err error)

	CountBy(column, operator, id string, value interface{}) (res int, err error)
}
