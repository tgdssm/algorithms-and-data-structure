package main

import "fmt"

type DoublyNode struct {
	content  interface{}
	nextNode *DoublyNode
	prevNode *DoublyNode
}

func newDoublyNode(content interface{}) *DoublyNode {
	return &DoublyNode{content: content}
}

func (nd *DoublyNode) setNextNode(nextNode *DoublyNode) {
	nd.nextNode = nextNode
}

func (nd *DoublyNode) setPrevNode(prevNode *DoublyNode) {
	nd.prevNode = prevNode
}

type DoublyLinkedList struct {
	head *DoublyNode
	tail *DoublyNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dl *DoublyLinkedList) isEmpty() bool {
	return dl.head == nil
}

func (dl *DoublyLinkedList) InsertBegin(content interface{}) {
	if dl.isEmpty() {
		node := newDoublyNode(content)
		dl.head = node
		dl.tail = node
	} else {
		node := newDoublyNode(content)
		node.setNextNode(dl.head)
		dl.head.setPrevNode(node)
		dl.head = node
	}
}

func (dl *DoublyLinkedList) InsertMiddle(content interface{}, index int) {

	if index < 0 || index >= dl.Size() {
		return
	}

	if dl.isEmpty() {
		node := newDoublyNode(content)
		dl.head = node
		dl.tail = node
	} else if index == 0 {
		dl.InsertBegin(content)
	} else {
		currentNode := dl.head
		for i := 0; i < index-1; i++ {
			currentNode = currentNode.nextNode
		}
		newNode := newDoublyNode(content)
		newNode.setPrevNode(currentNode)
		newNode.setNextNode(currentNode.nextNode)
		if currentNode.nextNode != nil {
			currentNode.nextNode.setPrevNode(newNode)
		}
		currentNode.setNextNode(newNode)
	}
}

func (dl *DoublyLinkedList) InsertEnd(content interface{}) {
	if dl.isEmpty() {
		node := newDoublyNode(content)
		dl.head = node
		dl.tail = node
	} else {
		node := newDoublyNode(content)
		node.setPrevNode(dl.tail)
		dl.tail.setNextNode(node)
		dl.tail = node
	}
}

func (dl *DoublyLinkedList) RemoveFirst() {
	if dl.isEmpty() {
		return
	}

	if dl.head.nextNode == nil {
		dl.head = nil
		dl.tail = nil
	}
	dl.head = dl.head.nextNode
	dl.head.setPrevNode(nil)
}

func (dl *DoublyLinkedList) RemoveLast() {
	if dl.isEmpty() {
		return
	}

	if dl.tail.prevNode == nil {
		dl.head = nil
		dl.tail = nil
	}

	dl.tail = dl.tail.prevNode
	dl.tail.setNextNode(nil)
}

func (dl *DoublyLinkedList) Remove(index int) {

	size := dl.Size()
	if index < 0 || index >= size || dl.isEmpty() {
		return
	}

	if index == 0 {
		dl.RemoveFirst()
	} else if index == size-1 {
		dl.RemoveLast()
	} else {
		currentNode := dl.head
		for i := 0; i < index-1; i++ {
			currentNode = currentNode.nextNode
		}

		remove := currentNode.nextNode
		currentNode.setNextNode(remove.nextNode)
		if remove.nextNode != nil {
			remove.nextNode.setPrevNode(currentNode)
		}
	}
}

func (dl *DoublyLinkedList) Clear() {
	dl.head = nil
	dl.tail = nil
}

func (dl *DoublyLinkedList) Print() {
	currentNode := dl.head
	for currentNode != nil {
		fmt.Println(currentNode.content)
		currentNode = currentNode.nextNode
	}
}

func (dl *DoublyLinkedList) Size() (size int) {
	currentNode := dl.head
	for currentNode != nil {
		size += 1
		currentNode = currentNode.nextNode
	}

	return
}
