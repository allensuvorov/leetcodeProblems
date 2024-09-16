/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    if root == nil {
        return nil
    }
    ans := [][]int{}
    q := []*Node{root}
    for len(q) > 0 {
        levelSize := len(q)
        level := []int{}
        for range levelSize {
            now := q[0]
            q = q[1:]
            level = append(level, now.Val)
            q = append(q, now.Children...)
        }
        ans = append(ans, level)
    }
    return ans
}
