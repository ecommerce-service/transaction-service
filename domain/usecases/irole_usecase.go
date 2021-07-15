package usecases

import "booking-car/domain/view_models"

type IRoleUseCase interface {
	BrowseAll(search string) (res []view_models.RoleVm, err error)
}
