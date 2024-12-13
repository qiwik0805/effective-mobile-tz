package main

import (
	musicInfoDTO "effective-mobile-tz/internal/domain/dto/service/music_info"
	"effective-mobile-tz/internal/infra/service/music_info"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	r := gin.Default()

	fakeClient := music_info.NewFakeClient()

	r.GET("/info", func(c *gin.Context) {
		var request InfoRequest
		if err := c.ShouldBind(&request); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		infoResponse, err := fakeClient.Info(c, musicInfoDTO.InfoRequest{
			Group: request.Group,
			Song:  request.Song,
		})
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Error().Msgf("info: %s", err)
			return
		}

		response := &InfoResponse{
			ReleaseDate: infoResponse.ReleaseDate,
			Text:        infoResponse.Text,
			Link:        infoResponse.Link,
		}

		c.JSON(http.StatusNotFound, response)
	})

	if err := r.Run(":35090"); err != nil {
		log.Fatal().Msgf("run: %s", err)
	}
}

type InfoRequest struct {
	Group string `form:"group" binding:"required"`
	Song  string `form:"song" binding:"required"`
}

type InfoResponse struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
