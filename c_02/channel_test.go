package main

import (
	"sync"
	"testing"
)

func TestChannel(t *testing.T) {
	// create channel, must create since intChan is a declaration only
	// intChan will be nil if not make this chan, and send and recv will be blocked forever
	intChan = make(IntChan)
	var wg *sync.WaitGroup = &sync.WaitGroup{}

	// finally close channel
	defer close(intChan)

	// async send data
	wg.Add(1)
	go SendData(intChan, wg)
	wg.Add(1)
	go SendData(intChan, wg)

	// wait and print data
	go RecvData(intChan)

	// wait
	wg.Wait()
}
