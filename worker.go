package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kollinchu/reife-worker/trello"
)

func main() {
	for {
		t := trello.NewTrello("111", "1111")
		fmt.Println(t)
		time.Sleep(600)
		trello_key := os.Getenv("TRELLO_KEY")
		fmt.Println(trello_key)
		fmt.Println("ending.....")
	}
}
