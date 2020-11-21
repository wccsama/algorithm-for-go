package concurrent

import (
	"sync"
)

type con struct {
	mux sync.Mutex
	muxRW sync.RWMutex
	w sync.WaitGroup{}
	make(chan int)
	Chan <-chan int
	chan chan int-> 
	select
	ctx
}
