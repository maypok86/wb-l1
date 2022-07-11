package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

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

	// также горутины можно останавливать с помощью signal.Notify,
	// как продемонстрировано в предыдущем задании

	/*
		Не очень понятно, что имелось в виду в задании.
		Горутины останавливаются после завершения работы либо с паникой
	*/

	wg.Wait()
}
