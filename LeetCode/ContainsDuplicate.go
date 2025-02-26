/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Input: nums = [1,2,3,1]
Output: true

Input: nums = [1,2,3,4]
Output: false
*/

func containsDuplicate(nums []int) bool {
    duplicate := make(map[int]bool,len(nums))
    for i:=0;i<len(nums);i++{
        _,ok:=duplicate[nums[i]]
        if ok{
            return true
        }else{
            duplicate[nums[i]]=true
        }
    }
    return false
}
