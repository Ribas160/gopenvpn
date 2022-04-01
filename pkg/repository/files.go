package repository

import (
	"errors"
	"io"
	"os"
)

type FilesRepository struct {
}

func newFilesRepository() *FilesRepository {
	return &FilesRepository{}
}

func (r *FilesRepository) Copy(from string, to string) error {
	buf, err := r.Read(from)
	if err != nil {
		return err
	}

	dFile, err := os.Create(to)
	if err != nil {
		return errors.New("New file can't be created: " + err.Error())
	}

	defer dFile.Close()

	dFile.WriteString(buf)

	return nil
}

func (r *FilesRepository) Read(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", errors.New("File can't be opened: " + err.Error())
	}

	stat, err := f.Stat()
	if err != nil {
		return "", err
	}

	defer f.Close()

	buf := make([]byte, stat.Size())
	for {
		_, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}

	return string(buf), nil
}
