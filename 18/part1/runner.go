package aoc

import (
	"fmt"
	"strconv"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	arr := make([]*binaryNode, len(data))
	for i := range data {
		sf := createBinaryTree(data[i])
		arr[i] = sf
	}
	root := arr[0]
	var notFirst bool
	for i := range arr {
		if notFirst {
			root = &binaryNode{
				left:  root,
				right: arr[i],
			}
		}
		notFirst = true
		explodeAndSplitBinaryTree(root)
		// depthFirstBinaryPrint(root)
		// fmt.Println()
	}
	ans := calcMagnitudeBinaryTree(root)
	return ans
}

type binaryNode struct {
	left  *binaryNode
	right *binaryNode
	data  int
}

func createBinaryTree(s string) *binaryNode {
	var open int
	var pointer int
	// Split to 2 tokens first
	for pointer = 1; pointer < len(s); pointer++ {
		switch s[pointer] {
		case '[':
			open++
		case ']':
			open--
		}
		if open == 0 && s[pointer] == ',' {
			break
		}
	}
	token1 := s[1:pointer]
	// fmt.Println(token1)
	pointer++
	token2 := s[pointer : len(s)-1]
	// fmt.Println(token2)

	var cur binaryNode
	if token1[0] == '[' {
		cur.left = createBinaryTree(token1)
	} else {
		num, err := strconv.Atoi(token1)
		if err != nil {
			return nil
		}
		cur.left = &binaryNode{data: num}
	}
	if token2[0] == '[' {
		cur.right = createBinaryTree(token2)
	} else {
		num, err := strconv.Atoi(token2)
		if err != nil {
			return nil
		}
		cur.right = &binaryNode{data: num}
	}
	return &cur
}

func explodeAndSplitBinaryTree(root *binaryNode) {
	for {
		prev, tuple, next := explodeBinaryTree(root, nil, 1)
		if tuple != nil {
			if prev != nil {
				prev.data += tuple.left.data
			}
			if next != nil {
				if isLeaf(next) {
					next.data += tuple.right.data
				} else {
					addLeftMost(next, tuple.right.data)
				}
			}
			continue
		}
		splitted := splitBinaryTree(root)
		if splitted {
			continue
		}
		break
	}
}

func explodeBinaryTree(s, prev *binaryNode, depth int) (*binaryNode, *binaryNode, *binaryNode) {
	if isLeaf(s) {
		return s, nil, nil
	}
	if depth > 4 {
		if isLeaf(s.left) && isLeaf(s.right) {
			cur := &binaryNode{
				left:  s.left,
				right: s.right,
			}
			s.left = nil
			s.right = nil
			s.data = 0
			return prev, cur, nil
		}
	}
	prev, tuple, next := explodeBinaryTree(s.left, prev, depth+1)
	if next != nil {
		return prev, tuple, next
	}
	if tuple != nil {
		return prev, tuple, s.right
	}
	prev, tuple, next = explodeBinaryTree(s.right, prev, depth+1)
	return prev, tuple, next
}

func isLeaf(s *binaryNode) bool {
	return s.left == nil && s.right == nil
}

func splitBinaryTree(s *binaryNode) bool {
	if s == nil {
		return false
	}
	if isLeaf(s) {
		if s.data > 9 {
			half := s.data / 2
			s.left = &binaryNode{data: half}
			s.right = &binaryNode{data: s.data - half}
			s.data = 0
			return true
		}
	}
	splitted := splitBinaryTree(s.left)
	if splitted {
		return true
	}
	return splitBinaryTree(s.right)
}

func addLeftMost(s *binaryNode, num int) {
	if s.left == nil {
		s.data += num
		return
	}
	addLeftMost(s.left, num)
}

func depthFirstBinaryPrint(s *binaryNode) {
	if s.left == nil && s.right == nil {
		fmt.Print(s.data)
		return
	}
	fmt.Print("[")
	depthFirstBinaryPrint(s.left)
	fmt.Print(",")
	depthFirstBinaryPrint(s.right)
	fmt.Print("]")
}

func calcMagnitudeBinaryTree(s *binaryNode) int {
	var left, right int
	if isLeaf(s.left) {
		left = s.left.data
	} else {
		left = calcMagnitudeBinaryTree(s.left)
	}
	if isLeaf(s.right) {
		right = s.right.data
	} else {
		right = calcMagnitudeBinaryTree(s.right)
	}
	return 3*left + 2*right
}
