func strongPasswordCheckerII(password string) bool {
    return (
        check8chars(password) &&
        check1Lower(password) &&
        check1Upper(password) &&
        check1Digit(password) &&
        check1Spec(password) &&
        checkNoAdj(password))
}

func check8chars(s string) bool {
    if len(s) >= 8 {
        return true
    } else {
        fmt.Println("Need at least 8 characters")
        return false
    }
}

func check1Lower(s string) bool {
    for _, v := range s {
        if unicode.IsLower(v) {
            return true
        }
    }
    fmt.Println("Must have at least one lowercase letter")
    return false
}

func check1Upper(s string) bool {
    for _, v := range s {
        if unicode.IsUpper(v) {
            return true
        }
    }
    fmt.Println("Must have at least one uppercase letter")
    return false
}

func check1Digit(s string) bool {
    for _, v := range s{
        if unicode.IsDigit(v){
            return true
        }
    }
    fmt.Println("Must have at least one digit letter")
    return false
}

func check1Spec(s string) bool {
    scm := map[byte]bool{}
    sc := "!@#$%^&*()-+"
    for i := range sc {
        scm[sc[i]] = true
    }
    for i := range s {
        if scm[s[i]] {
            return true
        }
    }
    fmt.Println("Must have at least one special letter")
    return false
}

func checkNoAdj(s string) bool {
    for i := 0; i < len(s) - 1; i++{
        if s[i] == s[i+1] {
            fmt.Println("Must not have same in adjacent postion")
            return false
        }
    }
    return true
}
