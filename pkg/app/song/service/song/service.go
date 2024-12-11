package song

import (
	"context"
	songRepoDTO "effective-mobile-tz/pkg/domain/dto/repository/song"
	musicInfoDTO "effective-mobile-tz/pkg/domain/dto/service/music_info"
	songServiceDTO "effective-mobile-tz/pkg/domain/dto/service/song"
	"effective-mobile-tz/pkg/domain/model"
	"fmt"
	"strings"
)

type Repository interface {
	GetAll(ctx context.Context, r songRepoDTO.GetAllRequest) (*songRepoDTO.GetAllResponse, error)
	GetText(ctx context.Context, r songRepoDTO.GetTextRequest) (*songRepoDTO.GetTextResponse, error)
	Remove(ctx context.Context, id model.SongID) error
	Update(ctx context.Context, r songRepoDTO.UpdateRequest) error
	Add(ctx context.Context, r songRepoDTO.AddRequest) (model.SongID, error)
}

type MusicInfoClient interface {
	Info(ctx context.Context, r musicInfoDTO.InfoRequest) (*musicInfoDTO.InfoResponse, error)
}

type Service struct {
	repository      Repository
	musicInfoClient MusicInfoClient
}

func NewService(repository Repository, musicInfoClient MusicInfoClient) *Service {
	return &Service{repository: repository, musicInfoClient: musicInfoClient}
}

func (s *Service) GetAll(ctx context.Context, r songServiceDTO.GetAllRequest) (*songServiceDTO.GetAllResponse, error) {
	request := songRepoDTO.GetAllRequest{Filter: songRepoDTO.NewGetAllFilter(r.Filter)}
	getAllResponse, err := s.repository.GetAll(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get all: %w", err)
	}

	response := &songServiceDTO.GetAllResponse{Songs: getAllResponse.Songs}
	return response, nil
}

func (s *Service) GetText(ctx context.Context, r songServiceDTO.GetTextRequest) (*songServiceDTO.GetTextResponse, error) {
	getTextResponse, err := s.repository.GetText(ctx, songRepoDTO.GetTextRequest{
		ID: r.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("get text: %w", err)
	}

	text := getTextResponse.Text
	split := strings.Split(text, "\n\n")

	// Убираем пустые куплеты
	verses := make([]string, 0, len(split))
	for _, verse := range split {
		if verse == "" {
			continue
		}
		verses = append(verses, verse)
	}

	startIndex := (r.Filter.Page - 1) * r.Filter.PageSize
	endIndex := startIndex + r.Filter.PageSize

	if startIndex >= len(verses) {
		return &songServiceDTO.GetTextResponse{Verses: []string{}}, nil // Возвращаем пустой слайс, если страница выходит за пределы
	}

	if endIndex > len(verses) {
		endIndex = len(verses)
	}

	paginatedVerses := verses[startIndex:endIndex]

	return &songServiceDTO.GetTextResponse{Verses: paginatedVerses}, nil
}

func (s *Service) Remove(ctx context.Context, id model.SongID) error {
	if err := s.repository.Remove(ctx, id); err != nil {
		return fmt.Errorf("remove: %w", err)
	}

	return nil
}

func (s *Service) Update(ctx context.Context, r songServiceDTO.UpdateRequest) error {
	if err := s.repository.Update(ctx, songRepoDTO.UpdateRequest{
		ID:          r.ID,
		Group:       r.Group,
		Song:        r.Song,
		ReleaseData: r.ReleaseData,
		Text:        r.Text,
		Link:        r.Link,
	}); err != nil {
		return fmt.Errorf("update: %w", err)
	}

	return nil
}

func (s *Service) Add(ctx context.Context, r songServiceDTO.AddRequest) (model.SongID, error) {
	infoResponse, err := s.musicInfoClient.Info(ctx, musicInfoDTO.InfoRequest{
		Group: r.Group,
		Song:  r.Song,
	})
	if err != nil {
		return 0, fmt.Errorf("info: %w", err)
	}

	songID, err := s.repository.Add(ctx, songRepoDTO.AddRequest{
		Group:       r.Group,
		Song:        r.Song,
		ReleaseData: infoResponse.ReleaseDate,
		Text:        infoResponse.Text,
		Link:        infoResponse.Link,
	})
	if err != nil {
		return 0, fmt.Errorf("add: %w", err)
	}

	return songID, nil
}
