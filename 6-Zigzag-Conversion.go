func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    up := false
    rowTracker := 1
    
    stringContainer := make([]string, numRows)
    
    stringContainer[0] = string(s[0])
    
    for i := 1; i < len(s); i++ {
        if up {
            stringContainer[numRows - rowTracker - 1] += string(s[i])
        } else {
            stringContainer[rowTracker] += string(s[i])
        }
        
        if rowTracker == (numRows-1) {
            rowTracker = 1
            up = !up
        } else {
            rowTracker++
        }
        
    }
    
    s = \\
    for i := 0; i<len(stringContainer); i++ {
        s += string(stringContainer[i])
    }
    
    return s
}