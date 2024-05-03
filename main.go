package main

import "fmt"

func main() {
	//hashTable := NewHashTable()
	//hashTable.Insert("Thalisson", 23)
	//hashTable.Keys()
	//
	//linkedList := NewSinglyLinkedList()
	//
	//linkedList.InsertBegin(52)
	//linkedList.InsertBegin(51)
	//linkedList.InsertBegin(50)
	//linkedList.InsertEnd(53)
	//linkedList.Remove(50)
	//linkedList.Print()

	//doublyLinkedList := NewDoublyLinkedList()
	//doublyLinkedList.InsertEnd(7)
	//doublyLinkedList.InsertEnd(8)
	//doublyLinkedList.InsertBegin(6)
	//doublyLinkedList.InsertMiddle(9, 2)
	//doublyLinkedList.InsertMiddle(10, 3)
	//doublyLinkedList.InsertMiddle(11, 4)
	//doublyLinkedList.InsertEnd(12)
	//doublyLinkedList.Print()
	//fmt.Println("-------------------")
	//doublyLinkedList.Remove(6)
	//doublyLinkedList.Remove(5)
	//doublyLinkedList.Remove(0)
	//doublyLinkedList.RemoveLast()
	//doublyLinkedList.RemoveFirst()
	//doublyLinkedList.Print()

	//stack := NewStack()
	//stack.Push(1)
	//stack.Push(2)
	//stack.Peek()
	//fmt.Println(stack.Pop())

	//queue := NewQueue()
	//queue.Enqueue(1)
	//queue.Enqueue(2)
	//queue.Enqueue(3)
	//queue.Enqueue(4)
	//
	//queue.Dequeue()
	//queue.Dequeue()
	//queue.Dequeue()
	//queue.Dequeue()
	//
	//fmt.Println(queue.Peek())

	binaryTree := NewBinaryTree()
	binaryTree.Insert(5)
	binaryTree.Insert(3)
	binaryTree.Insert(7)
	binaryTree.Insert(4)

	fmt.Println(binaryTree.Find(7).content)
}
