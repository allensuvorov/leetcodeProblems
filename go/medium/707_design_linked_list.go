type MyLinkedList struct {
    head *listNode
    tail *listNode
    len int
}

type listNode struct {
    val int
    next *listNode
}


func Constructor() MyLinkedList {
    return MyLinkedList{}
}


func (this *MyLinkedList) Get(index int) int {
    node := this.getNode(index)
    if node != nil {
        return node.val
    }
    return -1
}


func (this *MyLinkedList) getNode(index int) *listNode {
    i := 0
    for runner := this.head; runner != nil; runner = runner.next {
        if i == index {
            return runner
        }
        i++
    }
    return nil
}

func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := &listNode{val:val}
    if this.len == 0 {
        this.head = newNode
        this.tail = newNode
    } else {
        newNode.next = this.head
        this.head = newNode
    }
    this.len++
}


func (this *MyLinkedList) AddAtTail(val int)  {
    newNode := &listNode{val:val}
    if this.len == 0 {
        this.tail = newNode
        this.head = newNode
    } else {
        this.tail.next = newNode
        this.tail = newNode
    }
    this.len++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    newNode := &listNode{val:val}
    if index == this.len {
        this.AddAtTail(val)
    } else if index == 0 {
        this.AddAtHead(val)
    } else if index < this.len {
        prev := this.getNode(index - 1)
        next := prev.next
        prev.next = newNode
        newNode.next = next
        this.len++
    }
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if this.len == 0 || index >= this.len  {
        return
    }

    if index == 0 { // delete head
        this.head = this.head.next
    } else if index == this.len - 1 { // delete tail
        prev := this.getNode(index - 1)
        this.tail = prev
        prev.next = nil
    } else {
        prev := this.getNode(index - 1)
        prev.next = prev.next.next

    }
    this.len-- 
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
