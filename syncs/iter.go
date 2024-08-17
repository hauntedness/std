package syncs

import (
	"sync"
)

type latch struct {
	total int
}

func Latch(total int) latch {
	return latch{total: total}
}

func (l latch) ForEach(f func(i int)) {
	wg := &sync.WaitGroup{}
	wg.Add(l.total)
	for i := range l.total {
		go func() {
			defer wg.Done()
			f(i)
		}()
	}
	wg.Wait()
}

type latch2 struct {
	total     int
	parallism int
}

func Latch2(total, parallism int) latch2 {
	return latch2{total: total, parallism: parallism}
}

func (l latch2) ForEach(f func(i int)) {
	wg := &sync.WaitGroup{}
	ch := make(chan struct{}, l.parallism)
	wg.Add(l.total)
	for i := range l.total {
		ch <- struct{}{}
		go func() {
			defer func() {
				<-ch
				wg.Done()
			}()
			f(i)
		}()
	}
	wg.Wait()
}
