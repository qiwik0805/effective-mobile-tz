package song

import (
	"context"
	"effective-mobile-tz/internal/app/song/service/song/mock"
	songRepoDTO "effective-mobile-tz/internal/domain/dto/repository/song"
	songServiceDTO "effective-mobile-tz/internal/domain/dto/service/song"
	"effective-mobile-tz/internal/infra/service/music_info"
	"fmt"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GetText(t *testing.T) {
	repoController := minimock.NewController(t)
	repoMock := mock.NewRepositoryMock(repoController)

	service := NewService(repoMock, music_info.NewFakeClient())
	ctx := context.Background()

	t.Run("Success with pagination", func(t *testing.T) {
		testText := "Verse 1\nLine 1\nLine 2\n\nVerse 2\nLine 1\n\nVerse 3\nLine 1\nLine 2\nLine 3\n\nVerse 4"
		repoMock.GetTextMock.Expect(ctx, songRepoDTO.GetTextRequest{ID: 1}).Return(&songRepoDTO.GetTextResponse{Text: testText}, nil)

		testCases := []struct {
			name           string
			request        songServiceDTO.GetTextRequest
			expectedVerses []string
		}{
			{
				name: "first page",
				request: songServiceDTO.GetTextRequest{
					ID: 1,
					Filter: songServiceDTO.GetTextFilter{
						Page:     1,
						PageSize: 2,
					},
				},
				expectedVerses: []string{
					"Verse 1\nLine 1\nLine 2",
					"Verse 2\nLine 1",
				},
			},
			{
				name: "second page",
				request: songServiceDTO.GetTextRequest{
					ID: 1,
					Filter: songServiceDTO.GetTextFilter{
						Page:     2,
						PageSize: 2,
					},
				},
				expectedVerses: []string{
					"Verse 3\nLine 1\nLine 2\nLine 3",
					"Verse 4",
				},
			},
			{
				name: "page out of range",
				request: songServiceDTO.GetTextRequest{
					ID: 1,
					Filter: songServiceDTO.GetTextFilter{
						Page:     3,
						PageSize: 2,
					},
				},
				expectedVerses: []string{},
			},
			{
				name: "page with big pagesize",
				request: songServiceDTO.GetTextRequest{
					ID: 1,
					Filter: songServiceDTO.GetTextFilter{
						Page:     1,
						PageSize: 10,
					},
				},
				expectedVerses: []string{
					"Verse 1\nLine 1\nLine 2",
					"Verse 2\nLine 1",
					"Verse 3\nLine 1\nLine 2\nLine 3",
					"Verse 4",
				},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				response, err := service.GetText(ctx, tc.request)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedVerses, response.Verses)
			})
		}
	})

	t.Run("Error from repository", func(t *testing.T) {
		repoMock.GetTextMock.Expect(ctx, songRepoDTO.GetTextRequest{ID: 1}).Return(nil, fmt.Errorf("repository error"))

		request := songServiceDTO.GetTextRequest{
			ID: 1,
			Filter: songServiceDTO.GetTextFilter{
				Page:     1,
				PageSize: 2,
			},
		}
		response, err := service.GetText(ctx, request)
		assert.Error(t, err)
		assert.Nil(t, response)

	})
}
