package main

import (
	"GoAnotherOpera/api"
	"GoAnotherOpera/models"
	"os"

	qt "github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/multimedia"
)

var currentArtist *models.Artist
var currentAlbum *models.Album
var ui *MainWindowUi
var player *multimedia.QMediaPlayer
var playing = false
var started = false
var stopPosition int64 = 0

func main() {
	qt.NewQApplication(os.Args)
	ui = NewMainWindowUi()
	ui.MainWindow.Show()
	player = multimedia.NewQMediaPlayer()
	audioOutput := multimedia.NewQAudioOutput()
	player.SetAudioOutput(audioOutput)

	ui.songs_list.OnItemClicked(showSong)
	ui.artists_list.OnItemClicked(getAlbums)
	ui.albums_list.OnItemClicked(getSongs)
	ui.artist_search.OnTextEdited(searchArtist)
	ui.album_search.OnTextEdited(searchAlbum)
	ui.song_search.OnTextEdited(searchSong)
	player.OnPositionChanged(positionChanged)
	ui.song_slider.OnValueChanged(sliderMoved)
	ui.play_stop_button.OnClicked(playStop)
	qt.QApplication_Exec()
}

func playStop() {
	if !playing {
		if !started {
			started = true
		} else {
			player.SetPosition(stopPosition)
		}
		player.Play() // TODO: goroutine?
		playing = true
		ui.play_stop_button.SetText("STOP")
	} else {
		stopPosition = player.Position()
		playing = false
		player.Stop()
		ui.play_stop_button.SetText("PLAY")
	}
}

func sliderMoved(value int) {
	player.SetPosition(int64(value))
}

func positionChanged(position int64) {
	if playing {
		ui.song_slider.SetSliderPosition(int(position))
	}
}
func changeSong(song *models.Track) {
	url := qt.NewQUrl3(song.Audio)
	player.Stop()
	player.SetSource(url)
	playing = false
	ui.play_stop_button.SetText("PLAY")
	player.SetPosition(1)
	started = false
}

func searchSong(param1 string) {
	dewarn()
	songs, ok := api.GetMany[models.Track](map[string]string{"album_name": currentAlbum.Name, "namesearch": param1}, 15)
	if !ok {
		warn("Песни не найдены")
		return
	}
	ui.songs_list.Clear()
	var items []string
	for _, song := range *songs {
		items = append(items, song.Name)
	}
	ui.songs_list.AddItems(items)
}

func searchAlbum(param1 string) {
	dewarn()
	albums, ok := api.GetMany[models.Album](map[string]string{"artist_name": currentArtist.Name, "namesearch": param1}, 10)
	if !ok {
		warn("Альбомов не найдено")
		return
	}
	ui.albums_list.Clear()
	var items []string
	for _, album := range *albums {
		items = append(items, album.Name)
	}
	ui.albums_list.AddItems(items)
}

func getAlbums(item *qt.QListWidgetItem) {
	dewarn()
	ok := false
	currentArtist, ok = api.GetOne[models.Artist](map[string]string{"name": item.Text()})
	if !ok {
		warn("Артист не найден")
		return
	}
	albums, ok := api.GetMany[models.Album](map[string]string{"artist_name": currentArtist.Name}, 10)
	if !ok {
		warn("Альбомов не найдено")
		return
	}
	ui.albums_list.Clear()
	var items []string
	for _, album := range *albums {
		items = append(items, album.Name)
	}
	ui.albums_list.AddItems(items)
}

func getSongs(item *qt.QListWidgetItem) {
	dewarn()
	var ok = false
	currentAlbum, ok = api.GetOne[models.Album](map[string]string{"name": item.Text()})
	if !ok {
		warn("Альбом не найден")
		return
	}
	songs, ok := api.GetMany[models.Track](map[string]string{"album_name": currentAlbum.Name}, 15)
	if !ok {
		warn("Песни не найдены")
		return
	}
	ui.songs_list.Clear()
	var items []string
	for _, song := range *songs {
		items = append(items, song.Name)
	}
	ui.songs_list.AddItems(items)
}
func searchArtist(_search string) {
	dewarn()
	artists, ok := api.GetMany[models.Artist](map[string]string{"namesearch": _search}, 15)
	if !ok {
		warn("Артист не найден")
		return
	}
	ui.artists_list.Clear()
	var items []string
	for _, artist := range *artists {
		items = append(items, artist.Name)
	}
	ui.artists_list.AddItems(items)
}

func showSong(item *qt.QListWidgetItem) {
	song, ok := api.GetOne[models.Track](map[string]string{"name": item.Text(), "artist_name": currentArtist.Name})
	if !ok {
		warn("Песня не найдена")
		return
	}
	changeSong(song)
	ui.song_slider.SetMaximum(song.Duration)
	ui.song_slider.SetSliderPosition(1)
	ui.name_label.SetText(song.Name)
	ui.artist_label.SetText(currentArtist.Name)
	ui.record_label.SetText(currentAlbum.Name)
}

func warn(text string) {
	ui.warning_label.SetText(text)
}

func dewarn() {
	ui.warning_label.Clear()
}
