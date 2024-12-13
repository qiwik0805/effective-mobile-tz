package music_info

import (
	"context"
	musicInfoDTO "effective-mobile-tz/internal/domain/dto/service/music_info"
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpTransport interface {
	Do(r *http.Request) (*http.Response, error)
}

type Client struct {
	httpTransport HttpTransport
	baseURL       string
}

func NewClient(httpTransport HttpTransport, baseURL string) *Client {
	return &Client{httpTransport: httpTransport, baseURL: baseURL}
}

func (c *Client) Info(ctx context.Context, r musicInfoDTO.InfoRequest) (*musicInfoDTO.InfoResponse, error) {
	endpointURL := fmt.Sprintf("%s/info?group=%s&song=%s", c.baseURL, r.Group, r.Song)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request with context: %w", err)
	}

	resp, err := c.httpTransport.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request %s: %w", endpointURL, err)
	}

	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d, expected %d", statusCode, http.StatusOK)
	}

	var response musicInfoDTO.InfoResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	return &response, nil
}
