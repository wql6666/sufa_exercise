package easy

type MinStack struct {
	Stack    []int
	MinStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack:    make([]int, 0),
		MinStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	length := len(this.Stack)
	if length == 0 {
		this.MinStack = append(this.MinStack, val)
		this.Stack = append(this.Stack, val)
		return
	}
	if val < this.MinStack[length-1] {
		this.MinStack = append(this.MinStack, val)
	} else {
		this.MinStack = append(this.MinStack, this.MinStack[length-1])
	}
	this.Stack = append(this.Stack, val)

}

func (this *MinStack) Pop() {
	if len(this.Stack) > 0 {
		this.Stack = this.Stack[:len(this.Stack)-1]
		this.MinStack = this.MinStack[:len(this.MinStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.Stack) > 0 {
		return this.Stack[len(this.Stack)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.MinStack) > 0 {
		return this.MinStack[len(this.MinStack)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/*二位数组解决
type MinStack struct {
	vals [][]int
}

 //initialize your data structure here.
func Constructor() MinStack {
	return MinStack{vals: [][]int{}}
}


func (this *MinStack) Push(x int)  {
	if len(this.vals) == 0 {
		this.vals = [][]int{{x,x}}
	} else {
		this.vals = append(this.vals, []int{x, min(x, this.GetMin())})
	}
}


func (this *MinStack) Pop()  {
	this.vals = this.vals[:len(this.vals)-1]
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1][0]
}


func (this *MinStack) GetMin() int {
	return this.vals[len(this.vals)-1][1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
*/
