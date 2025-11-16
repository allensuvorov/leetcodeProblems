type TwoSum struct {
    numsCount map[int]int
}


func Constructor() TwoSum {
    return TwoSum{make(map[int]int)}
}


func (this *TwoSum) Add(number int)  {
    this.numsCount[number]++
}


func (this *TwoSum) Find(value int) bool {
    for num1, num1Count := range this.numsCount {
        num2 := value - num1
        if num1 != num2 {
            if _, exists := this.numsCount[num2]; exists { // main case
                return true
            }
        } else { // special case
            if num1Count >= 2 {
                return true
            }
        } 
    }
    return false
}


/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
