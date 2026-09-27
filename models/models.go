package models

type Artist struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Album struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Artist_id   string `json:"artist_id"`
	Artist_name string `json:"artist_name"`
	Image       string `json:"image"`
}

type Track struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	AlbumId     string `json:"album_id"`
	AlbumName   string `json:"album_name"`
	AlbumImage  string `json:"album_image"`
	ArtistId    string `json:"artist_id"`
	ArtistImage string `json:"artist_image"`
	Audio       string `json:"audio"`
	Duration    int    `json:"duration"`
	ArtistName  string `json:"artist_name"`
}
