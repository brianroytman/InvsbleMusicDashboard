package model

import (
	"time"
)

type Single struct {
	ParterName  string
	ReleaseDate time.Time
	SingleTitle string
	Genre       string
  UPC         int
  Publisher   string
  PublishingYear time.Year
  CustomID    string
  //CreatedAt   time.Time
  //UpdatedAt   time.Time
  //TagList     []string
	//FavoritedBy []User
	//Author      User
	//Comments    []Comment
}

/*
type Comment struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Body      string
	Author    User
}
*/
