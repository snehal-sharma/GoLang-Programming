/*
Example 1:

Input: n = 4
Output: 4
Explanation:
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4
*/
package main

import "fmt"

func main() {
	n := 5
	value := tribonacci(5)
	fmt.Printf("\nThe %d number in the sequence is %d", n, value)
}
func tribonacci(n int) int {

	first := 0
	second := 1
	third := 1
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		fmt.Println("The Tribonacci Sequence is : ")
		fmt.Printf("%d \t %d \t %d", first, second, third)
		for i := 3; i <= n; i++ {
			first, second, third = second, third, first+second+third
			fmt.Printf("\t %d", third)
		}
		return third
	}
}
