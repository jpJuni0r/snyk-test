package main

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/r3labs/diff/v3"
)

func main() {
	fmt.Println("Hello world")
	c := gomock.Nil()
	d := diff.DiffError{}
	fmt.Printf("%s", c)
	fmt.Printf("%s", d)
}
