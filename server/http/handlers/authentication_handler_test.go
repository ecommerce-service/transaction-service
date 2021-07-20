package handlers

import (
	"booking-car/domain/requests"
	"booking-car/domain/view_models"
	validator2 "booking-car/pkg/validator"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const baseUrl = "http://localhost:3000/api"

func newResponder(s int, c string, ct string) httpmock.Responder {
	resp := httpmock.NewStringResponse(s, c)
	resp.Header.Set("Content-Type", ct)
	return httpmock.ResponderFromResponse(resp)
}

func TestLoginSuccess(t *testing.T) {
	req := requests.LoginRequest{
		Username: "avivi27",
		Password: "27mei1993",
	}
	payload, _ := json.Marshal(&req)

	dummyLoginVm := view_models.LoginVm{
		Token:                 `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjY1Nzk5NTQsImlzcyI6ImF2aXZpMjciLCJwYXlsb2FkIjoiZXlKaGJHY2lPaUpTVTBFeFh6VWlMQ0psYm1NaU9pSkJNVEk0UTBKRExVaFRNalUySW4wLmxrNnQ4OWhoanZBdHV4WU40WW1wVHRzbGdfcmxJWmFlSVlsQXZiYmIybERpRW5EcDI5ZFVIUC1xYkZFUzhoQXNsN2lyTnV5cE5sZEVLTlhMaEM1dGtFRmE0S2w2ekNqVlhsQnZ1aUlXdGNMaEZ0MlpudDVIeDlWWXFwanI5TmMtZGtXSUczaWxfMzlyenRIalQ2M1ZfRWFrUGpmSTVRdHAxcUVtOG1ZaTZkZmRESTdaazRlSVJrSk50R1VhQUxtSFROdzdwUEZxb0x2UUFiMmQ1SE1YWHc3T3lDa1hSbTltRnZxd1ZNOVpCLTIxV1MzaS1hYVNraEsxYm1mQl9ZVTBtbHREX0RveTQ4b0JPdmxzeTBaZE9jSkRSbXFfQWdad19VbkxlTUtyOVRXUDV4bVJHX3BhSFBnYUNHMmI1aU9vdnYzcGNSbmFHSFU1dXVNSEdFbmFGTVJMdWNkaDJ2elMxd1VDM0J3Qk91OXQyMnpob1duNXZaMXFhbXhlQWtKNGJOVmRDWlFObnBFaDE0TGozX19MN1htcXF3VlhIcmJNWFl0cS1EX25hN3AzbTE0dDJFUC1uanIwMkhDX0p1MjJhaTBlV21NNURmNFhWaVBXeUQ3WGZqN29NS21kdkhmTnlod05QM2txdjJfNXFyRllPTUMxak50S3QtNW55bWFsWF91VE5WcXBFM3pjcHlQdUNQSDJodXNrSGl6M3M0NjNvSHZ5MWZpV25xUzhDYV9tczlRV2ROV0ZmeGFoa0hpal9UU1FGaGJGVWs2YnNhX0MwdnQ3am0xd0dhWlMzNnlzSTQ0akp3MlY5aktha2o4LVFBN24zR01ObERRQUZmQ25VS0xoOUFDSDVTd1FPU0JZcWlVRk45cmtvUUxnUWZxZzFyYmV5Vmo5VnkwLmpIa2hmZl9zSjZPaEZTNnF3em5Wd0EuV2VQWlc2VVBvaDFsZnhFM0JQeG1iaUk1dVJwYVBrQXZHZXV2VUU5NTNEU3NhdlJfcE80UVBBQXgxTjdEcHRHdWI4VHFPTnZQZmlxNkZ6OUJoenVDMXg5MkRHTWdvXzJ3cDhGWXpJTWpMNlUua2FrQVRHNXg4RlNvdjJ2Uk5KT2lYdyJ9.Nd_0asHYZPS91tgdmnLEBkAcMoJ7lVXGCEcw_8BFjfM`,
		TokenExpiredAt:        1626579954,
		RefreshToken:          `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjY1Nzk5NTQsImlzcyI6ImF2aXZpMjciLCJwYXlsb2FkIjoiZXlKaGJHY2lPaUpTVTBFeFh6VWlMQ0psYm1NaU9pSkJNVEk0UTBKRExVaFRNalUySW4wLmxrNnQ4OWhoanZBdHV4WU40WW1wVHRzbGdfcmxJWmFlSVlsQXZiYmIybERpRW5EcDI5ZFVIUC1xYkZFUzhoQXNsN2lyTnV5cE5sZEVLTlhMaEM1dGtFRmE0S2w2ekNqVlhsQnZ1aUlXdGNMaEZ0MlpudDVIeDlWWXFwanI5TmMtZGtXSUczaWxfMzlyenRIalQ2M1ZfRWFrUGpmSTVRdHAxcUVtOG1ZaTZkZmRESTdaazRlSVJrSk50R1VhQUxtSFROdzdwUEZxb0x2UUFiMmQ1SE1YWHc3T3lDa1hSbTltRnZxd1ZNOVpCLTIxV1MzaS1hYVNraEsxYm1mQl9ZVTBtbHREX0RveTQ4b0JPdmxzeTBaZE9jSkRSbXFfQWdad19VbkxlTUtyOVRXUDV4bVJHX3BhSFBnYUNHMmI1aU9vdnYzcGNSbmFHSFU1dXVNSEdFbmFGTVJMdWNkaDJ2elMxd1VDM0J3Qk91OXQyMnpob1duNXZaMXFhbXhlQWtKNGJOVmRDWlFObnBFaDE0TGozX19MN1htcXF3VlhIcmJNWFl0cS1EX25hN3AzbTE0dDJFUC1uanIwMkhDX0p1MjJhaTBlV21NNURmNFhWaVBXeUQ3WGZqN29NS21kdkhmTnlod05QM2txdjJfNXFyRllPTUMxak50S3QtNW55bWFsWF91VE5WcXBFM3pjcHlQdUNQSDJodXNrSGl6M3M0NjNvSHZ5MWZpV25xUzhDYV9tczlRV2ROV0ZmeGFoa0hpal9UU1FGaGJGVWs2YnNhX0MwdnQ3am0xd0dhWlMzNnlzSTQ0akp3MlY5aktha2o4LVFBN24zR01ObERRQUZmQ25VS0xoOUFDSDVTd1FPU0JZcWlVRk45cmtvUUxnUWZxZzFyYmV5Vmo5VnkwLmpIa2hmZl9zSjZPaEZTNnF3em5Wd0EuV2VQWlc2VVBvaDFsZnhFM0JQeG1iaUk1dVJwYVBrQXZHZXV2VUU5NTNEU3NhdlJfcE80UVBBQXgxTjdEcHRHdWI4VHFPTnZQZmlxNkZ6OUJoenVDMXg5MkRHTWdvXzJ3cDhGWXpJTWpMNlUua2FrQVRHNXg4RlNvdjJ2Uk5KT2lYdyJ9.Nd_0asHYZPS91tgdmnLEBkAcMoJ7lVXGCEcw_8BFjfM`,
		RefreshTokenExpiredAt: 1626579954,
	}
	dummyResponse,_ :=json.Marshal(&dummyLoginVm)

	
	rst := resty.New()
	httpmock.ActivateNonDefault(rst.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"POST",
		baseUrl+"/auth/login",
		newResponder(http.StatusOK, string(dummyResponse), "application/json"),
	)

	resp, err := rst.R().
		SetResult(&view_models.LoginVm{}).
		SetHeader("Content-Type", "application/json").
		SetBody(string(payload)).
		Post(baseUrl + "/auth/login")
	assert.NoError(t, err)
	assert.Equal(t, &dummyLoginVm,resp.Result())
	assert.Equal(t, http.StatusOK, resp.StatusCode())
}

func TestLoginValidationFailed(t *testing.T) {
	req := requests.LoginRequest{
		Username: "avivi27",
	}
	payload, _ := json.Marshal(&req)

	validator := validator2.NewValidator("en").SetValidator().SetTranslator()
	err := validator.GetValidator().Struct(req)
	assert.Error(t, err)

	rst := resty.New()
	httpmock.ActivateNonDefault(rst.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"POST",
		baseUrl+"/auth/login",
		newResponder(http.StatusBadRequest, string(payload), "application/json"),
	)

	resp, err := rst.R().
		SetResult(&view_models.LoginVm{}).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(baseUrl + "/auth/login")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
}

func TestLoginFailed(t *testing.T){
	req := requests.LoginRequest{
		Username: "avivi27",
		Password: "27mei1994",
	}
	payload, _ := json.Marshal(&req)

	rst := resty.New()
	httpmock.ActivateNonDefault(rst.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"POST",
		baseUrl+"/auth/login",
		newResponder(http.StatusUnauthorized, "", "application/json"),
	)

	resp, err := rst.R().
		SetResult(&view_models.LoginVm{}).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(baseUrl + "/auth/login")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode())
}

func TestRegisterSuccess(t *testing.T){
	req := requests.RegisterRequest{
		FirstName:   "syaikhul",
		LastName:    "hadi",
		Email:       "syaikhulhadime@gmail.com",
		Username:    "hadime27",
		Password:    "hadmime27",
		Address:     "",
		PhoneNumber: "082170311372",
	}
	payload, _ := json.Marshal(&req)

	validator := validator2.NewValidator("en").SetValidator().SetTranslator()
	err := validator.GetValidator().Struct(req)
	assert.NoError(t, err)

	rst := resty.New()
	httpmock.ActivateNonDefault(rst.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"POST",
		baseUrl+"/auth/register",
		newResponder(http.StatusOK, "", "application/json"),
	)

	resp, err := rst.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(baseUrl + "/auth/register")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
}

func TestRegisterFailed(t *testing.T){
	req := requests.RegisterRequest{
		FirstName:   "syaikhul",
		LastName:    "hadi",
		Username:    "hadime27",
		Address:     "",
		PhoneNumber: "082170311372",
	}
	payload, _ := json.Marshal(&req)

	validator := validator2.NewValidator("en").SetValidator().SetTranslator()
	err := validator.GetValidator().Struct(req)
	assert.Error(t, err)

	rst := resty.New()
	httpmock.ActivateNonDefault(rst.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"POST",
		baseUrl+"/auth/register",
		newResponder(http.StatusBadRequest, "", "application/json"),
	)

	resp, err := rst.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(baseUrl + "/auth/register")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
}
