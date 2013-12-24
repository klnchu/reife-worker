package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		time.Sleep(12)
		trello_key := os.Getenv("TRELLO_KEY")
		fmt.Println(trello_key)
		fmt.Println("ending.....")
	}
}
