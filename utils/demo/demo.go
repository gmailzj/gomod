package demo

import "fmt"

import "utils/uuid"
import uuid2 "github.com/eiblog/utils/uuid"
import uuid3 "github.com/google/uuid"

// Get uuid
func Get() {
	fmt.Println("Hello")
	fmt.Println(uuid.NewV4())
	fmt.Println(uuid2.NewV4())
	var uuid = uuid3.New()
	fmt.Println(uuid)
}
