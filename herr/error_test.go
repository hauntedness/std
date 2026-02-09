package herr_test

import (
	"log"
	"testing"

	"github.com/hauntedness/std/herr"
)

func TestTracedError_Error(t *testing.T) {
	err1 := herr.New("error1")
	err2 := herr.Wrap(err1, "error2")
	err3 := herr.Wrap(err2, "error3")
	actual := err3.Error()
	if actual != "error3: error2: error1" {
		t.Fatal(err3)
	}
	log.Printf("%+v", err3)
}
