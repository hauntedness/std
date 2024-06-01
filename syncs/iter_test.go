package syncs

import (
	"fmt"
	"testing"
	"time"
)

func TestWaitAll(t *testing.T) {
	start := time.Now()
	for i := range WaitAllWithLimitG(7, 3) {
		time.Sleep(time.Second * 1)
		fmt.Println(i)
	}
	duration := time.Since(start)
	if sec := duration.Seconds(); int(sec) != 3 {
		t.Fatalf("should take around 3 seconds, actual take %v seconds", sec)
	}
}
