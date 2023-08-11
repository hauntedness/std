package syncs

import "sync"

// WaitGroup is simalar to sync.WaitGroup but mixin a channel
// to limit parallelism
type WaitGroup struct {
	wg   sync.WaitGroup
	Chan chan struct{}
}

// Add will firstly apply channel send
// and then call underlying WaitGroup.Add(1)
// it blocks if too many goroutine call Add
func (wg *WaitGroup) Add() {
	wg.Chan <- struct{}{}
	wg.wg.Add(1)
}

// Done will firstly apply channel receive
// and then call underlying WaitGroup.Done()
// it allow some goroutine unblock
func (wg *WaitGroup) Done() {
	<-wg.Chan
	wg.wg.Done()
}

func (wg *WaitGroup) Wait() {
	wg.wg.Wait()
}
