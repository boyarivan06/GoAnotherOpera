package main

import (
	"GoAnotherOpera/api"
	"GoAnotherOpera/models"
	"fmt"
)

func main() {
	//api.Get_all_artists()
	artist, ok := api.GetOne[models.Artist]("artist", map[string]string{"id": "352045"})
	if ok {
		fmt.Println("artist name is", artist.Name)
		fmt.Println(artist)
	}
}
