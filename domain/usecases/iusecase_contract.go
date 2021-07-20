package usecases

import "booking-car/domain/view_models"

type IUseCaseContract interface {
	SetPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string)

	SetPaginationResponse(page, limit, total int) (res view_models.PaginationVm)
}
