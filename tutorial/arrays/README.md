# Arrays

Arrays are the most basic, important and  intutive data strctures to be learnt.

An array is an collection of data of the same type stored in memory in a sequential manner. 

**Note:** Memory of a computer can be considered as blocks which has some address. Data can be stored on the blocks. Sequential in this contex means that if you create an array of 10 items, those 10 items does not take random blocks from memory but in an orderly fashion. 

Every item in an array has an associated index with it and the item can be read/accessed using that index. In Golang the index starts with 0. For example if you have 3 items in an array then the first item has index 0, the second item has index 1 and the third item has index 2.   

Almost all the high level programming languages has array data structure inbuilt. We will see how we can create and manipulate array in golang but before that let us understand few examples and the application of arrays.



# Example 
Following is an example of array of integers:

[1,2,3,4,5,6,7]

Following is an example of array of floats:

[1.5,2.5,3.5,6.5]

# Applications

Some applications of arrays are the following:
- Organise similar data together to do some mathematical operations.
- Arrays are used to implement other data structures.

## Examples

Suppose you have to calculate sum of 10 numbers, what will you do? Will you declare 10 variables and do the sum?
Answer is no, as it will be too much of code and will not be maintainable and prone to errors. We can use an array in this case.

Arrays are also used to implement stack,queue,heap etc data structures.

# Code Examples

The code can also be found in arrays.go file but has been placed here
so that one can do a continuous read.

```go
func main() {
	// Declare and intialize array of integers of size 10
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Print value at index 3 -- this is known as random access.
	// To read the value at index 3 it is not necessary to go through
	// index 0,1 and 2. We will see a data structure where random access
	// is not possible.
	fmt.Println(arr[3])

	// Print all the values of array one by one using a loop
	for i:=0;i<len(arr);i++{
		fmt.Printf("At index %d:%d\n",i,arr[i])
	}

	// Add all the items of the array
	sum:=0
	for i:=0;i<len(arr);i++{
		sum = sum + arr[i]
	}
	fmt.Printf("Sum is %d\n",sum)

	
	// evenList is an array of size 10
	// All the items (i.e from index 0 to 9) will be zero by default
	// Let us populate this array with first 10 even integers.
	var evenList [10]int

	// evenList[0] is 0 -- so this is fine
	// We need to populate from index 1
	for i:=1;i<len(evenList);i++{
		evenList[i]=evenList[i-1]+2
	}
	
	// print the evenList
	for i:=0;i<len(evenList);i++{
		fmt.Println("At index %d:%d",i,evenList[i])
	}
}
```

## Array Bound
An array is bounded by its size. For example let us say that we have an array of size `n` then we can only access elements of indices `0 to n-1` and any index out of that bound is going to end in a error. (Try it yourself and check)
Also, array index cannot be a negative number.

## Operations On Arrays

You can access,add,delete and modify items in an array. But in golang array that we have discussed so far, we can only access and modify them.
Addition and removal cannot be done here and golang arrays are static in nature in this context.

Golang has another array like data structure known as slice that allows you to do all the opeations i.e. access,add,delete and modify. Do not confuse that slice is a new kind of data structure. Slices are nothing but arrays only but they allow a richer set of operations. This is how it is in Golang.
There can exist a language where an array simply allow all the operations.

Please go through the following link to understand the difference between array and slice in golang. It is highly recommended that you understand this if you are programming in Go.
https://blog.golang.org/slices-intro
Infact, above blog contains lot of useful information and I will not repeat it here.

## Slices and Operations

Walkthrough following code snippet to understand the operations. Code snippet is explained via comments.
```go
func sliceTutorial()  {
	fmt.Println("-----------------Slice tutorial------------------")
	// Make a slice of size 0
	sl:=make([]int,0)
	// add 1 item
	sl = append(sl,1)
	// add some more items
	sl = append(sl,2,3,4)
	// print the slice
	fmt.Println(sl)

	// Print the multiplication of all items in the slice
	result:=1
	for i:=0;i<len(sl);i++{
		result = result*sl[i]
	}
	fmt.Println("Mutiplication result:",result)
	// If sl is a slice then sl[k:n] is slice from index k to n-1
	// sl[:end] means --> sl[0:end]
	// sl[start:] means --> sl[start:n] [ n is size of the slice]
	// sl[:] mean ---> sl[0:n] [ n is size of the slice]
	fmt.Println(sl[:])
	fmt.Println(sl[:2])
	fmt.Println(sl[2:])
	fmt.Println(sl[1:4])

	// Remove an item at index 1
	sl = append(sl[0:1],sl[2:]...)
	fmt.Println(sl)
}

```

It is also worth to think about time complexity for operations on the slices and understanding of this blog https://blog.golang.org/slices-intro is required for this.
// ToDo: Answers
1. What is the time complexity for appending an item in a slice?
2. What is the time complexity for deleting an item at index `k` in a slice? (Assume that slice length >= `k`)
3. What is the time conplexity for modifying an item at index `k` in a slice?
4. What is the time complexity for adding an item after `kth` index in a slice? (Assume that slice length >= `k`)

## Experimenting with Arrays

// ToDo: 

## Interview Tip

// ToDo: 




 