package collections

import (
	"fmt"
)

/*
	linkedList

	Search O(n)
	Insert O(1)
	Delete O(1)
*/

type node struct {
	value int
	pLeft *node
	pRight *node 
}

func NewNode(v int) *node{
	return &node{
		value:  v,
		pLeft: nil,
		pRight: nil,
	}
}

func (n *node) Add(v int){
	node := NewNode(v)

	lastNode := n.getLastRight()
	
	lastNode.pRight = node
	node.pLeft = lastNode
}

func (n *node) Del(v int) bool {
	
	node := n.Search(v)
	if node == nil {
		return false
	}

	leftNode := node.pLeft
	rightNode := node.pRight

	node = nil
	leftNode.pRight = rightNode
	rightNode.pLeft = leftNode
	return true
}

func (n *node) Search(v int) *node{

	crtNode := n
	for true {
		if crtNode.value == v {
			return crtNode
		}

		if crtNode.pRight == nil{
			break
		}else{
			crtNode = crtNode.pRight
		}
	}

	return nil
}

func (n *node) getLastRight() *node {
	
	currentNode := n
	for true {
		if currentNode.pRight == nil {
			return currentNode
		}else{
			currentNode = currentNode.pRight
		}
	}

	fmt.Println(currentNode)

	return currentNode
}

func (n *node) ToString() {
	
	crtNode := n
	for true {
		fmt.Print(crtNode.value)

		if crtNode.pRight == nil{
			break
		}else{
			crtNode = crtNode.pRight
		}
	}
	fmt.Println()
}



