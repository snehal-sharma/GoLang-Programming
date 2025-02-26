//Re-attempt

type Vector struct{
    index1, index2 int
    visited bool
}
func containsNearbyDuplicate(nums []int, k int) bool {
    if len(nums)==1{
        return false
    }
    duplicate := make(map[int]Vector,len(nums))
    for i:=0;i<len(nums);i++{
        fmt.Println(duplicate[nums[i]])
        _,ok:=duplicate[nums[i]]
        if !ok{
            fmt.Println("Setting Default Values for",nums[i])
            duplicate[nums[i]]=Vector{-1,-1,false}
        }
        val,_:=duplicate[nums[i]]
        if val.index1!=-1{
            if val.visited{
                if val.index2-val.index1<=k{
                    return true
                }
                duplicate[nums[i]]=Vector{val.index2,i,true}
            }else{
                duplicate[nums[i]]=Vector{val.index1,i,true}
            }
        }else{
        duplicate[nums[i]]=Vector{i,-1,false}
        }
        fmt.Printf("%+v\n",duplicate[nums[i]])
    }
    fmt.Println(duplicate)
    for _,val := range duplicate{
            if val.visited==true && val.index2-val.index1 <= k{
                return true
            }
    }
    return false
}
