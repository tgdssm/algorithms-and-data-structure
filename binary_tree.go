package main

type TreeNode struct {
	content int
	left    *TreeNode
	right   *TreeNode
}

func (n *TreeNode) insertNode(newNode *TreeNode) {
	if n.content > newNode.content {
		if n.left == nil {
			n.left = newNode
		} else {
			n.right.insertNode(newNode)
		}
	} else {
		if n.right == nil {
			n.right = newNode
		} else {
			n.right.insertNode(newNode)
		}
	}
}

func (n *TreeNode) findNode(content int) *TreeNode {
	if n == nil {
		return nil
	} else if n.content > content {
		return n.left.findNode(content)
	} else if n.content < content {
		return n.right.findNode(content)
	} else {
		return n
	}
}

type Tree struct {
	root *TreeNode
}

func NewBinaryTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(content int) {
	newNode := &TreeNode{content: content}
	if t.root == nil {
		t.root = newNode
	} else {
		t.root.insertNode(newNode)
	}
}

func (t *Tree) Find(content int) *TreeNode {
	return t.root.findNode(content)
}
