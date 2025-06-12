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
