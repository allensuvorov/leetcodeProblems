type MyHashMap struct {
    HashTable []*Record
    MapSize int
}


type Record struct {
    Key int
    Val int
    Next *Record
}


func Constructor() MyHashMap {
    mapSize := 10
    hashTable := make([]*Record, mapSize)
    return MyHashMap{hashTable, mapSize}
}


func (this *MyHashMap) Put(key int, value int)  {
    bucket := key % this.MapSize
    runner := this.HashTable[bucket]
    keyExists := false
    for runner != nil {
        if runner.Key == key {
            keyExists = true
            runner.Val = value
        }
        runner = runner.Next 
    }
    if !keyExists { 
        this.HashTable[bucket] = &Record{
            Key: key, 
            Val: value,
            Next: this.HashTable[bucket],
            } 
    }
}


func (this *MyHashMap) Get(key int) int {
    bucket := key % this.MapSize
    runner := this.HashTable[bucket]
    for runner != nil {
        if runner.Key == key {
            return runner.Val
        }
        runner = runner.Next 
    }
    return -1
}


func (this *MyHashMap) Remove(key int)  {
    bucket := key % this.MapSize
    runner := this.HashTable[bucket]
    prev := runner
    for runner != nil {
        if runner.Key == key {
            if runner == this.HashTable[bucket] {
                this.HashTable[bucket] = runner.Next
            } else {
                prev.Next = runner.Next
            }
        }
        prev = runner
        runner = runner.Next 
    }
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
