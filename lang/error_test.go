package lang_test

import (
	"log"
	"testing"

	"github.com/hauntedness/std/lang"
)

func TestTracedError_Error(t *testing.T) {
	err1 := lang.NewError("error1")
	err2 := lang.Error(err1, "error2")
	err3 := lang.Error(err2, "error3")

	actual := err3.Error()
	if actual != "error3: error2: error1" {
		t.Fatal(err3)
	}

	log.Printf("%+v", err3)
}
