package main

import "fmt"
import "utils/uuid"
import uuid2 "github.com/eiblog/utils/uuid"

func main() {
	fmt.Println("Hello")
	fmt.Println(uuid.NewV4())
	fmt.Println(uuid2.NewV4())
}
