package queuelist

import (
	"fmt"
	"testing"
)

func TestLinkedQueue_Add(t *testing.T) {
	var jobs = new(LinkedQueue)
	fmt.Println("To add node...")
	jobs.Add("This is node 1")
	fmt.Println("end")
}

func TestSimpleQueue(t *testing.T) {
	var jobs = new(LinkedQueue)
	fmt.Println("To add node...")
	jobs.Add("This is node 1")
	fmt.Println("node:", jobs.Peek())
}
