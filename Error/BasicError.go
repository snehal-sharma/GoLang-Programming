package main

import "fmt"

func ThrowError() error {
	return fmt.Errorf("Encountered an error while processing")
}

func main() {
	err := ThrowError()
	if err != nil {
		fmt.Println(err)
	}
}
