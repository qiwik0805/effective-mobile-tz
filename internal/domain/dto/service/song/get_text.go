package song

import "effective-mobile-tz/internal/domain/model"

type GetTextRequest struct {
	ID     model.SongID
	Filter GetTextFilter
}

type GetTextFilter struct {
	Page     int
	PageSize int
}

type GetTextResponse struct {
	Verses []string
}
