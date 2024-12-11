package song

import "effective-mobile-tz/pkg/domain/model"

type UpdateRequest struct {
	ID model.SongID `uri:"song_id" binding:"required"`

	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseData string `json:"releaseData"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
