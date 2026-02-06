package main

import (
	"fmt"
	"testing"
)

func TestPerson(t *testing.T) {
	person := Person{Name: "John", Age: 20, Email: "john@example.com", Tele: "1234567890"}
	fmt.Printf("addr:%p.\n", &person)
	p := &person
	p.ShowInfo()
}

func TestPerson_IShow(t *testing.T) {
	var show IShow = &Person{Name: "John", Age: 20, Email: "john@example.com", Tele: "1234567890"}
	fmt.Printf("addr:%p.\n", show)
	show.ShowInfo()
}
