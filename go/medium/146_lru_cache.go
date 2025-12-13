type item struct {
    le *list.Element
    v int
}

type LRUCache struct {
    c int
    q *list.List // DLL, que of keys
    m map[int]*item // map of pointers to items
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        c: capacity,
        q: list.New(),
        m: make(map[int]*item),
    }
}


func (this *LRUCache) Get(key int) int {
    if item, ok := this.m[key]; ok { // if key in que move to back
        this.q.MoveToBack(item.le)
        return item.v
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if obj, ok := this.m[key]; ok { // if element esixts
        this.q.MoveToBack(obj.le) // if key in m, move to back
        obj.v = value // update value
    } else { // else push to back
        le := this.q.PushBack(key)
        obj := &item {
            le: le,
            v: value,
        }
        this.m[key] = obj
    }

    // if over cap, evict, 
    if this.q.Len() > this.c {
        le := this.q.Front()
        this.q.Remove(le)
        delete(this.m, le.Value.(int))
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// q  key , key, key
// using moves item to back of queue, evict from front

// m k:v, k:v, k:v

//  front  <-    back
