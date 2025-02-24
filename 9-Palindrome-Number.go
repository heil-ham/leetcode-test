func isPalindrome(x int) bool {
    stringX := strconv.Itoa(x)
    invertedIndex := len(stringX)-1
    for i:=0; i<len(stringX)/2; i++ {
        if string(stringX[i]) != string(stringX[invertedIndex]) {
            return false
        }
        invertedIndex--
    }

    return true
}