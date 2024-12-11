package music_info

import (
	"context"
	musicInfoDTO "effective-mobile-tz/pkg/domain/dto/service/music_info"
	"math/rand"
	"strings"
	"time"
)

type FakeClient struct {
}

func NewFakeClient() *FakeClient {
	return &FakeClient{}
}

func (f FakeClient) Info(ctx context.Context, r musicInfoDTO.InfoRequest) (*musicInfoDTO.InfoResponse, error) {
	return &musicInfoDTO.InfoResponse{
		ReleaseDate: time.Now().Format("02.01.2006"),
		Text:        generateText(4, 4, 4),
		Link:        generateRandomLink(),
	}, nil
}

func generateText(numVerses int, linesPerVerse int, wordsPerLine int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	words := []string{"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua"}
	var textBuilder strings.Builder
	for i := 0; i < numVerses; i++ {
		for j := 0; j < linesPerVerse; j++ {
			for k := 0; k < wordsPerLine; k++ {
				textBuilder.WriteString(words[rand.Intn(len(words))] + " ")
			}
			textBuilder.WriteString("\n")
		}
		textBuilder.WriteString("\n")
	}
	return textBuilder.String()
}

func generateRandomLink() string {
	const protocol = "https://"
	const domain = "example.com/"
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	const urlLen = 10
	result := make([]byte, urlLen)
	for i := 0; i < urlLen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return protocol + domain + string(result)

}
