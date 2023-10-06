func nextGreatestLetter(letters []byte, target byte) byte {
    l := 0
    r := len(letters) - 1
    if letters[r] <= target  {
        return letters[0]
    }
    for l < r {
        m := l + (r - l)/2
        pick := letters[m]
        if pick > target {
            r = m
        } else {
            l = m + 1
        }
    }
    return letters[l]
}
