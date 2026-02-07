package main

import (
	"fmt"
	"os"
)

func OpenFile() error {
	_, err := os.Open("file.txt")
	if err != nil {
		return fmt.Errorf("Encountered an error : %w:", err)
	}
	return nil
}

func main() {
	err := OpenFile()
	if err != nil {
		fmt.Println(err)
	}
}
