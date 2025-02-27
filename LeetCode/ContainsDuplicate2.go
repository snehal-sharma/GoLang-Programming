/*
Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

Input: nums = [1,2,3,1], k = 3
Output: true

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/

func containsNearbyDuplicate(nums []int, k int) bool {
    duplicates:=make(map[int]int,len(nums))
    for i:=0;i<len(nums);i++{
        value,ok:=duplicates[nums[i]]
        if ok && i-value<=k{
            return true
        }else{
            duplicates[nums[i]]=i
        }
    }
    return false
}
