package trello

import (
	"fmt"
)

const (
	ENDPOINT = "https://api.trello.com/1"
)

type Trello struct {
	Key   string
	Token string
}

func NewTrello(key, token string) (t *Trello) {
	t = new(Trello)
	t.Key = key
	t.Token = token
	return t
}

func (t *Trello) GetBoards() {
	url := ENDPOINT + "/members/me/boards"
	fmt.Println(url)
}
