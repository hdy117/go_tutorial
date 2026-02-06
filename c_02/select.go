package main

import (
	"fmt"
	"time"
)

type AttachmentChan = chan string
type ChatContentChan = chan string

var attachChan = make(AttachmentChan)
var chatChan = make(ChatContentChan)

type IPayload interface {
	GetPayload() string
	SetPayload(payload *string)
}

type Attachment struct {
	FromUserID string
	ToUserID   string
	Payload    string
}

func (obj *Attachment) GetPayload() string {
	fmt.Printf("this is attachment payload %s from %s to %s.\n", obj.Payload, obj.FromUserID, obj.ToUserID)
	return obj.Payload
}

func (obj *Attachment) SetPayload(payload *string) {
	obj.Payload = *payload
	fmt.Printf("set attachment payload %s from %s to %s.\n", obj.Payload, obj.FromUserID, obj.ToUserID)
}

type Message struct {
	FromUserID string
	ToUserID   string
	Content    string
}

func (obj *Message) GetPayload() string {
	fmt.Printf("this is message content %s from %s to %s.\n", obj.Content, obj.FromUserID, obj.ToUserID)
	return obj.Content
}

func (obj *Message) SetPayload(payload *string) {
	obj.Content = *payload
	fmt.Printf("set content %s from %s to %s.\n", obj.Content, obj.FromUserID, obj.ToUserID)
}

var attachment IPayload = &Attachment{FromUserID: "Jim", ToUserID: "Lily", Payload: "A task specification"}
var chatContent IPayload = &Message{FromUserID: "Jim", ToUserID: "Lily", Content: "Drink after work."}

// SendChanData, template to send chan data
func SendChanData[T ~chan TData, TData string | int | float32 | float64](channel T, data TData, timeout time.Duration) bool {
	if channel == nil {
		fmt.Println("chan is nil")
		return false
	}

	select {
	case channel <- data:
		{
			return true
		}
	case <-time.After(timeout):
		{
			return false
		}
	}
}

// Process user chat
func Router(attachmentChan AttachmentChan, chatContentChan ChatContentChan, attachment IPayload, chat IPayload) {
	fmt.Println("router start of chat")
loop:
	for {
		select {
		case attachmentPayload, ok := <-attachmentChan:
			{
				if ok {
					attachment.SetPayload(&attachmentPayload)
					attachment.GetPayload()
				}
			}
		case chatContent, ok := <-chatContentChan:
			{
				if ok {
					chat.SetPayload(&chatContent)
					chat.GetPayload()
				}
			}
		case <-time.After(time.Second * 10):
			{
				fmt.Println("timeout no new info, quit.")
				break loop
			}
		}

	}

	fmt.Println("router quit, end of chat")
}
