/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    numCPU := runtime.NumCPU()
    fmt.Printf("numCPU: %v", numCPU )
    sem := make(chan int, numCPU)
    return concurrentDFS(root, root.Val, sem)
}

func concurrentDFS(root *TreeNode, good int, sem chan int) int {
    if root == nil {
        return 0
    }
    res := 0
    if root.Val >= good {
        res += 1
        good = root.Val
    }
    var wg sync.WaitGroup
    var l, r int
    
    select {
    case sem <- 1:
        wg.Add(1)
        go func(){
            l = concurrentDFS(root.Left, good, sem)
            wg.Done()
        }()
    default:
        l = dfs(root.Left, good)
    }

    select {
    case sem <- 1:
        wg.Add(1)
            go func(){
        r = concurrentDFS(root.Right, good, sem)
        wg.Done()
        }()
    default:
        r = dfs(root.Right, good)
    }

    wg.Wait()
    return res + l + r
}

func dfs(root *TreeNode, good int) int {
    if root == nil {
        return 0
    }
    res := 0
    if root.Val >= good {
        res += 1
        good = root.Val
    }
    
    l := dfs(root.Left, good)
    r := dfs(root.Right, good)

    return res + l + r
}
