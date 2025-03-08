type LockingTree struct {
    parent []int // the input; need for checking ancestors
    locker []int // user who locked node i, 0 if the node is unlocked
    children [][]int // children for each node, need for checking descendants
}


func Constructor(parent []int) LockingTree {
    n := len(parent)
    children := make([][]int, n)

    for i := 1; i < n; i++ {
        children[parent[i]] = append(children[parent[i]], i)
    }

    locker := make([]int, n)
    return LockingTree{parent, locker, children}
}


func (this *LockingTree) Lock(num int, user int) bool {
    if this.locker[num] == 0 {
        this.locker[num] = user
        return true
    }
    return false
}


func (this *LockingTree) Unlock(num int, user int) bool {
    if this.locker[num] == user {
        this.locker[num] = 0
        return true
    }
    return false
}


func (this *LockingTree) Upgrade(num int, user int) bool {
    if this.locker[num] != 0 {
        return false 
    }
    // Traverse parents as a linked list
    // to check if there are any locked ancestors
    for node := num; this.parent[node] != -1; node = this.parent[node] {
        if this.locker[this.parent[node]] != 0 {
            return false
        }
    }
    
    // DFS over descendants
    // to check there is at least one locked descendant (by any user)
    // unlock them along the way
    stack := []int{num}
    lockedDescendantExists := false

    for len(stack) > 0 {
        top := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]

        if this.locker[top] > 0 {
            lockedDescendantExists = true
            this.locker[top] = 0
        }

        for _, child := range this.children[top] {
            stack = append(stack, child)
        }
    }

    if !lockedDescendantExists {
        return false
    }

    this.locker[num] = user
    return true
}


/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */
