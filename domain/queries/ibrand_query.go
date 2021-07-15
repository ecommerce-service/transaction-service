package queries

type IBrandQuery interface {
	IBaseQuery
	BrowseAll(search string) (interface{}, error)
}
