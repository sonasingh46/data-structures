package dequeue

// Dequeue is a double ended queue
type Dequeue struct {
	Front *Node
	Rear  *Node
}

// NewDequeue returns an empty instance of Dequeue
func NewDequeue() *Dequeue {
	return &Dequeue{}
}

// Node is the element of a dequeue.
type Node struct {
	Data int
	Next *Node
	Prev *Node
}

// NewNode returns a new node.
func NewNode(data int) *Node {
	return &Node{
		Data: data,
	}
}

// GetFront returns the front node from the dequeue.
// Note: This methods should not be called on empty dequeue as it panics.
func (d *Dequeue) GetFront() *Node {
	if d.IsEmpty() {
		panic("Get front on empty queue!")
	}
	return d.Front
}

// PushFront pushes a node to the front of the dequeue.
func (d *Dequeue) PushFront(data int) {
	newNode := NewNode(data)
	if d.IsEmpty() {
		d.Front = newNode
		d.Rear = newNode
	} else {
		newNode.Next = d.Front
		d.Front.Prev = newNode
		d.Front = newNode
	}
}

// PopFront pops a node from the front of the dequeue.
// Note: This method should not be called on empty dequeue as it panics
func (d *Dequeue) PopFront() *Node {
	if d.IsEmpty() {
		panic("Pop front on empty queue!")
	}
	temp := d.GetFront()
	// If only one node is present
	if d.Front == d.Rear {
		d.Front = nil
		d.Rear = nil
		return temp
	}

	d.Front = d.Front.Next
	d.Front.Prev = nil

	return temp
}

// GetRear returns the rear node from the dequeue.
// Note: This methods should not be called on empty dequeue as it panics.
func (d *Dequeue) GetRear() *Node {
	if d.IsEmpty() {
		panic("Get rear on empty queue!")
	}
	return d.Rear
}

// PushRear pushes a node to the rear of the dequeue.
func (d *Dequeue) PushRear(data int) {
	newNode := NewNode(data)
	if d.IsEmpty() {
		d.Front = newNode
		d.Rear = newNode
	} else {
		d.Rear.Next = newNode
		newNode.Prev = d.Rear
		d.Rear = newNode
	}
}

// PopRear pops a node from the rear of the dequeue.
// Note: This method should not be called on empty dequeue as it panics
func (d *Dequeue) PopRear() *Node {
	if d.IsEmpty() {
		panic("Pop rear on empty queue!")
	}
	temp := d.GetRear()
	// If only one node is present
	if d.Front == d.Rear {
		d.Rear = nil
		d.Front = nil
		return temp
	}
	d.Rear = d.Rear.Prev
	d.Rear.Next = nil
	return temp
}

// IsEmpty returns true if the dequeue is empty
func (d *Dequeue) IsEmpty() bool {
	return d.Front == nil || d.Rear == nil
}
