package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int, data chan int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case n := <-data:
			fmt.Println("worker id:", id)
			fmt.Println("data:", n)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: go run t4/main.go <number of workers>")
		return
	}
	timeout, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("'%s' is not a number", os.Args[1])
		return
	}
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	data := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	const numberOfWorkers = 5
	wg.Add(numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		go worker(ctx, &wg, i, data)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("time's up")
			return
		case <-quit:
			cancel()
			wg.Wait()
			fmt.Println("all workers finished")
			return
		default:
			data <- rand.Int()
		}
	}
}
