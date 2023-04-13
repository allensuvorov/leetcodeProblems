package easy

func numDifferentIntegers(word string) int {
    nums := make(map[string]struct{})
    var num strings.Builder 
    inNum := false
    for i, v := range word {
        if inNum {
            if unicode.IsDigit(v) {         // inside num
                num.WriteRune(v)
            } else {                        // end num
                fmt.Println(num.String())

                nums[num.String()]=struct{}{}
                num.Reset()
                inNum = false
            }   
        } else {                            // start num
            if unicode.IsDigit(v) {
                if (v -'0') == 0 {
                    // if zero not leading - write
                    if i+1 == len(word) || !unicode.IsDigit(rune(word[i+1])) {
                        num.WriteRune(v)
                        inNum = true
                    }
                } else {
                    num.WriteRune(v)
                    inNum = true
                }
            }
        }
    }
    if unicode.IsDigit(rune(word[len(word)-1])) {
        nums[num.String()]=struct{}{}
    }

    fmt.Println(nums)
    return len(nums)
}
