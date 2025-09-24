package box

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type StoreAdapter interface {
	Store([]byte) (int64, error)
	Fetch(id int64) ([]byte, error)
}

type MapStore struct {
	m map[int64][]byte
}

func NewMapStore() *MapStore {
	return &MapStore{m: map[int64][]byte{}}
}

// TODO: generate a better ID
func (s *MapStore) Store(data []byte) (int64, error) {
	id := time.Now().UnixNano()
	s.m[id] = data
	return id, nil
}
func (s *MapStore) Fetch(id int64) ([]byte, error) {
	return s.m[id], nil
}

type FileStore struct {
	storePath string
}

func NewFileStore(path string) *FileStore {
	os.MkdirAll(path, os.ModePerm)
	return &FileStore{storePath: path}
}

// TODO: generate a better ID
func (s *FileStore) Store(data []byte) (int64, error) {
	id := time.Now().UnixNano()
	fileName := filepath.Join(s.storePath, fmt.Sprintf("%d", id))
	err := os.WriteFile(fileName, data, 0666)
	return id, err
}

func (s *FileStore) Fetch(id int64) ([]byte, error) {
	fileName := filepath.Join(s.storePath, fmt.Sprintf("%d", id))
	return os.ReadFile(fileName)
}
