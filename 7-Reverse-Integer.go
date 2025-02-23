func reverse(x int) int {
    minus := false

    if x < 0 {
        x = x * -1
        minus = true
    }

    stringX := strconv.Itoa(x)
    sliceX := strings.Split(stringX, \\)
    invertedIndex := len(sliceX)-1
    
    fmt.Println(stringX)

    for i := 0; i<len(sliceX)/2; i++ {
        temp := sliceX[i]
        sliceX[i] = sliceX[invertedIndex]
        sliceX[invertedIndex] = temp
        invertedIndex--
    }

    stringX = strings.Join(sliceX, \\)
    
    x,_ = strconv.Atoi(stringX)

    if minus {
        x *= -1
    }
    
    if x == 0 || x > int(math.Pow(2,31)-1) || x < int(math.Pow(2,31))*(-1) {
        return 0
    }

    return x
}