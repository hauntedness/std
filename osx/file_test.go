package osx

import (
	"io"
	"os"
	"testing"
)

func TestCreateFor(t *testing.T) {
	filename := "testdata/some.txt"
	{
		actual, err := CreateFor(filename, func(file *os.File) (string, error) {
			_, err := file.WriteString("some text")
			return "success", err
		})
		if err != nil {
			t.Fatal(err)
		}
		if actual != "success" {
			t.Fatalf("not success")
		}
	}
	{
		actual, err := OpenFor(filename, func(file *os.File) (string, error) {
			data, err := io.ReadAll(file)
			if err != nil {
				return "", err
			}
			return string(data), nil
		})
		if err != nil {
			t.Fatal(err)
		}
		if actual != "some text" {
			t.Fatalf("file content")
		}
	}
}
