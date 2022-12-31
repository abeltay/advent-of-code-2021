package aoc

import (
	"strings"
	"unicode"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	var nodes []*Node
	for _, v := range data {
		sarr := strings.Split(v, "-")
		var startNode *Node
		for i := range nodes {
			if nodes[i].name == sarr[0] {
				startNode = nodes[i]
				break
			}
		}
		if startNode == nil {
			startNode = &Node{name: sarr[0]}
			nodes = append(nodes, startNode)
		}
		var endNode *Node
		for i := range nodes {
			if nodes[i].name == sarr[1] {
				endNode = nodes[i]
				break
			}
		}
		if endNode == nil {
			endNode = &Node{name: sarr[1]}
			nodes = append(nodes, endNode)
		}
		startNode.items = append(startNode.items, endNode)
		endNode.items = append(endNode.items, startNode)
	}
	/*
		for i := range nodes {
			fmt.Printf("%s\n", nodes[i].name)
		}
	*/
	var start *Node
	for i := range nodes {
		if nodes[i].name == "start" {
			start = nodes[i]
			break
		}
	}
	count := Delve([]*Node{start})
	/*
		for i := range start.items {
			fmt.Printf("%s\n", start.items[i].name)
		}
	*/
	return count
}

// Node represents a node
type Node struct {
	name  string
	items []*Node
}

// SizeBig checks if a cave is big
func (f Node) SizeBig() bool {
	return unicode.IsUpper(rune(f.name[0]))
}

// CloneNodes creates a clone
func CloneNodes(nodes []*Node) []*Node {
	newNodes := make([]*Node, len(nodes))
	for i := range nodes {
		newNodes[i] = nodes[i]
	}
	return newNodes
}

// Delve travels through the graph
func Delve(path []*Node) int {
	cur := path[len(path)-1]
	if cur.name == "end" {
		return 1
	}
	var count int
	for i := range cur.items {
		if !cur.items[i].SizeBig() {
			if cur.items[i].name == "start" {
				continue
			}
			m := make(map[string]int)
			for j := range path {
				if !path[j].SizeBig() {
					m[path[j].name]++
				}
			}
			if m[cur.items[i].name] >= 2 {
				continue
			}
			var illegal bool
			if m[cur.items[i].name] == 1 {
				for _, v := range m {
					if v == 2 {
						// fmt.Printf("%s, %v\n", cur.items[i].name, m)
						illegal = true
						break
					}
				}
			}
			if illegal {
				continue
			}
		}
		newPath := CloneNodes(path)
		newPath = append(newPath, cur.items[i])
		count += Delve(newPath)
	}
	return count
}
