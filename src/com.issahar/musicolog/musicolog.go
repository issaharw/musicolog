package musicolog

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"log"
	"time"
	"net/url"
)


type Link struct {
	Url		url.URL		`json:"url"`
	LogTime		time.Time	`json:"logTime"`
	Description 	string 		`json:"string"`
	UserID		string		`json:"userId"`
}

type Playlist struct {
	ID	string		`json:"id"`
	Name	string		`json:"name"`
	Links	[]Link		`json:"links"`
}

type SearchResults struct {
	Playlists 	[]Playlist	`json:"playlists"`
	Links		[]Link		`json:"links"`
}


type MusicService interface {
	Log(link *Link) error
	Playlist(id string) (*Playlist, error)
	CreatePlaylist(playlist *Playlist) (*Playlist, error)
	AddLinksToPlaylist(links *[]Link, playlist *Playlist) error
	Search(str string) (*SearchResults, error)
}



