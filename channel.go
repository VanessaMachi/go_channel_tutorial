package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func test(messages chan (string), str string) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		messages <- str
	}()
	wg.Wait()
	close(messages)

}

func wow(messages chan (string), str string) {
	messages <- str
}

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}
func main() {

	messages := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				messages <- result
			}()
		}
		wg.Wait()
		close(messages)
	}()
	for n := range messages {
		fmt.Printf("n = %d\n", n)
	}
}
