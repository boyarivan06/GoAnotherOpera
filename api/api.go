package api

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/imroc/req/v3"
)

const ClientId = "e6ffd643"

func GetOne[T any](userParams map[string]string) (*T, bool) {
	client := req.C()
	queryParams := map[string]string{"client_id": ClientId, "format": "json"}
	for k, v := range userParams {
		queryParams[k] = v
	}
	var response Response[T]
	var t T
	_, err := client.R().SetPathParam("object", strings.ToLower(reflect.TypeOf(t).Name())).SetQueryParams(queryParams).SetSuccessResult(&response).Get("https://api.jamendo.com/v3.0/{object}s/")
	if err != nil {
		return nil, false
	}
	if len(response.Results) <= 0 {
		return nil, false
	}
	return &response.Results[0], true

}

func GetMany[T any](userParams map[string]string, limit int) (*[]T, bool) {
	client := req.C()
	queryParams := map[string]string{"client_id": ClientId, "format": "json", "limit": strconv.Itoa(limit)}
	for k, v := range userParams {
		queryParams[k] = v
	}
	var t T
	var response Response[T]
	_, err := client.R().SetPathParam("object", strings.ToLower(reflect.TypeOf(t).Name())).SetQueryParams(queryParams).SetSuccessResult(&response).Get("https://api.jamendo.com/v3.0/{object}s/")
	if err != nil {
		return nil, false
	}
	return &response.Results, true
}
