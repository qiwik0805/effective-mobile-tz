package song

import "effective-mobile-tz/internal/domain/model"

type GetTextRequest struct {
	ID model.SongID
}

type GetTextResponse struct {
	Text string
}
