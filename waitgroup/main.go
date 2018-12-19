package main

import (
	"fmt"
	"sync"
	"time"
)

func sleep(sec time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(sec * time.Second)
	fmt.Println("exit sleep")
}

func main() {
	fmt.Println("Main goroutine start")
	var wg sync.WaitGroup
	wg.Add(2)
	go sleep(2, &wg)
	go sleep(3, &wg)
	wg.Wait()
	fmt.Println("Main gorutine exit")
}
