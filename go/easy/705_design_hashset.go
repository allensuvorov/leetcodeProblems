type MyHashSet struct {
    set [][]int
    setSize int
}


func Constructor() MyHashSet {
    setSize := 10
    set := make([][]int, setSize)
    return MyHashSet {set, setSize}
}


func (this *MyHashSet) Add(key int)  {
    group := key % this.setSize
    if !this.Contains(key) {
        this.set[group] = append(this.set[group], key)
    }
}


func (this *MyHashSet) Remove(key int)  {
    group := key % this.setSize
    keySlice := this.set[group]
    lastIndex := len(keySlice)-1
    for i, v := range keySlice {
        if v == key {
            keySlice[i] = keySlice[lastIndex]
            this.set[group] = keySlice[:lastIndex]
            return
        }
    }
}


func (this *MyHashSet) Contains(key int) bool {
    for _, v := range this.set[key % this.setSize] {
        if v == key {
            return true
        }
    }
    return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
