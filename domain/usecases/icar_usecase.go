package usecases

import (
	"booking-car/domain/requests"
	"booking-car/domain/view_models"
)

type ICarUseCase interface {
	GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CarVm, pagination view_models.PaginationVm, err error)

	GetByID(id string) (res view_models.CarVm, err error)

	Edit(req *requests.CarRequest, id string) (res string, err error)

	EditStock(id string, reduceStock int) (err error)

	Add(req *requests.CarRequest) (res string, err error)

	Delete(id string) (err error)

	Count(search string) (res int, err error)

	CountBy(column, operator, id string, value interface{}) (res int, err error)

	ValidateDuplication(productionYear, carTypeId, carColorId, id string) (bool, error)

	ReduceStock(ids []string, reducedBy int) (err error)
}
