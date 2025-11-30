/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    res := 0
    var dfs func(root *TreeNode, l, r int)
    dfs = func(root *TreeNode, l, r int) {
        if root == nil {
            return
        }
        res = max(res, l, r)
        dfs(root.Left, r+1, 0)
        dfs(root.Right, 0, l+1)
    }

    dfs(root, 0, 0)
    return res
}

// backtracking
func longestZigZag(root *TreeNode) int {
    res := 0
    var dfs func(root *TreeNode, isRightNode bool) int // return the curr zigzag path len
    dfs = func(root *TreeNode, isRightNode bool) int { 
        if root == nil {
            return 0
        }
        left := dfs(root.Left, false)
        right := dfs(root.Right, true)
        res = max(res, left, right)

        if isRightNode {
            return left + 1
        }
        return right + 1
    }
    dfs(root,true)
    dfs(root,false)
    return res
}


//     0
//   /   \
// 0       0
