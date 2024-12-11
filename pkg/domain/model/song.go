package model

type SongID int

type Song struct {
	ID          SongID
	Group       string
	Song        string
	ReleaseDate string
	Text        string
	Link        string
}
