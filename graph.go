package main

import (
	"math"
)

//4.1 Route between nodes: Given a directed graph, design an algorithm to find out whether there is a
// route between two nodes

type Node struct {
	Name    string
	Node    []Node
	IsVisit bool
}

var (
	path []string
)

func routeBetweenNodes(graph Node, a string, b string) []string {
	return dfs(graph, a, b)
}

func dfs(graph Node, a string, b string) []string {
	if graph.Node == nil && graph.Name != b {
		path = path[:len(path)-1]
		return path
	}
	graph.IsVisit = true

	if graph.Name == a {
		path = append(path, graph.Name)
	}

	for i := range graph.Node {
		if !graph.Node[i].IsVisit {
			graph.Node[i].IsVisit = true
			if len(path) > 1 && path[0] == a && path[len(path)-1] == b {
				return path
			}
			path = append(path, graph.Node[i].Name)
			dfs(graph.Node[i], a, b)
		}
	}

	return path
}

//4.2 Minimal tree: Given a sorted (increasing order) array with unique integer elements, write an algorithm
// to create a binary search tree with minimal height

type NodeInt struct {
	Value     int
	LeftNode  *NodeInt
	RightNode *NodeInt
}

func minimalTree(arr []int) *NodeInt {
	arrLength := len(arr)
	mid := int(math.Floor(float64(arrLength) / 2))

	if len(arr) == 1 {
		return &NodeInt{
			Value: arr[mid],
		}
	}
	if mid <= 1 {
		var right *NodeInt
		if len(arr) > 2 {
			right = &NodeInt{Value: arr[mid+1]}
		}
		return &NodeInt{
			Value:     arr[mid],
			LeftNode:  &NodeInt{Value: arr[mid-1]},
			RightNode: right,
		}
	}
	leftNode := minimalTree(arr[0:mid])
	rightNode := minimalTree(arr[mid+1 : arrLength])

	nodeInt := NodeInt{
		Value:     arr[mid],
		LeftNode:  leftNode,
		RightNode: rightNode,
	}

	return &nodeInt
}

// 4.3 List of Depths: given a binary tree, design an algorithm which creates a linked list
// of all the nodes at each depth (e.g if you have a tree with depth D, you'll have D linked lists)

/*func listOfDepths(binaryTree *NodeInt) []LinkList {
	var (
		q            Queue
		lls          []LinkList
		depth        = 1
		currentCount int
	)
	q.Push(binaryTree)

	for !q.IsEmpty() {
		binaryTree := q.Pop().(*NodeInt)
		if binaryTree.LeftNode != nil {
			q.Push(binaryTree.LeftNode)
		}
		if binaryTree.RightNode != nil {
			q.Push(binaryTree.RightNode)
		}

		if len(lls) == 0 {
			rootLinkList := LinkList{
				root: &LinkListNode{
					Value: binaryTree.Value,
				},
				length: 1,
			}
			lls = append(lls, rootLinkList)
			continue
		}

		if int(lls[len(lls)-1].length)+currentCount == depth {
			rootLinkList := LinkList{
				root: &LinkListNode{
					Value: binaryTree.Value,
				},
				length: 1,
			}
			lls = append(lls, rootLinkList)
			currentCount = 0
			// if binaryTree.LeftNode == nil && binaryTree.RightNode == nil {
			// 	depth /= 2
			// } else {
			depth *= 2
			// }
			if binaryTree.LeftNode == nil {
				currentCount += 1
			}
			if binaryTree.RightNode == nil {
				currentCount += 1
			}
			continue
		}

		if binaryTree.LeftNode == nil {
			currentCount += 1
		}
		if binaryTree.RightNode == nil {
			currentCount += 1
		}

		lls[len(lls)-1].AddTail(binaryTree.Value)
	}

	return lls
}*/

func listOfDepths(binaryTree *NodeInt) []LinkList {
	var (
		result       []LinkList
		lastLinklist LinkList
	)

	lastLinklist = LinkList{
		root: &LinkListNode{
			Value: binaryTree,
		},
		length: 1,
	}

	for lastLinklist.root != nil {
		result = append(result, lastLinklist)

		var newList LinkList
		current := lastLinklist.root
		for current != nil {
			node := current.Value.(*NodeInt)
			if node.LeftNode != nil {
				newList.AddTail(node.LeftNode)
			}
			if node.RightNode != nil {
				newList.AddTail(node.RightNode)
			}
			current = current.Next
		}

		lastLinklist = newList
	}

	return result
}
