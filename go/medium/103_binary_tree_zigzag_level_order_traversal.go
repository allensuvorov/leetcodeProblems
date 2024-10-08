/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    ans := [][]int{}
    zig := true
    q := []*TreeNode{root}
    for len(q) > 0 {
        levelSize := len(q)
        level := make([]int, levelSize)
        for i := range levelSize {
            now := q[0]
            q = q[1:]
            if zig {
                level[i] = now.Val
            } else {
                level[levelSize - 1 - i] = now.Val
            }
            if now.Left != nil {
                q = append(q, now.Left)
            }
            if now.Right != nil {
                q = append(q, now.Right)
            }
        }
        ans = append(ans, level)
        zig = !zig
    }
    return ans
}
