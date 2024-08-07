func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }

    ones := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
    teen := []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
    tens := []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
    htmb := []string{"Hundred", "Thousand", "Million", "Billion"}

    stack := []string{}
    position := 0
    subNum := 0

    for x := num; x > 0; x /= 10 {
        position++
        switch position % 3 {
        case 1: 
            subNum = x % 10
            if x < 10 {
                stack = append(stack, ones[subNum])
            }
        case 2: 
            subNum += x % 10 * 10
            switch {
            case subNum > 0 && subNum < 10: // 1 .. 9
                stack = append(stack, ones[subNum])
            case subNum > 9 && subNum < 20: // 10 - 19
                stack = append(stack, teen[subNum - 10])
            case subNum > 19 && subNum < 100: // 20 - 99
                if subNum % 10 != 0 {
                    stack = append(stack, ones[subNum % 10])
                }
                stack = append(stack, tens[subNum / 10 - 2])
            }
        case 0:
            subNum = 0
            if x % 10 > 0 {
                stack = append(stack, htmb[0])
                stack = append(stack, ones[x % 10])
            }
            if x > 9 && (x / 10) % 1000 != 0 {
                stack = append(stack, htmb[position/3])
            }
        }
    }

    var res string
    for i, v := range stack {
        res = v + res
        if i < len(stack) - 1 {
            res = " " + res
        }
    }
    return res
}
