package main

import "fmt"

type player interface {
	play()
}
type boy struct {
}

func (b boy) play() { fmt.Println("balloon") }
