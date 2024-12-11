package song

import (
	"context"
	"effective-mobile-tz/pkg/domain/domainerr"
	songServerDTO "effective-mobile-tz/pkg/domain/dto/server/song"
	songServiceDTO "effective-mobile-tz/pkg/domain/dto/service/song"
	"effective-mobile-tz/pkg/domain/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type Service interface {
	GetAll(ctx context.Context, r songServiceDTO.GetAllRequest) (*songServiceDTO.GetAllResponse, error)
	GetText(ctx context.Context, r songServiceDTO.GetTextRequest) (*songServiceDTO.GetTextResponse, error)
	Remove(ctx context.Context, id model.SongID) error
	Update(ctx context.Context, r songServiceDTO.UpdateRequest) error
	Add(ctx context.Context, r songServiceDTO.AddRequest) (model.SongID, error)
}

type Server struct {
	service Service
}

func NewServer(service Service) *Server {
	srv := &Server{service: service}
	return srv
}

// GetAll godoc
// @Summary Get all songs with filter and pagination
// @Description Get all songs with filter and pagination
// @Tags songs
// @Accept  json
// @Produce  json
// @Param group query string false "Filter by group"
// @Param song query string false "Filter by song"
// @Param releaseDate query string false "Filter by releaseDate"
// @Param text query string false "Filter by text"
// @Param link query string false "Filter by link"
// @Param page query int false "Page number for pagination" default(1)
// @Param pageSize query int false "Page size for pagination" default(10)
// @Success 200 {object} songServerDTO.GetAllResponse
// @Failure 400 {object} domainerr.BadRequest
// @Router /api/v1/songs [get]
func (srv *Server) GetAll(c *gin.Context) {
	var filter songServerDTO.GetAllFilter
	if err := c.ShouldBind(&filter); err != nil {
		c.Error(domainerr.NewBadRequest(err.Error()))
		return
	}

	log.Debug().Msgf(fmt.Sprintf("filter: %#v", filter))

	getAllRequest := songServiceDTO.GetAllRequest{
		Filter: songServiceDTO.NewGetAllFilter(filter),
	}
	getAllResponse, err := srv.service.GetAll(c, getAllRequest)
	if err != nil {
		c.Error(fmt.Errorf("get all: %w", err))
		log.Error().Msgf(fmt.Sprintf("get all: %s", err))
		return
	}

	songs := make([]songServerDTO.GetAllHelper, 0, len(getAllResponse.Songs))
	for _, song := range getAllResponse.Songs {
		songs = append(songs, songServerDTO.NewGetAllHelper(song))
	}

	response := songServerDTO.GetAllResponse{Songs: songs}
	c.JSON(http.StatusOK, response)
}

func (srv *Server) GetText(c *gin.Context) {
	var request songServerDTO.GetTextRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.Error(domainerr.NewBadRequest(err.Error()))
		return
	}

	var filter songServerDTO.GetTextFilter
	if err := c.ShouldBind(&filter); err != nil {
		c.Error(domainerr.NewBadRequest(err.Error()))
		return
	}

	if filter.Page < 1 {
		c.Error(domainerr.NewBadRequest(fmt.Sprintf("query page param must be positive, got %d instead", filter.Page)))
		return
	}

	if filter.PageSize < 1 {
		c.Error(domainerr.NewBadRequest(fmt.Sprintf("query pageSize param must be positive, got %d instead", filter.PageSize)))
		return
	}

	request.Filter = filter

	log.Debug().Int("id", int(request.ID)).Interface("filter", filter).Msg("Get Text")

	getTextRequest := songServiceDTO.GetTextRequest{
		ID: request.ID,
		Filter: songServiceDTO.GetTextFilter{
			Page:     request.Filter.Page,
			PageSize: request.Filter.PageSize,
		},
	}
	getTextResponse, err := srv.service.GetText(c, getTextRequest)
	if err != nil {
		c.Error(fmt.Errorf("get text: %w", err))
		return
	}

	verses := make([]songServerDTO.GetTextHelper, 0, len(getTextResponse.Verses))
	for _, verse := range getTextResponse.Verses {
		verses = append(verses, songServerDTO.NewGetTextHelper(verse))
	}
	response := songServerDTO.GetTextResponse{Text: verses}

	c.JSON(http.StatusOK, response)
}

func (srv *Server) Remove(c *gin.Context) {
	var id model.SongID
	if err := c.ShouldBindUri(&id); err != nil {
		c.Error(domainerr.NewBadRequest(err.Error()))
		return
	}

	if err := srv.service.Remove(c, id); err != nil {
		c.Error(err)
	}

	c.Status(http.StatusNoContent)
}

func (srv *Server) Update(c *gin.Context) {
	var updateRequest songServerDTO.UpdateRequest
	if err := c.ShouldBindUri(&updateRequest); err != nil {
		c.Error(domainerr.NewBadRequest(err.Error()))
		return
	}

	if err := c.ShouldBind(&updateRequest); err != nil {
		c.Error(domainerr.NewBadRequest(err.Error()))
		return
	}

	if err := srv.service.Update(c, songServiceDTO.UpdateRequest{
		ID:          updateRequest.ID,
		Group:       updateRequest.Group,
		Song:        updateRequest.Song,
		ReleaseData: updateRequest.ReleaseData,
		Text:        updateRequest.Text,
		Link:        updateRequest.Link,
	}); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)

}

func (srv *Server) Add(c *gin.Context) {
	var addRequest songServerDTO.AddRequest

	if err := c.ShouldBind(&addRequest); err != nil {
		c.Error(domainerr.NewBadRequest(err.Error()))
		return
	}

	songID, err := srv.service.Add(c, songServiceDTO.AddRequest{
		Group: addRequest.Group,
		Song:  addRequest.Song,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.Header("X-Created-ID", strconv.Itoa(int(songID)))
	c.Status(http.StatusCreated)
}
