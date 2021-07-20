package queries

type ICarBrandQuery interface {
	IBaseQuery
	BrowseAll(search string) (interface{}, error)
}
