package main

import (
	"fmt"

	"github.com/imroc/req/v3"
)

func Get_all_artists() {
	client := req.C()
	resp, err := client.R().Get("https://api.jamendo.com/v3.0/artists/?client_id=e6ffd643&format=json&namesearch=''")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.String())
}
