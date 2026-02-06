package main

import (
	"fmt"
	"sync"
	"time"
)

// IntChan means a channel of int type
type IntChan = chan int

// global variable of channel
var intChan IntChan

func RecvData(ch IntChan) {
	for data := range ch {
		fmt.Println(time.Now().Local().Format(time.Layout))
		fmt.Printf("recv data:%d.\n", data)
	}
}

func SendData(ch IntChan, wg *sync.WaitGroup) {
	// notify one finish of send
	defer wg.Done()

	duration := 2 * time.Second
	for i := range 3 {
		// send to channel
		ch <- i
		// sleep
		time.Sleep(duration)
	}
}
