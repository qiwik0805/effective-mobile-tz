package song

import "effective-mobile-tz/pkg/domain/model"

type GetAllRequest struct {
	Filter GetAllFilter
}

type GetAllFilter struct {
	Group       *string `form:"group"`
	Song        *string `form:"song"`
	ReleaseDate *string `form:"releaseDate"`
	Text        *string `form:"text"`
	Link        *string `form:"link"`
	Page        int     `form:"page"`
	PageSize    int     `form:"pageSize"`
}

type GetAllResponse struct {
	Songs []GetAllHelper `json:"songs"`
}

type GetAllHelper struct {
	ID          model.SongID `json:"id"`
	Group       string       `json:"group"`
	Song        string       `json:"song"`
	ReleaseDate string       `json:"releaseDate"`
	Text        string       `json:"text"`
	Link        string       `json:"link"`
}

func NewGetAllHelper(song model.Song) GetAllHelper {
	return GetAllHelper{
		ID:          song.ID,
		Group:       song.Group,
		Song:        song.Song,
		ReleaseDate: song.ReleaseDate,
		Text:        song.Text,
		Link:        song.Link,
	}
}
