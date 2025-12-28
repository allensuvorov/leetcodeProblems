/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaves1 := []int{}
    dfs(root1, &leaves1)
    leaves2 := []int{}
    dfs(root2, &leaves2)
    if len(leaves1) != len(leaves2) {
        return false
    }
    for i := range leaves1 {
        if leaves1[i] != leaves2[i] {
            return false
        }
    }
    return true
}

func dfs(root *TreeNode, leaves *[]int){
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        *leaves = append(*leaves, root.Val)
    }
    dfs(root.Left, leaves)
    dfs(root.Right, leaves)
}


/// channels solution

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    c1 := make(chan int)
    c2 := make(chan int)
    go func() {
        findLeaves(root1, c1)
        close(c1)
    }()
    go func() {
        findLeaves(root2, c2)
        close(c2)
    }()

    for {
        v1, ok1 := <- c1
        v2, ok2 := <- c2
        if !ok1 && !ok2 {
            return true
        }
        if ok1 != ok2 || v1 != v2 {
            return false
        }
    }
    return true
}


func findLeaves(root *TreeNode, c chan int) {
    if root != nil {
        if root.Left == nil && root.Right == nil {
            c <- root.Val
        } else {
            findLeaves(root.Left, c)
            findLeaves(root.Right, c)
        }
    }
}
