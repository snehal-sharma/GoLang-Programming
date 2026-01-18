// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputStr string
	inputStr = "Split this string"

	//fmt.Println("Enter String to Split")
	//fmt.Scanf("%s",ß&inputStr)
	//fmt.Scanln(ß&inputStr)

	strSplit := strings.Split(inputStr, "")
	fmt.Println(strSplit)

	for i, val := range strSplit {
		fmt.Printf("\nIndex : %d Value: %s Type :%T", i, val, val)
	}

	//This approach converts the string to a slice of rune values, which are integer aliases representing Unicode code points.

	runes := []rune(inputStr)
	fmt.Println("\nRunes : ", runes)

	var runeStr []string

	for i, val := range runes {
		fmt.Printf("\nIndex : %d Value: %v Type :%T", i, val, val)
		runeStr = append(runeStr, string(val))
	}

	fmt.Println("\n Rune to String : ", runeStr)
}
