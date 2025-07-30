package storage

import (
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	Dir string
}

func NewLocalStorage(dir string) *LocalStorage {
	os.MkdirAll(dir, os.ModePerm)
	return &LocalStorage{Dir: dir}
}

func (s *LocalStorage) Save(name string, r io.Reader) error {
	filePath := filepath.Join(s.Dir, name)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, r)
	return err
}

func (s *LocalStorage) Load(filename string) (string, error) {
	filePath := filepath.Join(s.Dir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", err
	}
	return filePath, nil
}
