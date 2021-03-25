package models

import (
	"time"
)

type YoutubePremiumSingle struct {
	ParterName  string
	ReleaseDate time.Time
	SingleTitle string
	Genre       string
  UPC         int
  Publisher   string
  PublishingYear time.Time
  CustomID    string
  RightsController string
}
