/*
	https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/529/week-2/3292/

	Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

	push(x) -- Push element x onto stack.
	pop() -- Removes the element on top of the stack.
	top() -- Get the top element.
	getMin() -- Retrieve the minimum element in the stack.


	Example:

	MinStack minStack = new MinStack();
	minStack.push(-2);
	minStack.push(0);
	minStack.push(-3);
	minStack.getMin();   --> Returns -3.
	minStack.pop();
	minStack.top();      --> Returns 0.
	minStack.getMin();   --> Returns -2.

		* Your MinStack object will be instantiated and called as such:
		* obj := Constructor();
		* obj.Push(x);
		* obj.Pop();
		* param_3 := obj.Top();
		* param_4 := obj.GetMin();
*/
package main

type MinStack struct {
	stack    []int
	stackMin []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		stackMin: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.stackMin) == 0 {
		this.stackMin = append(this.stackMin, x)
		return
	}
	top := this.GetMin()
	if top > x {
		this.stackMin = append(this.stackMin, x)
	} else {
		this.stackMin = append(this.stackMin, top)
	}
}

func (this *MinStack) Pop() {
	this.stackMin = this.stackMin[:len(this.stackMin)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.stackMin[len(this.stackMin)-1]
}
