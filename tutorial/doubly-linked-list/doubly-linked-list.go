package doublylinkedlist

// DoublyLinkedList data structure
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	length int
}

// NewDoublyLinkedList returns an empty doubly linked list
func NewDoublyLinkedList()*DoublyLinkedList{
	return &DoublyLinkedList{}
}

// Node of the doubly linked list
type Node struct {
	Data int
	Prev *Node
	Next *Node
}

// NewNode returns a new Node.
func NewNode(data int)*Node  {
	return &Node{
		Data:data,
	}
}

// AddNodeToEnd adds a node to the end of doubly linked list.
func (dll *DoublyLinkedList)AddNodeToEnd(data int)  {
	newNode:=NewNode(data)
	dll.length++
	if dll.Head==nil{
		dll.Head=newNode
		dll.Tail=newNode
		return
	}
	dll.Tail.Next=newNode
	newNode.Prev=dll.Tail
	dll.Tail=newNode
}

// AddNodeToBeg adds a node to the beginning of the doubly linked list.
func (dll *DoublyLinkedList)AddNodeToBeg(data int)  {
	newNode:=NewNode(data)
	dll.length++
	if dll.Head==nil{
		dll.Head=newNode
		dll.Tail=newNode
		return
	}
	newNode.Next=dll.Head
	dll.Head.Prev=newNode
	dll.Head=newNode
}

// Len return the length of doubly linked list.
func (dll *DoublyLinkedList)Len() int  {
	return dll.length
}

// ConvertArrayToDLL converts a given array into doubly linked list
func ConvertArrayToDLL(arr []int)*DoublyLinkedList  {
	newDLL:=NewDoublyLinkedList()
	for i:=0;i<len(arr);i++{
		newDLL.AddNodeToEnd(arr[i])
	}
	return newDLL
}
