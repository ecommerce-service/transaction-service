package usecases

import (
	"booking-car/domain/requests"
	"booking-car/domain/view_models"
)

type ICarTypeUseCase interface {
	GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CarTypeVm, pagination view_models.PaginationVm, err error)

	GetAll(search,brandID string) (res []view_models.CarTypeVm, err error)

	GetByID(id string) (res view_models.CarTypeVm, err error)

	Edit(req *requests.CarTypeRequest, id string) (res string, err error)

	Add(req *requests.CarTypeRequest) (res string, err error)

	Delete(id string) (err error)

	Count(search string) (res int, err error)

	CountBy(column, operator, id string, value interface{}) (res int, err error)
}
