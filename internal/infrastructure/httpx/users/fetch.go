package users

import (
	"asana/internal/domain/users"
	"asana/internal/infrastructure/httpx"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type HttpUser struct {
	client *httpx.Client
}

func New(client *httpx.Client) users.Repository {
	return &HttpUser{
		client: client,
	}
}

type getUsersResponse struct {
	Data []*users.User `json:"data"`
}

func (hu *HttpUser) GetUsers(ctx context.Context) ([]*users.User, error) {
	params := url.Values{}

	data, err := hu.client.Do(ctx, http.MethodGet, "/users", params)
	if err != nil {
		return nil, err
	}

	var resp getUsersResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
