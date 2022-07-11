package main

import (
	"fmt"
	"time"
)

func Sleep1(t time.Duration) {
	// похожим образом можно сделать с помощью контекста, но, на мой взгляд, это лишнее
	<-time.After(t) // aka <-time.NewTimer(t).C
}

func Sleep2(t time.Duration) {
	<-time.NewTicker(t).C
}

func testSleep(t time.Duration, sleep func(time.Duration)) {
	ts := time.Now().Second()
	sleep(t)
	fmt.Println(time.Now().Second() - ts)
}

func main() {
	testSleep(time.Second, Sleep1)
	testSleep(time.Second, Sleep2)
}
