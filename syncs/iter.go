package syncs

import (
	"sync"
)

// WaitAll Wrap [sync.WaitGroup]
//
// note that you should not break the loop func
func WaitAll(total int) func(yield func(i int) bool) {
	return func(yield func(i int) bool) {
		wg := &sync.WaitGroup{}
		wg.Add(total)
		for i := range total {
			go func() {
				defer wg.Done()
				yield(i)
			}()
		}
		wg.Wait()
	}
}

// WaitAllWithLimitG is similar to WaitAll but also limits the number of active goroutines in this group to at most limit.
func WaitAllWithLimitG(total int, limit int) func(yield func(i int) bool) {
	return func(yield func(i int) bool) {
		ch := make(chan struct{}, limit)
		wg := &sync.WaitGroup{}
		wg.Add(total)
		for i := range total {
			ch <- struct{}{}
			go func() {
				defer func() {
					wg.Done()
					<-ch
				}()
				yield(i)
			}()
		}
		wg.Wait()
		close(ch)
	}
}
