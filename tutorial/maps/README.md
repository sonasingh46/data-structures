# Maps

Map is a key value data structure used for fast lookup. It is very important data structure and used extensively in real
world. Before we understand internals of map, let us first understand how does a map look like, what does fast lookup 
mean and how we can use it in Golang.

```
      +-------------+
      | Key  | Value|
      |      |      |
      +-------------+
      |  2   |  8   |
      +-------------+
      |  4   |  9   |
      +-------------+
      |  12  |  18  |
      +-------------+
      |  45  |  56  |
      +-------------+
      |  22  |  78  |
      +-------------+
      |  456 |  59  |
      +-------------+
      |  70  |  80  |
      +-------------+
```

If you look at the above diagram, you can realise a key value structure. We can say that for key `2` the value is `8`,
for key `22` the value is `78` and similarly for others.

Consider an example when you have to save score(out of 100) of 1000 students. Each student has a roll number and you
need to find/tell the score of a student given the roll number. 
Now, what you can do is, use a map. Roll number will be the `key` and score will be the `value`. Now the value
corresponding to a key can be get in almost constant time. This is what fast lookup means. Notice `almost constant time`. 
We will analayse why it is almost constant and not strictly constant in coming sections.

# Maps In Golang

Go thorugh following code to understand more about maps in golang. The code can also be found in the same directory
here in `map.go`

```go
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
```

# Internals of Map

// ToDo: 
