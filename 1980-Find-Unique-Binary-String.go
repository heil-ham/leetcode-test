import (
    "fmt"
    "math"
    "strconv"
    )

func convertBinaryToInt(bin string) int {
    result := 0
    for index,binary := range bin {
        convertedBinaryToInt,_ := strconv.Atoi(string(binary))
        result += convertedBinaryToInt * int(math.Pow(float64(2), float64(len(bin)-1-index)))
    }
    return result
}

func convertIntToBinary(num int,digit int) string{
    binary := ""
    invertedBinary := ""
    zeroes := ""
    
    for i := 0; i < digit; i++ {
            zeroes += "0"
    }
    
    if num > 1 {
        for num != 1{
            binary += strconv.Itoa(num % 2)
            num = int(num / 2)
            if num == 1 {
                binary += strconv.Itoa(num)
            }
        }
        
        for i := len(binary)-1; i >=0; i-- {
            invertedBinary += string(binary[i])
        }
        
        trimmedZeroes := zeroes[0:digit-len(invertedBinary)]
        trimmedZeroes += invertedBinary
        
        return trimmedZeroes
        
        
    } else {
        if num == 1 {
            trimmedZeroes := zeroes[0:digit-1]
            trimmedZeroes += "1"
            return trimmedZeroes
        } else {
            return zeroes
        }
    }
}

func findDifferentBinaryString(nums []string) string {
    numOfBinary := int(math.Pow(2,float64(len(nums))))

    for i := 0; i < numOfBinary; i++ {
        binaryFound := false
        binary := convertIntToBinary(i, len(nums))
        for j := 0; j < len(nums); j++ {
            if binary == nums[j] {
                binaryFound = true
                break
            }
        }
        if !binaryFound {
            return binary
        }
    }
    
    return "All Found !"

}