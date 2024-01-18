package main

type node struct {
	data  int
	left  *node
	right *node
}

func Search(curr *node, needle int) bool {
	if curr == nil {
		return false
	}
	if curr.data == needle {
		return	true
	}
	if curr.data < needle {
		return Search(curr.right, needle)
	}

	return Search(curr.left, needle)
}

func dfs (head *node, needle int) bool {
	return Search(head, needle)

}


// func main() {
// 	binary := &node{
// 		data: 1,
// 		left: &node{
// 			data: 2,
// 			left: &node{
// 				data: 4,
// 				left: nil,
// 				right: nil,
// 			},
// 			right: &node{
// 				data: 5,
// 				left: nil,
// 				right: nil,
// 			},
// 		},
// 		right: &node{
// 			data: 3,
// 			left: &node{
// 				data: 6,
// 				left: nil,
// 				right: nil,
// 			},
// 			right: &node{
// 				data: 7,
// 				left: nil,
// 				right: nil,
// 			},
// 		},
// 	}
	
// 	fmt.Println(dfs(binary, 1))
// }