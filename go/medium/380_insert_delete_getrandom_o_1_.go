type RandomizedSet struct {
    nums []int
    set map[int]int
}


func Constructor() RandomizedSet {
    return RandomizedSet{[]int{}, map[int]int{}}
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.set[val]; ok {
        return false
    }

    this.nums = append(this.nums, val)
    this.set[val] = len(this.nums) - 1

    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.set[val]; !ok {
        return false
    }

    i := this.set[val]
    end := len(this.nums) - 1
    this.nums[i] = this.nums[end] // swap val from end
    this.set[this.nums[end]] = i // update index in set
    this.nums = this.nums[:end] // pop
    delete(this.set, val)
    return true
}


func (this *RandomizedSet) GetRandom() int {
    return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

 // 2, 2 
