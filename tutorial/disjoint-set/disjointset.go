package disjointset

// Node is the element of the disjoint set.
type Node struct {
	Data int
	Rank int
	Parent *Node
}

// NewNode returns a new instance of Node.
func NewNode(data int)*Node  {
	return &Node{
		Data:data,
	}
}

// DisjointSet data structure.
type DisjointSet struct {
	Set map[int]*Node
	totalDisjointSets int
}

// NewDisjointSet returns a new disjoint set.
func NewDisjointSet()*DisjointSet  {
	return &DisjointSet{
		Set: map[int]*Node{},
	}
}

// MakeSet pushes a node into a brand new set.
func (ds *DisjointSet)MakeSet(data int)  {
	newNode:=NewNode(data)
	newNode.Parent=newNode
	ds.Set[data]=newNode
	ds.totalDisjointSets++
}

func (ds *DisjointSet)Union(data1,data2 int)  {
	// Get the nodes corresponding to data1 and data2
	node1:=ds.Set[data1]
	node2:=ds.Set[data2]

	// Get parent of node1 and node2
	parent1:=node1.Parent
	parent2:=node2.Parent

	// If both the parents are same, return

	if parent1.Data == parent2.Data {
		return
	}

	if parent1.Rank >= parent2.Rank{
		if parent1.Rank == parent2.Rank{
			parent1.Rank++
		}
		parent2.Parent = parent1
	}else {
		parent1.Parent = parent2
	}
	ds.totalDisjointSets--
}

func (ds *DisjointSet)Find(data int) int {
	return ds.FindNode(ds.Set[data]).Data
}

func (ds *DisjointSet)FindNode(node *Node) *Node {
	parent:=node.Parent
	if parent == node{
		return parent
	}
	// Does path compression here
	node.Parent = ds.FindNode(node.Parent)
	return node.Parent
}