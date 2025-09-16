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
// memory: O(n), time compl : O(n)
