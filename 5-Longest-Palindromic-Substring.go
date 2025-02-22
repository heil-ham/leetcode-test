func longestPalindrome(s string) string {
    longestPalin := string(s[0])
    
    
    for i := 0; i<len(s); i++ {
        
        if len(s) > 2 && (len(longestPalin) > len(s)-i) {
            return longestPalin
        }
        substring := string(s[i])
        for j := i+1; j<len(s); j++ {
            substring += string(s[j])
            if isPalindrome(substring) && (len(substring) > len(longestPalin)) {
                longestPalin = substring
            }
        }
    }

    return longestPalin
}

func isPalindrome(s string) bool {
    invertedIndex := len(s)-1
    
    for i:=0; i<int(len(s)/2); i++ {
        if s[i] != s[invertedIndex] {
            return false
        }
        invertedIndex--
    }
    
    return true
}