package main

type Stack struct {
	stack []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (st *Stack) IsEmpty() bool {
	return len(st.stack) == 0
}

func (st *Stack) Size() int {
	return len(st.stack)
}

func (st *Stack) Push(content interface{}) {
	st.stack = append(st.stack, content)
}

func (st *Stack) Pop() (removedItem interface{}) {
	if !st.IsEmpty() {
		removedItem = st.stack[st.Size()-1]
		st.stack = st.stack[:st.Size()-1]
		return
	}

	st.shrink()

	return nil
}

func (st *Stack) Peek() interface{} {
	if !st.IsEmpty() {
		return st.stack[st.Size()-1]
	}

	return nil
}

func (st *Stack) shrink() {
	currentCapacity := cap(st.stack)
	if len(st.stack) < currentCapacity/4 && currentCapacity > 16 {
		// Reduz a capacidade para metade se a pilha estiver utilizando menos de 25% e capacidade for maior que 16
		newStack := make([]interface{}, len(st.stack), st.max(currentCapacity/2, 16))
		copy(newStack, st.stack)
		st.stack = newStack
	}
}

func (st *Stack) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
