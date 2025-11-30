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
        size := len(q)
        res = append(res, q[size - 1].Val)
        for range size {
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
// iterative bfs with q, take value from end of q
