// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"io"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/get", nil)

	client := http.DefaultClient
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error Making a request : ", err)
		return
	}

	defer resp.Body.Close()

	body,err:=io.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("Error Reading Response")
		return
	}

	fmt.Println("Status Code : ", resp.Status)
	fmt.Println("Response : ",string(body))

}
