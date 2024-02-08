type MyHashMap struct {
    hashTable []*Record
    mapSize int
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
    
}


func (this *MyHashMap) Get(key int) int {
    
}


func (this *MyHashMap) Remove(key int)  {
    
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
