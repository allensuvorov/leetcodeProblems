/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil { // root.Val = 6
        return 0
    }
    
    lh, rh := getHeight(root.Left), getHeight(root.Right)

    if lh == rh {
        return 1 << lh + countNodes(root.Right) // left with math
    } else {
        return 1 << rh + countNodes(root.Left) //  right with math
    }

    return 0
}

func getHeight (root *TreeNode) int {
    h := 0
    for root != nil {
        h++
        root = root.Left
    }
    return h
}
