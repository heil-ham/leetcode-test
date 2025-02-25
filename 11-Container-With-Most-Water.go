func maxArea(height []int) int {
    i := 0
    j := len(height)-1
    maxVolume := 0
    volume := 0

    for i < j {
        if height[i] <= height[j] {
            volume = height[i] * (j-i)
            i++
        } else {
            volume = height[j] * (j-i)
            j--
        }

        if volume > maxVolume {
            maxVolume = volume
        }
    }

    return maxVolume
}