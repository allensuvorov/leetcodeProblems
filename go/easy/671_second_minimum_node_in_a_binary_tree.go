/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
    if root == nil {
        return -1
    }

    if root.Left == nil { // no children
        return -1
    }

    var l, r int

    if root.Left.Val == root.Val {
        l = findSecondMinimumValue(root.Left)
    } else {
        l = root.Left.Val
    }
    
    if root.Right.Val == root.Val {
        r = findSecondMinimumValue(root.Right)
    } else {
        r = root.Right.Val
    }

    switch {
    case l == -1 && r == -1:
        return -1
    case l != -1 && r != -1:
        return min(l,r)
    case l != -1:
        return l
    case r != -1:
        return r
    }

    return -1
}
