func myAtoi(s string) int {
minus := false
    newString := \\
    var num int

    for i := 0; i<len(s); i++ {
        singleString, err := strconv.Atoi(string(s[i]))
        if string(s[i]) == \ \ {
            if len(newString) >= 1 {
                break
            }
            continue
        } else if string(s[i]) == \-\ || string(s[i]) == \+\ {
            if len(newString) >= 1 {
                break
            } else if string(s[i]) == \-\ {
                newString += \0\
                minus = true
            } else {
                newString += \0\
            }
        } else if err != nil {
            fmt.Println(\Terjadi Error\)
            if len(newString) >= 1 {
                break
            } else {
                return 0
            }
        } else {
            newString += strconv.Itoa(singleString)
            fmt.Println(newString)
        }
    }
    num,_ = strconv.Atoi(newString)

    if minus {
        num *= -1
        if num < int(math.Pow(2,31)) * -1 {
            return int(math.Pow(2,31)) * -1
        }
        return num
    }

    if num > int(math.Pow(2,31)) - 1 {
        return int(math.Pow(2,31)) - 1
    }

    return num
}