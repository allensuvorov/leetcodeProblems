type MyHashSet struct {
    a [][]int
    setSize int
}


func Constructor() MyHashSet {
    setSize := 10
    a := make([][]int, setSize)
    return MyHashSet {a, setSize}
}


func (this *MyHashSet) Add(key int)  {
    group := key % this.setSize
    keySlice:= this.a[group]
    for _, v := range keySlice {
        if v == key {
            return
        }
    }
    this.a[key % this.setSize] = append(this.a[key % this.setSize], key)
}


func (this *MyHashSet) Remove(key int)  {
    group := key % this.setSize
    keySlice:= this.a[group]
    for i, v := range keySlice {
        if v == key {
            keySlice[i] = keySlice[len(keySlice)-1]
            this.a[group] = keySlice[:len(keySlice)-1]
            return
        }
    }
}


func (this *MyHashSet) Contains(key int) bool {
    for _, v := range this.a[key % this.setSize] {
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
