/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func binaryTreePaths(root *TreeNode) []string {
    return traverse(root, "")
}

func traverse(node *TreeNode, cur string) []string {
    curRes := []string{}
    if node.Left != nil {
        curRes = append(curRes, traverse(node.Left, fmt.Sprintf("%s%d->", cur, node.Val))...)
    }
    if node.Right != nil {
        curRes = append(curRes, traverse(node.Right, fmt.Sprintf("%s%d->", cur, node.Val))...)
    }
    if  node.Left == nil && node.Right == nil {
        curRes = append(curRes, fmt.Sprintf("%s%d", cur, node.Val))
    }
    return curRes
}