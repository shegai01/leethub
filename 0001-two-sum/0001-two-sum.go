func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    // target :=0
    for i, val := range nums{
        res := target - val
        if ind, found := hash[res]; found{
            return []int{ind, i}
        }
        hash[val]= i
    }
    return []int{}
}
//Memory: 6.11 MB, Beats: 11.06%
