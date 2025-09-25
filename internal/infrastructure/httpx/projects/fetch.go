package projects

import (
	"asana/internal/domain/projects"
	"asana/internal/infrastructure/httpx"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type HttpProjects struct {
	client *httpx.Client
}

func New(client *httpx.Client) projects.Repository {
	return &HttpProjects{
		client: client,
	}
}

type getProjectsResponse struct {
	Data []*projects.Project `json:"data"`
}

func (hu *HttpProjects) GetProjects(ctx context.Context) ([]*projects.Project, error) {
	params := url.Values{}
	data, err := hu.client.Do(ctx, http.MethodGet, "/projects", params)
	if err != nil {
		return nil, err
	}

	var resp getProjectsResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
