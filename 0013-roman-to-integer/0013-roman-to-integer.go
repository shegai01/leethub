func romanToInt(s string) int {
    if len(s) < 1{
        return 0
    }
    res := 0
    table := map[byte]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }

for i:= 0; i < len(s);i++{
    if i < len(s)-1 && table[s[i]]< table[s[i+1]]{
        res-=table[s[i]]
    }else{
    res+=table[s[i]]
    }
}
    return res
}