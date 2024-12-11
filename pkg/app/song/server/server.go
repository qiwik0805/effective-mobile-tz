package server

import (
	"context"
	songRepository "effective-mobile-tz/pkg/app/song/repository/song"
	songServer "effective-mobile-tz/pkg/app/song/server/song"
	songService "effective-mobile-tz/pkg/app/song/service/song"
	"effective-mobile-tz/pkg/domain/domainerr"
	"effective-mobile-tz/pkg/infra/service/music_info"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Server struct {
	config Config
}

func NewServer(config Config) *Server {
	return &Server{config: config}
}

func (s *Server) Run(ctx context.Context) error {
	var musicInfoClient songService.MusicInfoClient
	if s.config.MusicInfoConfig.UseFake == "" {
		httpTransport := &http.Client{}
		baseURL := s.config.MusicInfoConfig.BaseURL
		musicInfoClient = music_info.NewClient(httpTransport, baseURL)
		log.Info().Msgf(fmt.Sprintf("using music info client with base url=%s", baseURL))
	} else {
		musicInfoClient = music_info.NewFakeClient()
		log.Info().Msgf(fmt.Sprintf("using fake music info client"))
	}

	pool, err := pgxpool.New(ctx, s.config.SongRepositoryConfig.DSN)
	if err != nil {
		return fmt.Errorf("pgxpool new: %w", err)
	}
	if err = pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping db: %w", err)
	}

	songRepository := songRepository.NewRepository(pool)

	songService := songService.NewService(songRepository, musicInfoClient)
	songServer := songServer.NewServer(songService)

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 { // Если нет ошибок в контексте
			return
		}

		lastError := c.Errors.Last() // Получаем последнюю ошибку
		err := lastError.Err         // Получаем саму ошибку из c.Error

		var (
			errBadRequest *domainerr.BadRequest
		)

		statusCode := http.StatusInternalServerError // по умолчанию 500
		if errors.As(err, &errBadRequest) {
			statusCode = http.StatusBadRequest
		}

		log.Info().Err(err).Msg("endpoint error")
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})

	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	v1.POST("/songs", songServer.Add)
	v1.GET("/songs", songServer.GetAll)
	v1.GET("/songs/:songID/text", songServer.GetText)
	v1.PUT("/songs/:songID", songServer.Update)
	v1.DELETE("/songs/:songID", songServer.Remove)

	if err := r.Run(); err != nil {
		return fmt.Errorf("run: %w", err)
	}

	return nil
}
