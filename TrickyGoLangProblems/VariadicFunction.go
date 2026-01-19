//Create a variadic function which accepts any input and calculates the sum of all integer arguments passed to it.

package main

import (
	"fmt"
	"reflect"
)

// func VarSum(varInter ...any)
func VarSum(varInter ...interface{}) {
	sum := 0
	for _, value := range varInter {
		if reflect.TypeOf(value).Kind() == reflect.Int {
			fmt.Printf("\nValue : %v is of Type : %T", value, value)
			sum += value.(int)
		}
	}
	fmt.Printf("\n Sum : %d", sum)
}

func main() {
	VarSum(1, 2, 3, 4)
	VarSum(4, 6, "a")
	VarSum("a", "b", "c")
}
