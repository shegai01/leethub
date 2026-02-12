func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)


    for i, v := range nums{
        res := target - v
        if ind, ok := hash[res]; ok{
            return []int{ind, i}
        }
        hash[v] = i
    }
    return nil
}