package gin

import (
	"effective-mobile-tz/internal/domain/domainerr"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func ErrorMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 { // Если нет ошибок в контексте
		return
	}

	lastError := c.Errors.Last()
	err := lastError.Err

	var (
		errBadRequest *domainerr.BadRequest
		errNotFound   *domainerr.NotFound
	)

	statusCode := http.StatusInternalServerError // по умолчанию 500
	if errors.As(err, &errBadRequest) {
		statusCode = http.StatusBadRequest
	} else if errors.As(err, &errNotFound) {
		statusCode = http.StatusNotFound
	} else {
		log.Error().Err(err).Msg("Internal Error")
	}

	log.Debug().Err(err).Msg("Endpoint Error")

	c.AbortWithStatus(statusCode)
}
