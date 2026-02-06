package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRouter(t *testing.T) {
	// start router
	go Router(attachChan, chatChan, attachment, chatContent)

	task := "this is a task, finish it"
	fmt.Println("now sending task")
	if !SendChanData(attachChan, task, 10*time.Second) {
		fmt.Println("fail to send task")
	}

	time.Sleep(15 * time.Second)

	chat := "drink after work"
	fmt.Println("now sending chat")
	if !SendChanData(chatChan, chat, 10*time.Second) {
		fmt.Println("fail to send chat")
	}

	close(attachChan)
	close(chatChan)
	fmt.Println("close all chan")
}
