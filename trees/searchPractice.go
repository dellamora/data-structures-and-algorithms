// Practice Problema
// Comparing two binary trees to seeif they are equal in both shape and structure

package main

// type node struct {
// 	data int
// 	left *node
// 	right *node
// }

func compare (binary * node, binary2 * node) bool {
	if binary == nil && binary2 == nil {
		return true
	}
	if binary == nil || binary2 == nil {
		return false
	}
	return binary.data == binary2.data && compare(binary.left, binary2.left) && compare(binary.right, binary2.right)
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
// 	binary2 := &node{
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
// 	fmt.Println(compare(binary, binary2))
// }

