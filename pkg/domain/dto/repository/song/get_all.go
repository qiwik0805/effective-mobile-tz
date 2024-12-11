package song

import (
	songServiceDTO "effective-mobile-tz/pkg/domain/dto/service/song"
	"effective-mobile-tz/pkg/domain/model"
)

type GetAllRequest struct {
	Filter GetAllFilter
}

type GetAllResponse struct {
	Songs []model.Song
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

func NewGetAllFilter(filter songServiceDTO.GetAllFilter) GetAllFilter {
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
