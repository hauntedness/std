package osx

import (
	"errors"
	"os"
)

// CreateFor Create file via call [os.Create] and then call handler f and finally close the file
func CreateFor[T any](filename string, f func(file *os.File) (T, error)) (t T, err error) {
	var file *os.File
	file, err = os.Create(filename)
	if err != nil {
		return
	}
	defer func() {
		closeerr := file.Close()
		if closeerr != nil {
			err = errors.Join(err, closeerr)
		}
	}()
	t, err = f(file)
	return
}

// OpenFor Open file via call [os.Open] and then call handler f and finally close the file
func OpenFor[T any](filename string, f func(file *os.File) (T, error)) (t T, err error) {
	var file *os.File
	file, err = os.Open(filename)
	if err != nil {
		return
	}
	defer func() {
		closeerr := file.Close()
		if closeerr != nil {
			err = errors.Join(err, closeerr)
		}
	}()
	t, err = f(file)
	return
}
