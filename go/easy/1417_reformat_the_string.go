func reformat(s string) string {
    
    if len(s) == 1 {
        return s
    }
    
    n, ok := isPossible(s)
    
    if !ok {
        return ""
    }
    
    numbers := make([]rune, 0, n)
    letters := make([]rune, 0, len(s) - n)

    for _, v := range s {
        if unicode.IsNumber(v) {
            numbers = append(numbers, v)
        }
        if unicode.IsLetter(v) {
            letters = append(letters, v)
        }
    }

    res := make([]rune, 0, len(s))

    switch {
    case n > len(s) - n: 
        res = fillSlice(numbers, letters)
    case n == len(s) - n:
        res = fillSlice(numbers, letters)
    case n < len(s) - n:
        res = fillSlice(letters, numbers)
    }
    return string(res)
}

func fillSlice(long, short []rune) []rune {
    res := make([]rune, 0, len(long)+len(short))
    for i := range short {
        res = append(res, long[i])
        res = append(res, short[i])
    }
    if len(long)> len(short) {
        res = append(res, long[len(long)-1])
    }
    return res
}

func isPossible(s string) (int, bool) {
    n := 0
    for _, v := range s {
        if unicode.IsNumber(v) {
            n++
        }
    }
    if len(s) % 2 == 0 {
        if len(s) / 2 != n {
            return 0, false
        }
    } else {
        if len(s) / 2 != n && len(s) / 2 + 1 != n {
            return 0, false
        }
    }
    return n, true
}
