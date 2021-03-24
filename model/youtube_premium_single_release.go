package model

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
  PublishingYear time.Year
  CustomID    string
  RightsController string
}
