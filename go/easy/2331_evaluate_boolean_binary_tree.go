/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
    if root == nil {
        return false
    }
    numToBool := map[int]bool{0:false, 1:true}
    
    // leaf - convert to bool
    if root.Left == nil {
        return numToBool[root.Val]
    }

    // non leaf - do operation
    numToOperation := map[int] func(a,b bool) bool {
        2: func(a,b bool) bool {
            return a || b
        },
        3: func(a,b bool) bool {
            return a && b
        },
    }

    left := evaluateTree(root.Left)
    right := evaluateTree(root.Right)
    return numToOperation[root.Val](left, right)
}
