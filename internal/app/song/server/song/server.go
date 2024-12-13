package song

import (
	"context"
	"effective-mobile-tz/internal/domain/domainerr"
	songServerDTO "effective-mobile-tz/internal/domain/dto/server/song"
	songServiceDTO "effective-mobile-tz/internal/domain/dto/service/song"
	"effective-mobile-tz/internal/domain/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	vld     *validator.Validate
}

func NewServer(service Service) *Server {
	srv := &Server{service: service, vld: validator.New()}
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
// @Param page query int true "Page number for pagination" default(1)
// @Param pageSize query int true "Page size for pagination" default(10)
// @Success 200 {object} songServerDTO.GetAllResponse
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /api/v1/songs [get]
func (srv *Server) GetAll(c *gin.Context) {
	var filter songServerDTO.GetAllFilter
	if err := c.ShouldBind(&filter); err != nil {
		c.Error(domainerr.NewBadRequest(err.Error()))
		return
	}

	if err := srv.vld.Struct(filter); err != nil {
		errors := err.(validator.ValidationErrors)
		c.Error(domainerr.NewBadRequest(errors.Error()))
		return
	}
	log.Debug().Msgf(fmt.Sprintf("filter: %#v", filter))

	getAllRequest := songServiceDTO.GetAllRequest{
		Filter: songServiceDTO.NewGetAllFilter(filter),
	}
	getAllResponse, err := srv.service.GetAll(c, getAllRequest)
	if err != nil {
		c.Error(fmt.Errorf("get all: %w", err))
		return
	}

	songs := make([]songServerDTO.GetAllHelper, 0, len(getAllResponse.Songs))
	for _, song := range getAllResponse.Songs {
		songs = append(songs, songServerDTO.NewGetAllHelper(song))
	}

	response := songServerDTO.GetAllResponse{Songs: songs}
	c.JSON(http.StatusOK, response)
}

// GetText godoc
// @Summary Get song text with pagination
// @Description Get song text with pagination
// @Tags songs
// @Accept  json
// @Produce  json
// @Param songID path int true "Song ID"
// @Param page query int false "Page number for pagination" default(1)
// @Param pageSize query int false "Page size for pagination" default(10)
// @Success 200 {object} songServerDTO.GetTextResponse
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /api/v1/songs/{songID}/text [get]
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
	request.Filter = filter

	if err := srv.vld.Struct(request); err != nil {
		errors := err.(validator.ValidationErrors)
		c.Error(domainerr.NewBadRequest(errors.Error()))
		return
	}

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

// Remove godoc
// @Summary Remove a song by ID
// @Description Remove a song by ID
// @Tags songs
// @Accept  json
// @Produce  json
// @Param songID path int true "Song ID to remove"
// @Success 204 "No Content"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /api/v1/songs/{songID} [delete]
func (srv *Server) Remove(c *gin.Context) {
	var removeRequest songServerDTO.RemoveRequest
	if err := c.ShouldBindUri(&removeRequest); err != nil {
		c.Error(domainerr.NewBadRequest(err.Error()))
		return
	}

	if err := srv.vld.Struct(removeRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		c.Error(domainerr.NewBadRequest(errors.Error()))
		return
	}

	if err := srv.service.Remove(c, removeRequest.ID); err != nil {
		c.Error(err)
	}

	c.Status(http.StatusNoContent)
}

// Update godoc
// @Summary Update a song by ID
// @Description Update a song by ID
// @Tags songs
// @Accept  json
// @Produce  json
// @Param songID path int true "Song ID to update"
// @Param request body songServerDTO.UpdateRequest true "Song update request"
// @Success 204 "No Content"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /api/v1/songs/{songID} [put]
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

	if err := srv.vld.Struct(updateRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		c.Error(domainerr.NewBadRequest(errors.Error()))
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

// Add godoc
// @Summary Add a new song
// @Description Adds a new song to the system.
// @Tags songs
// @Accept json
// @Produce json
// @Param request body songServerDTO.AddRequest true "Request body for adding a song"
// @Success 201 {string} string "Created"  header(X-Created-ID)
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /api/v1/songs [post]
func (srv *Server) Add(c *gin.Context) {
	var addRequest songServerDTO.AddRequest

	if err := c.ShouldBind(&addRequest); err != nil {
		c.Error(domainerr.NewBadRequest(err.Error()))
		return
	}

	if err := srv.vld.Struct(addRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		c.Error(domainerr.NewBadRequest(errors.Error()))
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
