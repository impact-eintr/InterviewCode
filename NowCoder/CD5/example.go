package CD5

type MinStack struct {
	data     []interface{}
	minIndex []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var instance MinStack
	return instance
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.minIndex) == 0 {
		this.minIndex = append(this.minIndex, 0)
	} else {
		//当前最小值比新值大
		if this.data[this.minIndex[len(this.minIndex)-1]].(int) > x {
			//当前最小值索引压入minIndex
			this.minIndex = append(this.minIndex, len(this.data)-1)
		}
	}
}

func (this *MinStack) Pop() {
	length := len(this.data)
	currentMinIndex := length - 1

	if currentMinIndex == this.minIndex[len(this.minIndex)-1] {
		this.minIndex = this.minIndex[:len(this.minIndex)-1]
	}

	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1].(int)
}

func (this *MinStack) GetMin() int {
	return this.data[this.minIndex[len(this.minIndex)-1]].(int)
}
