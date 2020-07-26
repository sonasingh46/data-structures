# Singly Linked List

We will refer singly linked list as linked list in this tutorial and hence single linked list or linked list will mean 
the same.

Linked list like array is a collection of same type of items but unlike of arrays the items are not stored in contiguous
memory but are stored randomly in the memory. The items in an array is accessed by index but the items in a linked list 
is accesed by a pointer. An item of a linked list is also most commonly called as `node`.

In single linked list each item has a pointer to the next item in the list except the last item which does not point to 
anything. So if you have pointer to the first item of a linked list, you can traverse/access the whole linked list. We 
can say that an item(or node) of a linked list has two parts:
- **Data Part:** It holds the data that we want to store.
- **Pointer Part:** It points to the next node in the linked list.

 Linked list is created using self referential structure. 


```
+---------------+        +-------+-------+         +-------+-------+         +-------+-------+         +-------+-------+
|XXXXXXX|       |        |       |       |         |       |       |         |       |       |         |       |       |
|XXXXXXX|  2004 +------->+ Data  |  2008 +-------->+ Data  |  4000 +-------->+ Data  |  3000 +-------->+ Data  |  nil  |
|XXXXXXX|       |        |       |       |         |       |       |         |       |       |         |       |       |
+---------------+        +-------+-------+         +-------+-------+         +-------+-------+         +-------+-------+

                              Node1                     Node2                     Node3                     Node4
      Head                Address:2004              Address:2008              Address:4000              Address:3000
```

```go
// Self referential structure
// Node the item type of a linked list 
type Node struct {
    // This is the data part of the node
    Data int
    Next *Node
}
```
`Head` is of type Node that points to the first item of the linked list. The `Data` part of head is not used.

# Operations On Linked List
- Add a node in linked list.
- Delete a node in linked list.
- Search a node in linked list.

To understand the above operations on linked list go through the code. The code is extensively commented and exaplained
for easier understanding. Code is present in `singly-linked-list.go` file in the current directory.