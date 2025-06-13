package utils

import (
	"errors"
	"os"
)

func CheckFolder(folder string) error {
	if _, err := os.Stat(folder); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(folder, os.FileMode(0o744))
		if err != nil {
			return err
		}
	}
	return nil
}

func GetConfigFile(p string) (*os.File, error) {
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(p)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	return f, nil
}
