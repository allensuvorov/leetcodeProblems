/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    minDiff := 100_001
    prev := -100_001

    var lnr func (node *TreeNode) // LNR DFS
    lnr = func (node *TreeNode) {
        if node != nil {
            lnr(node.Left)
            diff := node.Val - prev
            minDiff = min(minDiff, diff)
            prev = node.Val
            lnr(node.Right)
        }
    }
    
    lnr(root)
    return minDiff
}
