package main

import "fmt"

func main()  {
	mapDemo()
}

// mapDemo function illustrates how maps are used
// in golang.
func mapDemo()  {
	// This is how a map is created and initialized in golang
	// using make.
	// Here key type is int and also value type is int
	studentScores:=make(map[int]int)

	// Let say us student with roll number 1 scored 94 marks
	// out of 100. Let us push this in the studentScores map.
	studentScores[1]=100
	// Similarly let us push for students with roll number 2,3,4
	// and 5 who scored 90,80,85 and 65 respectively.
	studentScores[2]=90
	studentScores[3]=80
	studentScores[4]=85
	studentScores[5]=65
	/*
	Above push operations on map also take almost constant time and we can
	say:
	Time Complexity for a Push Operation in Map : O(1)
	*/

	// Let us now try to get score of student with roll number 3
	fmt.Printf("Marks of student with roll number %d:%d\n",3,studentScores[3])
	/*
	We just now fetched score of student with roll number 3 and this is a get operations.
	Time Complexity for a Get Operation in Map: O(1)
	*/

	// Let us now try to delete a student entry from the map. Say we want to
	// delete entry for roll number 1.
	// Use delete inbuilt function and pass the map and key to delete
	// the entry corresponding to the key.
	delete(studentScores,1)
	/*
	Time Complexity for Delete Operation in Map : O(1)
	*/

	// We can print length of a map in following way :
	fmt.Println("Length of Map:",len(studentScores))

	// we can also iterate through all the items in a map
	for k,v:=range studentScores{
		fmt.Printf("Roll Number:%d; Score:%d\n",k,v)
	}

	/*
	It should be noted that looping over the map does not guarantee that
	you will get the map entries in the same order you have pushed. We should
	never assume that items of a map will be in the order in which we have
	pushed and base our coding implementation on this assumption.
	Though it can come sometimes in the same order we have pushed, it is
	just a coincidence. It is just random.
	*/
}