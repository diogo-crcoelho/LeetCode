func revNum(x int) int {
    res := 0
    for x > 0{
        res = res * 10 + x % 10
        x = x / 10
    }
    return res
}

func isPalindrome(x int) bool {
    if (x < 0){
        return false
    } else if revNum(x) == x{
        return true
    }
    return false
}
