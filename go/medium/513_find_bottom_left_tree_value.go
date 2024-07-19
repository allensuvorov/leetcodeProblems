/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    q := []*TreeNode{root}
    ans := 0
    for len(q) > 0 {
        ans = q[0].Val
        rowLen := len(q)
        for i := range rowLen {
            now := q[i]
            if now.Left != nil {
                q = append(q, now.Left)
            }
            if now.Right != nil {
                q = append(q, now.Right)
            }
        }
        q = q[rowLen:]
    }
    return ans
}
