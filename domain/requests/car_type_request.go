package requests

type CarTypeRequest struct {
	Name    string `json:"name"`
	BrandID string `json:"brand_id"`
}
