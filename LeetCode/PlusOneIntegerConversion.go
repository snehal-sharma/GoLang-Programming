/*
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Note : The below solution does not work for large Arrays.
*/

import "math"
func plusOne(digits []int) []int {
    sum:=0
    for i:=0;i<len(digits);i++{
        power:=math.Pow(10,float64(len(digits)-1-i))
        sum+=digits[i]*int(power)
    }
    sum+=1
    fmt.Printf("Final Sum %v",sum)
    var finalArray []int
    for sum!=0{
        value:=sum%10
        finalArray=append([]int{value},finalArray...)
        sum=sum/10
    }
    return finalArray
}
