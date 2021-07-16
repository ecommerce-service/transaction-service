package queries

type ICarColorQuery interface {
	IBaseQuery

	BrowseAll(search string) (interface{}, error)
}
