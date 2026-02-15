func removeElement(nums []int, val int) int {
    if len(nums) == 0{
        return 0
    }
    count := 0
    for i := 0; i < len(nums); i++{
        if nums[i] != val{
            nums[count] = nums[i]
            count++
        }
    }
    return count
}
//Memory: O(1), time: O(n)
