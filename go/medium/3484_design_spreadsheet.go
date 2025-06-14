type Spreadsheet struct {
    data [][]int
}


func Constructor(rows int) Spreadsheet {
    data := make([][]int, rows)
    for i := range data {
        data[i] = make([]int, 26)
    }
    return Spreadsheet{data: data}
}

func getRC(s string) (int, int) {
    r, _ := strconv.Atoi(s[1:])
    c := int(s[0]-'A')
    return r - 1, c
}

func (this *Spreadsheet) getNum (s string) int {
    if num, err := strconv.Atoi(s); err == nil {
        return num
    }

    r, c := getRC(s)
    return this.data[r][c]
}

func (this *Spreadsheet) SetCell(cell string, value int)  {
    r, c := getRC(cell)
    this.data[r][c] = value
}


func (this *Spreadsheet) ResetCell(cell string)  {
    r, c := getRC(cell)
    this.data[r][c] = 0
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
