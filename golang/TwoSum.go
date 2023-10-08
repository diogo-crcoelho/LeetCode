func twoSum(nums []int, target int) []int {
    n := len(nums)
    for i := 0; i < n - 1; i++{
        for j := i + 1; i < n; j++{
            if nums[i] + nums[j] == target{
                return []int{i , j}
            }
        }
    }
    return []int{-1, -1}
}
