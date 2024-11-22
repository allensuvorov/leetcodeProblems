/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    stack := make([]*TreeNode, 0)
    for root != nil {
        stack = append(stack, root)
        root = root.Left
    }
    return BSTIterator{stack: stack}
}


func (this *BSTIterator) Next() int {
    top := len(this.stack) - 1
    result := this.stack[top]
    this.stack = this.stack[:top]
    for node := result.Right; node != nil; node = node.Left {
        this.stack = append(this.stack, node)
    }
    return result.Val
}


func (this *BSTIterator) HasNext() bool {
    return len(this.stack) != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
