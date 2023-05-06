func compress(chars []byte) int {
    // just return the chars
    lastSym := chars[0]
    lastPos := 0
    insertPos := 0
    
    for i := range chars {
        if chars[i] != lastSym{
            insertPos = pack(chars, lastSym, i - lastPos, insertPos)
            lastSym = chars[i]            
            lastPos = i
        } 
    }
    insertPos = pack(chars, lastSym, len(chars) - lastPos, insertPos)
    return insertPos
}

func pack(chars []byte, lastSym byte, cnt, insertPos int) int {
    chars[insertPos] = lastSym
    insertPos++
    
    if cnt > 1 {
        if cnt < 10{
            chars[insertPos] = '0' + byte(cnt)
            insertPos++
        } else {
            digits := strconv.Itoa(cnt)
            for i := range digits{
                chars[insertPos] = digits[i]
                insertPos++
            }
        }
    }
    return insertPos
}
