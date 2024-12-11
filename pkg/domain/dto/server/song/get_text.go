package song

import "effective-mobile-tz/pkg/domain/model"

type GetTextRequest struct {
	ID     model.SongID `uri:"songID" binding:"required"`
	Filter GetTextFilter
}

type GetTextFilter struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
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
