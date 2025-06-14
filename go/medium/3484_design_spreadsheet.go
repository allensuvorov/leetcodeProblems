type Spreadsheet struct {
    data map[string]int
}


func Constructor(rows int) Spreadsheet {
    data := make(map[string]int)
    return Spreadsheet{data: data}
}


func (this *Spreadsheet) getNum (s string) int {
    if num, err := strconv.Atoi(s); err == nil {
        return num
    }

    return this.data[s]
}


func (this *Spreadsheet) SetCell(cell string, value int)  {
    this.data[cell] = value
}


func (this *Spreadsheet) ResetCell(cell string)  {
    this.data[cell] = 0
}


func (this *Spreadsheet) GetValue(formula string) int {
    terms := strings.Split(formula[1:], "+")
    return this.getNum(terms[0]) + this.getNum(terms[1])
}


/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
