package main

import "fmt"

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
		fmt.Printf("At index %d:%d\n",i,evenList[i])
	}

	sliceTutorial()
}

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
