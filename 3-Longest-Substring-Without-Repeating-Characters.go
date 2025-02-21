func lengthOfLongestSubstring(s string) int {
    if s == \\ {
        return 0
    }
    appendedString := string(s[0])
    longest := 1
    currentLong := 1
    
    for i := 0; i < len(s); i++ {
        for j := i+1; j < len(s); j++ {
            if alreadyExist(appendedString, string(s[j])) {
                if currentLong > longest {
                    longest = currentLong
                }
                currentLong = 1
                appendedString = string(s[i+1])
                break
            } else {
                appendedString +=  string(s[j])
                currentLong++
            }
            
            if j == len(s)-1 && currentLong > longest {
                longest = currentLong
                return longest
            }
        }
    }
    
    return longest
}

func alreadyExist(s string, char string) bool {
    for i := 0; i < len(s); i++ {
        if char == string(s[i]) {
            return true
        }
    }

    return false
}