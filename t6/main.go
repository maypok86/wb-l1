package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		fmt.Println("default cancel")
	}()
	time.Sleep(time.Second)
	cancel()

	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		fmt.Println("cancel by deadline")
	}()

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		fmt.Println("cancel by timeout")
	}()

	wg.Wait()
}
