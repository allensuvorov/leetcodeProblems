/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return root
    }
    // search
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else { // delete
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        }
        cur := root.Right
        for cur.Left != nil {
            cur = cur.Left
        }
        root.Val = cur.Val
        root.Right = deleteNode(root.Right, root.Val) // recursive delete
    }
    return root
}
