package song

import (
	"effective-mobile-tz/pkg/domain/model"
)

type UpdateRequest struct {
	ID          model.SongID
	Group       string
	Song        string
	ReleaseData string
	Text        string
	Link        string
}
