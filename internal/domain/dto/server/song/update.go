package song

import "effective-mobile-tz/internal/domain/model"

type UpdateRequest struct {
	ID model.SongID `json:"-" uri:"songID" binding:"required" validate:"required,gte=1"`

	Group       string `json:"group" validate:"required"`
	Song        string `json:"song" validate:"required"`
	ReleaseData string `json:"releaseData" validate:"required"`
	Text        string `json:"text" validate:"required"`
	Link        string `json:"link" validate:"required"`
}
