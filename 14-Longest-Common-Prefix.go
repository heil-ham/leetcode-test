func longestCommonPrefix(strs []string) string {
    common := \\
    smallestIndex := smallestStringIndex(strs)

    for i:=0; i<smallestIndex; i++ {
        for j:=1; j<len(strs); j++ {
            if strs[j][i] != strs[j-1][i] {
                return common
            }
        }
        common += string(strs[0][i])
    }
    return common
}

func smallestStringIndex(strs []string) int {
    smallestIndex := len(strs[0])
    for i:=1; i<len(strs); i++ {
        if len(strs[i]) < smallestIndex {
            smallestIndex = len(strs[i])
        }
    }

    return smallestIndex
}