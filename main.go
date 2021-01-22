package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Session 1")
		time.Sleep(2 * time.Second)
	}()
	// fmt.Println("Session 1")
	// time.Sleep(2 * time.Second)
	fmt.Println("Session 2")
}
