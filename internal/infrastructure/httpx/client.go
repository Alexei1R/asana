package httpx

import (
	"asana/pkg/config"
	"asana/pkg/log"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	Http *http.Client
	cfg  *config.Config
}

func New() *Client {
	return &Client{
		Http: http.DefaultClient,
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
	maxRetries := c.cfg.Refresh.Retry
	wait := c.cfg.Refresh.Interval

	for i := 0; i < maxRetries; i++ {
		req, err := c.newRequest(ctx, method, path, params, nil)
		if err != nil {
			return nil, err
		}

		resp, err := c.Http.Do(req)
		if err != nil {
			return nil, err
		}

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode == 429 {
			retryAfter := wait
			if ra := resp.Header.Get("Retry-After"); ra != "" {
				if sec, err := strconv.Atoi(ra); err == nil {
					retryAfter = time.Duration(sec) * time.Second
				}
			}

			log.Info("Rate limited, waiting %s before retry", retryAfter)
			time.Sleep(retryAfter)
			wait *= 2 // exponential backoff
			continue
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return body, nil
		}

		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	return nil, fmt.Errorf("max retries exceeded for %s %s", method, path)
}
