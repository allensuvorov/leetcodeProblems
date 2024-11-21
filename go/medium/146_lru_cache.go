import "container/list"

type Node struct {
    Key int
    Val int
}


type LRUCache struct {
    capacity int
    cache map[int] *list.Element
    list *list.List
}


func Constructor(capacity int) LRUCache {
    return LRUCache { 
        capacity: capacity, 
        cache: make(map[int]*list.Element), 
        list: list.New(),
    }
}


func (this *LRUCache) Get(key int) int {
    if element, ok := this.cache[key]; ok {
        this.list.MoveToBack(element)
        return element.Value.(Node).Val
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    node := Node {Key: key, Val: value}
    if element, ok := this.cache[key]; ok {
        element.Value = node
        this.list.MoveToBack(element)
    } else {
        if this.list.Len() == this.capacity {
            oldest := this.list.Front()
            delete(this.cache, oldest.Value.(Node).Key)
            this.list.Remove(this.list.Front())
        }
        element = this.list.PushBack(node)
        this.cache[key] = element
    }
}



/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
