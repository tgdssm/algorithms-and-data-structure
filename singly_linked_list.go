package main

import "fmt"

type SinglyNode struct {
	content  interface{}
	nextNode *SinglyNode
}

func newSinglyNode(content interface{}) *SinglyNode {
	return &SinglyNode{content: content}
}

func (nd *SinglyNode) setNextNode(nextNode *SinglyNode) {
	nd.nextNode = nextNode
}

type SinglyLinkedList struct {
	head *SinglyNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (sl *SinglyLinkedList) isEmpty() bool {
	return sl.head == nil
}

func (sl *SinglyLinkedList) InsertBegin(content interface{}) {
	if sl.isEmpty() {
		sl.head = newSinglyNode(content)
	} else {
		item := newSinglyNode(content)
		item.setNextNode(sl.head)
		sl.head = item
	}
}

func (sl *SinglyLinkedList) InsertEnd(content interface{}) {
	if sl.isEmpty() {
		sl.head = newSinglyNode(content)
	} else {
		lastNode := sl.getLastNode()
		lastNode.nextNode = newSinglyNode(content)
	}
}

func (sl *SinglyLinkedList) Remove(content interface{}) {
	if sl.isEmpty() {
		return
	}

	if sl.head.content == content {
		sl.head = sl.head.nextNode
		return
	}

	prevNode := sl.head
	currentNode := sl.head.nextNode
	for currentNode != nil {
		if currentNode.content == content {
			prevNode.setNextNode(currentNode.nextNode)
			return
		} else {
			prevNode = currentNode
			currentNode = currentNode.nextNode
		}
	}
}

func (sl *SinglyLinkedList) Print() {
	currentNode := sl.head
	for currentNode != nil {
		fmt.Println(currentNode.content)
		currentNode = currentNode.nextNode
	}
}

func (sl *SinglyLinkedList) getLastNode() *SinglyNode {
	currentNode := sl.head
	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}

	return currentNode
}
