package main

 type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

 func inorderTraversal(root *TreeNode) []int {
    result := []int{} 
    inorderHelper(root, &result)
    return result
}

func inorderHelper(node *TreeNode, result *[]int) {
    if node != nil {
        inorderHelper(node.Left, result)
        
        *result = append(*result, node.Val)
        
        inorderHelper(node.Right, result)
    }
}
