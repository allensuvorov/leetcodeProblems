/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    // descendents - target sub-tree - DFS from target
    // ancestors 
    //  - bfs from root to learn level of target
    //  - go remaining distance in opposite subtree

    // but let's try and create a more universal approach 
    // may be it's in direction of treversal
    res := []int{}
    var dfs func(node *TreeNode, dist int) int
    dfs = func(node *TreeNode, dist int) int {
        if root == nil {
            return 0
        }

        if dist == k {
            res = append(res, node.Val)
        }

        if root == target || dist > 0{
            dist =+
        }

        if dfs(node.Left, dist) > 0 {
            dist++
        }
        if dfs(node.Right, dist) > 0 {
            dist++
        }
        
        return dist
    }

    dfs(root, 0)

    return res

}
