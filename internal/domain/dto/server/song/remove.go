package song

import "effective-mobile-tz/internal/domain/model"

type RemoveRequest struct {
	ID model.SongID `uri:"songID" binding:"required" validate:"required,gte=1"`
}
