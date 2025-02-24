
func plusOne(digits []int) []int {
    if len(digits)==1{
        if digits[0]<9{
            digits[0]++
            return digits
        }else{
            digits[0]=0
        }

    }else{
        totalLen:=len(digits)-1
        for i:=totalLen;i>=0;i--{
            fmt.Println(totalLen)
            fmt.Println(digits[i])
            if digits[i]==9{
                digits[i]=0
            }else{
                digits[i]++
                return digits
            } 
        }
    }
    digits=append([]int{1},digits...)
    return digits
}
