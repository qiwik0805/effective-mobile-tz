package song

import "effective-mobile-tz/pkg/domain/model"

type RemoveRequest struct {
	ID model.SongID `uri:"song_id" binding:"required"`
}
