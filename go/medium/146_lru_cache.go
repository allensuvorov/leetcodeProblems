import "container/list"

type LRUCache struct {
    C int
    L *list.List
    M map[int]*list.Element
}

type Node struct {
    Key int
    Value int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        C: capacity,
        L: list.New(),
        M: make(map[int]*list.Element, capacity), 
    }
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.M[key]; ok {
        this.L.MoveToBack(node)
        return node.Value.(*Node).Value
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.M[key]; ok { // exist
        this.L.MoveToBack(node)
        node.Value.(*Node).Value = value
    } else {
        node := &Node{
            Key: key,
            Value: value,
        }
        if this.L.Len() == this.C { // map is at capacity
            lru := this.L.Front() // get LRU from back of the list
            lruKey := lru.Value.(*Node).Key
            delete(this.M, lruKey)
            this.L.Remove(lru)
        }
        this.M[key] = this.L.PushBack(node)
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
 
