package server

import (
	"context"
	"database/sql"
	songRepository "effective-mobile-tz/internal/app/song/repository/song"
	songServer "effective-mobile-tz/internal/app/song/server/song"
	songService "effective-mobile-tz/internal/app/song/service/song"
	pgMigrator "effective-mobile-tz/internal/infra/db/migrations/postgres"
	ginMiddleware "effective-mobile-tz/internal/infra/http/middleware/gin"
	"effective-mobile-tz/internal/infra/service/music_info"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Migrator interface {
	Migrate(ctx context.Context) error
}

type Server struct {
	config Config
}

func NewServer(config Config) *Server {
	return &Server{config: config}
}

func (s *Server) Run(ctx context.Context) error {
	musicInfoClient := music_info.Factory(s.config.MusicInfoConfig)
	if s.config.MusicInfoConfig.BaseURL == "" && s.config.MusicInfoConfig.UseFake != "" {
		return fmt.Errorf("music info service base url is not set")
	}

	if s.config.MusicInfoConfig.UseFake == "" {
		log.Info().Msg(fmt.Sprintf("using music info client with base url=%s", s.config.MusicInfoConfig.BaseURL))
	} else {
		log.Info().Msg(fmt.Sprintf("using fake music info client"))
	}

	// Миграции
	{
		db, err := sql.Open("postgres", s.config.SongRepositoryConfig.DSN)
		if err != nil {
			return fmt.Errorf("sql open: %w", err)
		}

		if err = db.Ping(); err != nil {
			return fmt.Errorf("failed to ping db: %w", err)
		}

		pgZlgLogAdapter := pgMigrator.NewZlgLogAdapter(&log.Logger)
		defaultMigrator := pgMigrator.NewDefaultMigrator(db, pgZlgLogAdapter)
		if err = defaultMigrator.Migrate(ctx); err != nil {
			return fmt.Errorf("migrate: %w", err)
		}
	}

	pool, err := pgxpool.New(ctx, s.config.SongRepositoryConfig.DSN)
	if err != nil {
		return fmt.Errorf("pgxpool new: %w", err)
	}

	songRepository := songRepository.NewRepository(pool)
	songService := songService.NewService(songRepository, musicInfoClient)
	songServer := songServer.NewServer(songService)

	r := gin.New()
	r.Use(ginMiddleware.ErrorMiddleware)
	r.Use(ginMiddleware.DefaultStructuredLogger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	v1.POST("/songs", songServer.Add)
	v1.GET("/songs", songServer.GetAll)
	v1.GET("/songs/:songID/text", songServer.GetText)
	v1.PUT("/songs/:songID", songServer.Update)
	v1.DELETE("/songs/:songID", songServer.Remove)

	if err = r.Run(); err != nil {
		return fmt.Errorf("run: %w", err)
	}

	return nil
}
