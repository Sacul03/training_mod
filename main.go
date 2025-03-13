package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("Starting Go Docker App...")

	for {
		fmt.Println("Hello from inside Docker!")
		log.Println("Logging: Hello from inside Docker!")

		time.Sleep(1 * time.Second) // Wait 1 second before printing again
	}
}