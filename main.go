package main

import (
  gomock "github.com/golang/mock/gomock"

  "fmt"
)

func main() {
  fmt.Println("Hello world")
  c := gomock.Controller{}
  fmt.Printf("%s", c)
}

