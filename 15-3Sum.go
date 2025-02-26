func threeSum(nums []int) [][]int {
    var finalArr [][]int
    sort.Ints(nums)

    for i:=0; i<len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        left, right := i+1, len(nums)-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]

            if sum == 0 {
                finalArr = append(finalArr, []int{nums[i], nums[left], nums[right]})
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                for right > left && nums[right] == nums[right-1] {
                    right--
                }
                left++
                right--
            } else if sum > 0 {
                right--
            } else {
                left++
            }
        }

    }

    return finalArr
}