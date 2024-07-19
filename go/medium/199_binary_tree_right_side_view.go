/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    q := []*TreeNode{root}
    ans := []int{}
    for len(q) > 0 {
        rowLen := len(q)
        ans = append(ans, q[0].Val)
        for i := range rowLen {
            now := q[i]
            if now.Right != nil {
                q = append(q, now.Right)
            }
            if now.Left != nil {
                q = append(q, now.Left)
            }
        }
        q = q[rowLen:]
    }
    return ans
}
