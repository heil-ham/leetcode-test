func romanToInt(s string) int {
    num := 0

    for i:=0; i<len(s); i++ {
        if s[i] == 'M' {
            if i != 0 {
                if s[i-1] == 'C' {
                    num -= 200
                }
            } 
            num += 1000
        } else if s[i] == 'D' {
            if i != 0 {
                if s[i-1] == 'C' {
                    num -= 200
                }
            } 
            num += 500
        } else if s[i] == 'C' {
            if i != 0 {
                if s[i-1] == 'X' {
                    num -= 20
                }
            } 
            num += 100
        } else if s[i] == 'L' {
            if i != 0 {
                if s[i-1] == 'X' {
                    num -= 20
                }
            } 
            num += 50
        } else if s[i] == 'X' {
            if i != 0 {
                if s[i-1] == 'I' {
                    num -= 2
                }
            } 
            num += 10
        } else if s[i] == 'V' {
            if i != 0 {
                if s[i-1] == 'I' {
                    num -= 2
                }
            } 
            num += 5
        } else if s[i] == 'I' {
            num += 1
        }
    }

    return num
}