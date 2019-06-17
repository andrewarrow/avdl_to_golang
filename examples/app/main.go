package main

import "fmt"
import "sync"

var mutex sync.Mutex

func main() {
	fmt.Println("example app for avdl_to_golang")

	mutex.Lock()
	fields1 := schema["Thing"]
	fields2 := schema["OtherThing"]
	mutex.Unlock()

	fmt.Println(fields1, fields2)
}
