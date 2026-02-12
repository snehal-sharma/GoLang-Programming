package main

import (
	"errors"
	"fmt"
)

var ErrorNotFound = errors.New("Item not found")

func getItem(id int) error {
	if id == 0 {
		return ErrorNotFound
	}
	return nil
}

func main() {
	err := getItem(0)
	if errors.Is(err, ErrorNotFound) {
		fmt.Println(err)
	}
}
