package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words :=strings.Split(s," ")
	fmt.Println(words)
	WordCounter :=make(map[string]int,len(words))
	for _, word := range words{
		value, ok := WordCounter[word]
		if ok{
			WordCounter[word] = value+1
		}else{
			WordCounter[word] = 1
		}
	
	}
	return WordCounter
}
