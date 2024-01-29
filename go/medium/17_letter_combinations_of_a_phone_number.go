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
