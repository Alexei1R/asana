package projects_test

import (
	"encoding/json"
	"testing"

	"asana/internal/domain/projects"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalProjects(t *testing.T) {
	data := []byte(`{"data":[{"gid":"1","name":"Project 1"},{"gid":"2","name":"Project 2"}]}`)
	var resp struct {
		Data []*projects.Project `json:"data"`
	}
	err := json.Unmarshal(data, &resp)
	assert.NoError(t, err)
	assert.Len(t, resp.Data, 2)
	assert.Equal(t, "Project 1", resp.Data[0].Name)
	assert.Equal(t, "Project 2", resp.Data[1].Name)
}
