package httpx

import (
	"asana/pkg/config"
	"asana/pkg/log"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	http *http.Client
	cfg  *config.Config
}

func New() *Client {
	return &Client{
		http: http.DefaultClient,
		cfg:  config.Get(),
	}
}

func (c *Client) newRequest(ctx context.Context, method, path string, params url.Values, body any) (*http.Request, error) {
	fullURL := fmt.Sprintf("%s%s?%s", c.cfg.Asana.BaseURL, path, params.Encode())

	var req *http.Request
	var err error
	if body != nil {
		log.Error("Request body is not nil, but body handling is not implemented")
	} else {
		req, err = http.NewRequestWithContext(ctx, method, fullURL, nil)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Set("Authorization", "Bearer "+c.cfg.Asana.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (c *Client) Do(ctx context.Context, method, path string, params url.Values) ([]byte, error) {
	req, err := c.newRequest(ctx, method, path, params, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
