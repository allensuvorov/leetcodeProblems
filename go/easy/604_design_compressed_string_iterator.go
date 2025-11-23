type StringIterator struct {
    s string
    pos int
    letter byte
    remNum int
}


func Constructor(compressedString string) StringIterator {
    return StringIterator{
        s: compressedString,
        pos: 0,
        letter: 0,
        remNum: 0,
    }
    
}


func (this *StringIterator) Next() byte {
    if !this.HasNext() {
        fmt.Println("End of line")
        return ' '
    }

    // parse next
    if this.remNum == 0 { 
        // parse the letter, pos is at the new letter
        this.letter = this.s[this.pos]
        // parse number
        this.pos++
        var number int
        for this.pos < len(this.s) && unicode.IsDigit(rune(this.s[this.pos])) {
            number = number * 10
            number = number + int(this.s[this.pos] - '0')
            this.pos++
        }
        this.remNum = number
    }

    this.remNum--
    return this.letter
}


func (this *StringIterator) HasNext() bool {
    return !(this.pos == len(this.s) && this.remNum == 0)
}


/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
    // c num
 // L 123  e2t1C1o1d1e123
    // ^                    ^  
