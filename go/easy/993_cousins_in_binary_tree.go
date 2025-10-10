/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    xDepth := -1
    yDepth := -1
    
    xParent := -1
    yParent := -1

    var dfs func(root *TreeNode, parent, depth int)
    dfs = func(root *TreeNode, parent, depth int) {
        if root == nil {
            return
        }

        if root.Val == x {
            xDepth = depth
            xParent = parent
        }
        
        if root.Val == y {
            yDepth = depth
            yParent = parent
        }

        dfs(root.Left, root.Val, depth + 1)
        dfs(root.Right, root.Val, depth + 1)
    }

    dfs(root, -1 , 0)

    return (xDepth == yDepth) && (xParent != yParent)
}
