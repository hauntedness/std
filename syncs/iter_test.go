package syncs

import (
	"fmt"
	"testing"
	"time"
)

func TestLatch2(t *testing.T) {
	start := time.Now()
	Latch2(7, 3).ForEach(func(i int) {
		time.Sleep(time.Second * 1)
		fmt.Println(i)
	})
	duration := time.Since(start)
	if sec := duration.Seconds(); int(sec) != 3 {
		t.Fatalf("should take around 3 seconds, actual take %v seconds", sec)
	}
}
