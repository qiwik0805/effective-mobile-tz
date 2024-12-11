package song

import "effective-mobile-tz/pkg/domain/model"

type GetTextRequest struct {
	ID model.SongID
}

type GetTextResponse struct {
	Text string
}
