func removeDuplicates(nums []int) int {
    count := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[count] {
            count++
            nums[count] = nums[i]
        }
    }
    return count + 1
}


func containsDuplicate(nums []int) bool {
    unique := make(map[int]bool)
    for _, num := range nums {
        if unique[num] {
            return true
        }
        unique[num] = true
    }
    return false
}


func reverseString(s []byte)  {
    left, right := 0, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}
