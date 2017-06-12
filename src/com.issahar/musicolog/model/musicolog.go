package model

import (
	"strings"
	"io/ioutil"
	"fmt"
	"os"
)



const musicDir = "/Users/issaharw/Music/iTunes/iTunes Media/Music"

type Song struct {
	Filename    string   `json:"filename"`
	Size        int64    `json:"size,omitempty"`
}



func init() {

}

func FindSongByName(name string) ([]Song, error) {
	name = strings.ToLower(name)

	ret := make([]Song, 0)

	ch := make(chan []Song)

	files, _ := ioutil.ReadDir(musicDir)

	for _, dir := range files {
		go FindSongForArtist(dir, name, ch)
	}


	num := 0
	for songs := range ch {
		if (songs != nil && len(songs) > 0) {
			for _, song := range songs {
				ret = append(ret, song)
			}
		}
		num ++
		if (num == len(files)) {
			close(ch)
		}

	}

	return ret, nil
}


func FindSongForArtist(artist os.FileInfo, name string, ch chan []Song) error {
	found := make([]Song, 0)
	fmt.Println("Called with: ", artist.Name())
	artistFullPath := musicDir + "/" + artist.Name()
	albums, _ := ioutil.ReadDir(artistFullPath)
	for _, album := range albums {
		albumFullPath := artistFullPath + "/" + album.Name()
		songs, _ := ioutil.ReadDir(albumFullPath)
		for _, song := range songs  {
			if strings.Contains(strings.ToLower(song.Name()), name) {
				found = append(found, Song{albumFullPath + "/" + song.Name(), song.Size()})
			}
		}
	}


	ch <- found

	return nil
}
