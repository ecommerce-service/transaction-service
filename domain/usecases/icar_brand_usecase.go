package usecases

import (
	"booking-car/domain/requests"
	"booking-car/domain/view_models"
)

type ICarBrandUseCase interface {
	GetListWithPagination(search,orderBy,sort string,page,limit int) (res []view_models.CarBrandVm,pagination view_models.PaginationVm,err error)

	GetAll(search string) (res []view_models.CarBrandVm,err error)

	GetByID(id string) (res view_models.CarBrandVm,err error)

	Edit(req *requests.CarBrandRequest,id string) (res string,err error)

	Add(req *requests.CarBrandRequest) (res string,err error)

	Delete(id string) (err error)

	Count(search string) (res int,err error)

	CountBy(column,operator,id string,value interface{}) (res int,err error)
}
