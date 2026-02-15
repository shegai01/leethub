func moveZeroes(nums []int)  {
    if len(nums) <1 {
        return 
    }
    count := 0
    for i := 0; i < len(nums); i++{
        if nums[i] != 0{
            nums[i], nums[count] = nums[count], nums[i]
            count++
        }
    }
}
// memory: O(1); time O(n)
