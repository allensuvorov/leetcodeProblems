/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    var res, prev *TreeNode


    var lnr func(node *TreeNode)
    lnr = func (node *TreeNode) {
        if node == nil {
            return
        }

        lnr(node.Left)

        if prev == nil {
            res = node
            fmt.Println(res.Val)
        }

        if prev != nil {
            prev.Right = node
        }

        node.Left = nil
        prev = node

        lnr(node.Right)
    }
    lnr(root)
    return res
}
