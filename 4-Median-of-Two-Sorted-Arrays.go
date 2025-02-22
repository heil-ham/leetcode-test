func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // appending 2 slices
    combinedNums := append(nums1, nums2...)
    
    // sorting logic
    // cara 1
    // for i := 0; i < len(combinedNums); i++ {
    //     for j:=0; j < len(combinedNums); j++ {
    //         if j == len(combinedNums)-1 {
    //             break
    //         } else if combinedNums[j] > combinedNums[j+1] {
    //             temp := combinedNums[j]
    //             combinedNums[j] = combinedNums[j+1]
    //             combinedNums[j+1] = temp
    //             j = 0
    //         }
            
    //     }
    // }

    // cara 2
    for i := 0; i < len(combinedNums); i++ {
        tiniestIndex := i
        for j:=i; j < len(combinedNums); j++ {
            if combinedNums[j] < combinedNums[tiniestIndex] {
                tiniestIndex = j
            }
        }
        temp := combinedNums[i]
        combinedNums[i] = combinedNums[tiniestIndex]
        combinedNums[tiniestIndex] = temp
    }
    
    fmt.Println(combinedNums)
    
    var median float64
    
    if len(combinedNums) % 2 == 0 {
        median = float64(combinedNums[len(combinedNums)/2] + combinedNums[(len(combinedNums)/2) - 1]) / 2
    } else {
        index := int(len(combinedNums)/2)
        median = float64(combinedNums[index])
    }
    return median
}