package queries

import "booking-car/domain/models"

type IRoleQuery interface {
	BrowseAll(search string) (res []*models.Roles, err error)
}
