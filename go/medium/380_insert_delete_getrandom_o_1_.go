type RandomizedSet struct {
    m map[int]int
    a []int
}


func Constructor() RandomizedSet {
    return RandomizedSet {
        m: make(map[int]int), 
        a: make([]int, 0),
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.m[val]; exists {
        return false
    }
    this.a = append(this.a, val)
    this.m[val] = len(this.a) - 1
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if index, exists := this.m[val]; !exists {
        return false
    } else {
        last := len(this.a) - 1
        this.a[index] = this.a[last]
        this.m[this.a[index]] = index
        delete(this.m, val)
        this.a = this.a[:last]
    }
    return true

}


func (this *RandomizedSet) GetRandom() int {
    n := len(this.a)
    randIndex := rand.Intn(n)
    return this.a[randIndex]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
