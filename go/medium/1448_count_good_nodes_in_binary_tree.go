/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    var dfs func(root *TreeNode, greatest int) int
    dfs = func(root *TreeNode, greatest int) int {
        if root == nil {
            return 0
        }
        good := 0
        if root.Val >= greatest {
            good = 1
            greatest = root.Val
        }
        return good + dfs(root.Left, greatest) + dfs(root.Right, greatest)
    }
    return dfs(root, root.Val)
}

// closure
func goodNodes(root *TreeNode) int {
    result := 0

    var dfs func(root *TreeNode, maxSoFar int)
    dfs = func(root *TreeNode, maxSoFar int) {
        if root == nil {
            return
        }
        if root.Val >= maxSoFar {
            result++
            maxSoFar = root.Val
        }
        dfs(root.Left, maxSoFar)
        dfs(root.Right, maxSoFar)
    }
    dfs(root, math.MinInt)
    return result
}

