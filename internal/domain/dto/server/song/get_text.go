package song

import "effective-mobile-tz/internal/domain/model"

type GetTextRequest struct {
	ID     model.SongID `uri:"songID" binding:"required" validate:"required,gte=1"`
	Filter GetTextFilter
}

type GetTextFilter struct {
	Page     int `form:"page" validate:"required,gte=1"`
	PageSize int `form:"pageSize" validate:"required,gte=1"`
}

type GetTextResponse struct {
	Text []GetTextHelper `json:"text"`
}

type GetTextHelper struct {
	Verse string `json:"verse"`
}

func NewGetTextHelper(verse string) GetTextHelper {
	return GetTextHelper{Verse: verse}
}
