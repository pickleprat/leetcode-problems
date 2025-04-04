package main 

import "errors"

type Stack struct {
	top   	int
	stack 	[] byte
}

func (st *Stack) Push(num byte) byte  {
	st.stack = append(st.stack, num); st.top++; 
	return num  
}

func (st *Stack) View() (byte, error) {
	// views the topmost element of the stack 
	if st.top == -1 {
		return 0, errors.New("empty stack. nothing to view");  
	}
	return st.stack[st.top], nil
}

func (st *Stack) Pop() (byte, error) {
	// removes the topmost element of the stack 
	if st.top == -1 {
		return 0, errors.New("empty stack. nothing to remove"); 
	} 

	topByte := st.stack[st.top]; 
	st.stack = st.stack[:len(st.stack)-1]; 
	return topByte, nil  
}  

func NewStack(size int) *Stack {
	return &Stack{
		stack: make([]byte, 0),
		top:   -1,
	}
}