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
    res := []int{}
    q := []*TreeNode{root}
    for len(q) > 0 {
        res = append(res, q[len(q)-1].Val)
        levelSize := len(q)
        for range levelSize {
            now := q[0]
            q = q[1:]
            if now.Left != nil {
                q = append(q, now.Left)
            }
            if now.Right != nil {
                q = append(q, now.Right)
            }
        }
    }
    return res
}
