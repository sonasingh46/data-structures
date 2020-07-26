package singlylinkedlist

// Node is the linked list item type
type Node struct {
	// Data to store a integer data
	Data int
	// Next points to the next node in the linked list
	Next *Node
}

// NewNode returns a new instance of Node
func NewNode(data int)*Node{
	return &Node{
		Data:data,
	}
}

// CreateLinkedListFromArray takes array as an input
// and converts it to a linked list.
func CreateLinkedListFromArray(arr []int)*Node  {
	var head *Node
	if len(arr)>0{
		head = NewNode(arr[0])
		for i:=1;i<len(arr);i++{
			AddNodeToEnd(head,arr[i])
		}
	}
	return head
}

// AddNodeToEnd adds a node to the end of the linked list.
func AddNodeToEnd(head *Node, data int) *Node {
	newNode:=NewNode(data)
	if head==nil{
		return newNode
	}
	temp:=head
	for temp.Next!=nil{
		temp = temp.Next
	}
	temp.Next=newNode
	return head
}

// AddNodeToBeg adds a node to the begining of the linked list.
func AddNodeToBeg(head *Node, data int)*Node  {
	newNode:=NewNode(data)
	newNode.Next=head
	head=newNode
	return head
}

// GetLastNode returns last node of the linked list.
func GetLastNode(head *Node)*Node  {
	if head==nil{
		return nil
	}
	temp:=head
	for temp.Next!=nil{
		temp=temp.Next
	}
	return temp
}

// GetLength returns number of nodes in the linked list
func GetLength(head *Node)int  {
	var count int
	if head==nil{
		return count
	}
	temp:=head
	for temp!=nil{
		count++
		temp=temp.Next
	}
	return count
}