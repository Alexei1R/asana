package storage_test

import (
	"os"
	"testing"

	"asana/pkg/storage"
)

type User struct {
	Name string
	Gid  string
}

func TestWriteJson(t *testing.T) {
	dir := t.TempDir()
	data := User{Name: "Alice", Gid: "123"}

	filename, err := storage.WriteJson(dir, "user", data)
	if err != nil {
		t.Fatalf("WriteJson failed: %v", err)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("File was not created: %s", filename)
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if len(content) == 0 {
		t.Errorf("File content is empty")
	}
}
