func twoSum(nums []int, target int) []int {
    xMap := make(map[int]int)
    for ind, val := range nums{
        key := target - val
        if i, ok := xMap[key]; ok{
            return []int{i, ind}
        }
        xMap[val] = ind
    }
    return []int{}
    }
