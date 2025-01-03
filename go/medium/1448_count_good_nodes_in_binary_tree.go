/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    return dfs(root, root.Val)
}

func dfs(node *TreeNode, max int) int {
    if node == nil {
        return 0
    }
    good := 0
    if max <= node.Val {
        good = 1
        max = node.Val
    }
    return good + dfs(node.Left, max) + dfs(node.Right, max)
}

// redo 3 Jan 2025
func goodNodes(root *TreeNode) int {
    res := 0
    var dfs func(root *TreeNode, greatest int)
    dfs = func(root *TreeNode, greatest int) {
        if root == nil {
            return
        }
        if root.Val >= greatest {
            res++
            greatest = root.Val
        }
        
        dfs(root.Left, greatest)
        dfs(root.Right, greatest)
    }

    dfs(root, -1e4)
    return res
}
