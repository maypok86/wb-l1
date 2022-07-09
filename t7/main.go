package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Map struct {
	mutex sync.Mutex
	data  map[int]int
}

func NewMap() *Map {
	return &Map{
		data: make(map[int]int),
	}
}

func (m *Map) Set(key, value int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data[key] = value
}

func main() {
	rand.Seed(time.Now().UnixNano())

	m := NewMap()
	n := 10
	for i := 0; i < n; i++ {
		m.Set(rand.Int(), rand.Int())
	}

	fmt.Printf("%#v", m)
}
