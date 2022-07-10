package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter interface {
	Inc()
	GetValue() int64
}

type atomicCounter struct {
	value int64
}

func (ac *atomicCounter) Inc() {
	atomic.AddInt64(&ac.value, 1)
}

func (ac atomicCounter) GetValue() int64 {
	return ac.value
}

type mutexCounter struct {
	mutex sync.Mutex
	value int64
}

func (mc *mutexCounter) Inc() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	mc.value++
}

func (mc *mutexCounter) GetValue() int64 {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	return mc.value
}

func TestCounter(counter Counter, t string) {
	var wg sync.WaitGroup
	n := 100
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			counter.Inc()
		}()
	}
	wg.Wait()
	fmt.Printf("%s: %d\n", t, counter.GetValue())
}

func main() {
	TestCounter(&atomicCounter{}, "Atomic")
	TestCounter(&mutexCounter{}, "Mutex")
}
