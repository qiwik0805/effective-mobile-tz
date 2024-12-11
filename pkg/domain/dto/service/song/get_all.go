package song

import (
	songServerDTO "effective-mobile-tz/pkg/domain/dto/server/song"
	"effective-mobile-tz/pkg/domain/model"
)

type GetAllRequest struct {
	Filter GetAllFilter
}

type GetAllResponse struct {
	Songs []model.Song `json:"songs"`
}

type GetAllFilter struct {
	Group       *string
	Song        *string
	ReleaseDate *string
	Text        *string
	Link        *string
	Page        int
	PageSize    int
}

func NewGetAllFilter(filter songServerDTO.GetAllFilter) GetAllFilter {
	return GetAllFilter{
		Group:       filter.Group,
		Song:        filter.Song,
		ReleaseDate: filter.ReleaseDate,
		Text:        filter.Text,
		Link:        filter.Link,
		Page:        filter.Page,
		PageSize:    filter.PageSize,
	}
}
