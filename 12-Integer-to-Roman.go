func intToRoman(num int) string {
    roman := \\

    stringNum := strconv.Itoa(num)
    exp := len(stringNum)-1

    for i:=0; i<len(stringNum); i++ {
        if stringNum[i] == '0' {
            exp--
            continue
        }

        roman += getSingleRoman(string(stringNum[i]), exp)
        exp--
    }

    return roman
}

func getSingleRoman(stringNum string, exp int) string {
    num,_ := strconv.Atoi(stringNum)
    singleRoman := \\
    roman := \\

    if exp >= 3 {
        singleRoman = \M\
    } else if exp == 2 {
        if num == 4 {
            return \CD\
        } else if num == 9 {
            return \CM\
        } else if num < 5 {
            singleRoman = \C\
        } else if num >= 5 {
            roman = \D\
            singleRoman = \C\
        }
    } else if exp == 1 {
        if num == 4 {
            return \XL\
        } else if num == 9 {
            return \XC\
        } else if num < 5 {
            singleRoman = \X\
        } else if num >= 5 {
            roman = \L\
            singleRoman = \X\
        }
    } else {
        if num == 4 {
            return \IV\
        } else if num == 9 {
            return \IX\
        } else if num < 5 {
            singleRoman = \I\
        } else if num >= 5 {
            roman = \V\
            singleRoman = \I\
        }
    }

    if num >= 5 {
        num -= 5
    }

    for i:=0; i<num; i++ {
        roman += singleRoman
    }

    return roman
}