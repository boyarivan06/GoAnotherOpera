package api

import (
	"fmt"

	"github.com/imroc/req/v3"
)

const CLIENT_ID = "e6ffd643"

func Get_all_artists() {
	client := req.C()
	resp, err := client.R().Get("https://api.jamendo.com/v3.0/artists/?client_id=e6ffd643&format=json&namesearch=''")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.String())
}

func GetOne[T any](object string, user_params map[string]string) (*T, bool) {
	client := req.C()
	queryParams := map[string]string{"client_id": CLIENT_ID, "format": "json"}
	for k, v := range user_params {
		queryParams[k] = v
	}
	var responce Response[T]
	_, err := client.R().SetPathParam("object", object).SetQueryParams(queryParams).SetSuccessResult(&responce).Get("https://api.jamendo.com/v3.0/{object}s/")
	if err != nil {
		return nil, false
	}
	return &responce.Results[0], true

}

func GetMany(object string, user_params map[string]string) {
	client := req.C()
	query_params := map[string]string{"client_id": CLIENT_ID, "format": "json"}
	for k, v := range user_params {
		query_params[k] = v
	}
	resp, err := client.R().SetPathParam("object", object).SetQueryParams(query_params).Get("https://api.jamendo.com/v3.0/{object}s/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.String())
}
