package osx

import (
	"errors"
	"os"
)

// CreateFor Create file and then call handler f and finally close the file
func CreateFor[T any](filename string, f func(file *os.File) (T, error)) (t T, err error) {
	var file *os.File
	file, err = os.Create(filename)
	if err != nil {
		return t, err
	}
	defer func() {
		closeerr := file.Close()
		if closeerr != nil {
			err = errors.Join(err, closeerr)
		}
	}()
	t, err = f(file)
	return t, err
}

// OpenFor Open file and then call handler f and finally close the file
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
