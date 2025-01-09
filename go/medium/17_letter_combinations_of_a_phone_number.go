func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }
    m := map[byte]string{'2':"abc", '3':"def", '4':"ghi", '5':"jkl", '6':"mno", '7':"pqrs", '8':"tuv", '9':"wxyz"}
    res := []string{}
    for _, c := range m[digits[0]] {
        strs := []string{""}
        if len(digits) > 1 {
            strs = letterCombinations(digits[1:])
        }
        for i, str := range strs {
            strs[i] = string(c) + str
        }
        res = append(res, strs...)
    }
    return res
}

// redo
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    numLetters := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }
    res := []string{}
    var dfs func(digits, letters string)
    dfs = func(digits, letters string){
        if len(digits) == 0 {
            res = append(res, letters)
            return
        }
        for _, letter := range numLetters[digits[0]] {
            dfs(digits[1:], letters + string(letter))
        }
    }
    dfs(digits,"")
    return res
}
