/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
    if root == nil {
        return 0
    }

    sumGrandChildren := 0
    if root.Val % 2 == 0 {
        // get sum of grand children
        if root.Left != nil {
            if root.Left.Left != nil {
                sumGrandChildren += root.Left.Left.Val
            }
            if root.Left.Right != nil {
                sumGrandChildren += root.Left.Right.Val
            }
        }
        if root.Right != nil {
            if root.Right.Left != nil {
                sumGrandChildren += root.Right.Left.Val
            }
            if root.Right.Right != nil {
                sumGrandChildren += root.Right.Right.Val
            }
        }
    }

    l := sumEvenGrandparent(root.Left)
    r := sumEvenGrandparent(root.Right)
    return sumGrandChildren + l + r
}
