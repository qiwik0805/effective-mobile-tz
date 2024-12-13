package song

import (
	"context"
	"effective-mobile-tz/internal/domain/domainerr"
	songRepoDTO "effective-mobile-tz/internal/domain/dto/repository/song"
	"effective-mobile-tz/internal/domain/model"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) GetAll(ctx context.Context, r songRepoDTO.GetAllRequest) (*songRepoDTO.GetAllResponse, error) {
	query := squirrel.Select("id", "group_name", "song", "release_date", "text", "link").
		From("song")

	filter := r.Filter
	if filter.Group != nil {
		query = query.Where(squirrel.ILike{"group_name": "%" + *filter.Group + "%"})
	}
	if filter.Song != nil {
		query = query.Where(squirrel.ILike{"song": "%" + *filter.Song + "%"})
	}
	if filter.ReleaseDate != nil {
		query = query.Where(squirrel.ILike{"release_date": "%" + *filter.ReleaseDate + "%"})
	}
	if filter.Text != nil {
		query = query.Where(squirrel.ILike{"text": "%" + *filter.Text + "%"})
	}
	if filter.Link != nil {
		query = query.Where(squirrel.ILike{"link": "%" + *filter.Link + "%"})
	}

	// Добавление пагинации
	if filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		query = query.Limit(uint64(filter.PageSize)).Offset(uint64(offset))
	}

	// Генерация SQL-запроса
	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql: %w", err)
	}

	log.Debug().Str("sql", sql).Interface("args", args).Msg("Generated Query")

	rows, err := repo.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	var songs []model.Song
	for rows.Next() {
		var song model.Song
		if err = rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		songs = append(songs, song)
	}

	response := &songRepoDTO.GetAllResponse{Songs: songs}
	return response, nil
}

func (repo *Repository) GetText(ctx context.Context, r songRepoDTO.GetTextRequest) (*songRepoDTO.GetTextResponse, error) {
	internalDB := New(repo.db)
	intID := int32(r.ID)
	text, err := internalDB.SelectSongText(ctx, intID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domainerr.NewNotFound(fmt.Sprintf("select song text with id=%d: %s", intID, err))
		}
		return nil, fmt.Errorf("select song text with id=%d: %w", intID, err)
	}

	response := &songRepoDTO.GetTextResponse{Text: text}
	return response, nil
}

func (repo *Repository) Remove(ctx context.Context, id model.SongID) error {
	internalDB := New(repo.db)
	intID := int32(id)
	if err := internalDB.DeleteSong(ctx, intID); err != nil {
		return fmt.Errorf("delete song with id=%d: %w", intID, err)
	}
	return nil
}

func (repo *Repository) Update(ctx context.Context, r songRepoDTO.UpdateRequest) error {
	internalDB := New(repo.db)
	intID := int32(r.ID)
	if err := internalDB.UpdateSong(ctx, &UpdateSongParams{
		Song:        r.Song,
		GroupName:   r.Group,
		ReleaseDate: r.ReleaseData,
		Text:        r.Text,
		Link:        r.Link,
		ID:          intID,
	}); err != nil {
		return fmt.Errorf("update song with id=%d: %w", intID, err)
	}

	return nil
}

func (repo *Repository) Add(ctx context.Context, r songRepoDTO.AddRequest) (model.SongID, error) {
	internalRepo := New(repo.db)
	songID, err := internalRepo.InsertSong(ctx, &InsertSongParams{
		Song:        r.Song,
		GroupName:   r.Group,
		ReleaseDate: r.ReleaseData,
		Text:        r.Text,
		Link:        r.Link,
	})
	if err != nil {
		return 0, fmt.Errorf("insert song: %w", err)
	}

	return model.SongID(songID), nil
}
