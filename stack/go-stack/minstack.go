package main 

type MinStack struct {
    stack [] int 
    minStack [] int 
}


func Constructor() MinStack {
    return MinStack{
        stack : make([] int, 0), 
        minStack: make([] int, 0), 
    }
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val); 
    if (len(this.minStack) > 0 && this.minStack[len(this.minStack) - 1] >= val) || len(this.minStack) == 0 {
        this.minStack = append(this.minStack, val); 
    }
}


func (this *MinStack) Pop()  {
    top := this.Top(); 
    this.stack = this.stack[:len(this.stack) - 1]
    if top == this.GetMin() {
        this.minStack = this.minStack[:len(this.minStack) - 1]; 
    }

}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack) - 1]; 
}


func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack) - 1] ; 
}
