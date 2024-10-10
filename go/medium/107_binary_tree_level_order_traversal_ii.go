/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    // BFS
    res := [][]int{}
    q := []*TreeNode{root}
    for len(q) > 0 {
        level := make([]int, 0, len(q))
        for range cap(level) {
            now := q[0]
            q = q[1:]
            if now.Left != nil {
                q = append(q, now.Left)
            }
            if now.Right != nil {
                q = append(q, now.Right)
            }
            level = append(level, now.Val)
        }
        res = append(res, level)
    }
    
    // reverse
    for l, r := 0, len(res) - 1; l < r; {
        res[l], res[r] = res[r], res[l]
        l++
        r--
    }
    return res
}
