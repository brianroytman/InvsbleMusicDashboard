package models

import (
	"time"
)

type ContentIdSingle struct {
	ParterName  string
	ReleaseDate time.Time
	SingleTitle string
	Genre       string
  UPC         int
  Publisher   string
  PublishingYear time.Year
  CustomID    string
}
