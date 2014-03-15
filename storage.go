package main

import (
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)


const CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const ID_LENGTH = 6


func init() {
	rand.Seed(time.Now().UnixNano())
}


type Storage struct {
	root string
}

func NewStorage(root string) *Storage {
	return &Storage{root}
}

func (s *Storage) Create() (id string, dest io.WriteCloser, err error) {
	// create id
	var idbytes = make([]byte, ID_LENGTH)
	for idx := 0; idx < ID_LENGTH; idx++ {
		idbytes[idx] = CHARS[rand.Intn(len(CHARS))]
	}
	id = string(idbytes)

	// create file storage
	path := s.GetPath(id)
	baseDir := filepath.Dir(path)
	if _, err = os.Stat(baseDir); os.IsNotExist(err) {
		err = os.MkdirAll(baseDir, 0777)
		if err != nil {
			return "", nil, err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return "", nil, err
	}

	return id, file, nil
}

func (s *Storage) Open(id string) (io.ReadCloser, error) {
	path := s.GetPath(id)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *Storage) GetPath(id string) string {
	return filepath.Join(s.root, id[:2] + "/" + id[2:])
}
