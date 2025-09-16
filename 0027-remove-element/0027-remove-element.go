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
//Memory: 3.97 MB, Beats: 99.43%
