package models

type Artist struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Album struct {
	id          string
	name        string
	artist_id   string
	artist_name string
	image       string
}

type Track struct {
	id, name, album_id, album_name, album_image, artist_id, artist_image, audio string
	duration                                                                    float64
}
