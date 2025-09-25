package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

func CreatePath(path string) error {
	return os.MkdirAll(path, 0o777)
}

func WriteJson(dir, base string, data any) (string, error) {
	if err := CreatePath(dir); err != nil {
		return "", err
	}

	ts := time.Now().Format("20060102T150405Z")
	filename := base + "_" + ts + ".json"
	tmp := filepath.Join(dir, filename+".tmp")
	dst := filepath.Join(dir, filename)

	f, err := os.Create(tmp)
	if err != nil {
		return "", err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		f.Close()
		_ = os.Remove(tmp)
		return "", err
	}
	if err := f.Close(); err != nil {
		_ = os.Remove(tmp)
		return "", err
	}
	if err := os.Rename(tmp, dst); err != nil {
		return "", err
	}
	return dst, nil

}
