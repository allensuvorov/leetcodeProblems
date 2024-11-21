type Element struct {
    Key int
    Val int
    Next *Element
    Prev *Element
}


type LRUCache struct {
    Capacity int
    Elements map[int] *Element
    Front *Element
    Back *Element
}


func Constructor(capacity int) LRUCache {
    return LRUCache{ Capacity: capacity, Elements: make(map[int]*Element) }
}


func (this *LRUCache) MoveToBack(element *Element) {
    if element.Next == nil {
        return
    }

    if element.Prev == nil {
        this.Front = element.Next
        this.Front.Prev = nil
    } else {
        element.Prev.Next, element.Next.Prev = element.Next, element.Prev
    }

    element.Next = nil
    element.Prev = this.Back
    this.Back.Next = element
    this.Back = element
}


func (this *LRUCache) Get(key int) int {
    if element, ok := this.Elements[key]; ok {
        this.MoveToBack(element)
        return element.Val
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if element, ok := this.Elements[key]; ok {
        element.Val = value
        this.MoveToBack(element)
    } else {
        element = &Element{
            Key: key,
            Val: value,
        }

        if len(this.Elements) == 0 {
            this.Back = element
            this.Front = element 
        } else {
            element.Prev = this.Back
            this.Back.Next = element
            this.Back = element
        }

        this.Elements[key] = element
        if len(this.Elements) > this.Capacity {
            delete(this.Elements, this.Front.Key)

            this.Front = this.Front.Next
            this.Front.Prev = nil
        }
    }
}



/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
