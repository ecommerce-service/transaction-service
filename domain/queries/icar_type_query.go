package queries

type ICarTypeQuery interface {
	IBaseQuery

	BrowseAll(search, brandId string) (interface{}, error)
}
