func isValid(s string) bool {
    stack := []rune{}
    hash := map[rune]rune{
        ')':'(',
        '}':'{',
        ']':'[',
    }
    for _, val := range s{
        switch val {
            case '(', '{', '[':
            stack = append(stack, val)
            case ')', '}', ']':
            if len(stack) == 0 || stack[len(stack)-1] != hash[val]{
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}