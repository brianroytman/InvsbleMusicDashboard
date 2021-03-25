package models

import (
	"time"
)

type SingleJSON struct {
	PartnerName string `json:"PartnerName,omitempty"`
	ReleaseDate string `json:"ReleaseDate,omitempty"`
	SongTitle   string `json:"SongTitle,omitempty"`
	ArtistName  string `json:"ArtistName, omitempty"`
	Genre       string `json:"Genre, omitempty"`
}
