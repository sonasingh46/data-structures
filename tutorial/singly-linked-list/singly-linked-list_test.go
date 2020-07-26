package singlylinkedlist

import (
	"testing"
)

func TestAddNodeToEnd(t *testing.T) {
	type args struct {
		head *Node
		data int
	}
	tests := []struct {
		name           string
		args           args
		want           int
		expectedLength int
	}{
		{
			name: "Test Case #1: Add a node at end of linked list when head is nil",
			args: args{
				head: CreateLinkedListFromArray([]int{}),
				data: 5,
			},
			want:           5,
			expectedLength: 1,
		},

		{
			name: "Test Case #2: Add a node at end of linked list (3->4->8)",
			args: args{
				head: CreateLinkedListFromArray([]int{3, 4, 8}),
				data: 5,
			},
			want:           5,
			expectedLength: 4,
		},

		{
			name: "Test Case #3: Add a node at end of linked list (3)",
			args: args{
				head: CreateLinkedListFromArray([]int{3}),
				data: 7,
			},
			want:           7,
			expectedLength: 2,
		},

		{
			name: "Test Case #3: Add a node at end of linked list (3)",
			args: args{
				head: CreateLinkedListFromArray([]int{3}),
				data: 7,
			},
			want:           7,
			expectedLength: 2,
		},

		{
			name: "Test Case #3: Add a node at end of linked list (24->56->45)",
			args: args{
				head: CreateLinkedListFromArray([]int{24, 56, 45}),
				data: 45,
			},
			want:           45,
			expectedLength: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.head = AddNodeToEnd(tt.args.head, tt.args.data)
			lastNode := GetLastNode(tt.args.head)
			if lastNode.Data != tt.want {
				t.Errorf("Test case failed as: "+
					"want las node data to be %d but got %d", tt.want, lastNode.Data)
			}
			gotLength := GetLength(tt.args.head)
			if gotLength != tt.expectedLength {
				t.Errorf("Test case failed as: "+
					"want expected length be %d but got %d", tt.expectedLength, gotLength)
			}
		})
	}
}

func TestAddNodeToBeg(t *testing.T) {
	type args struct {
		head *Node
		data int
	}
	tests := []struct {
		name           string
		args           args
		want           int
		expectedLength int
	}{
		{
			name: "Test Case #1: Add a node at beginning of linked list when head is nil",
			args: args{
				head: CreateLinkedListFromArray([]int{}),
				data: 5,
			},
			want:           5,
			expectedLength: 1,
		},

		{
			name: "Test Case #2: Add a node at beginning of linked list (3->4->8)",
			args: args{
				head: CreateLinkedListFromArray([]int{3, 4, 8}),
				data: 5,
			},
			want:           5,
			expectedLength: 4,
		},

		{
			name: "Test Case #3: Add a node at beginning of linked list (3)",
			args: args{
				head: CreateLinkedListFromArray([]int{3}),
				data: 7,
			},
			want:           7,
			expectedLength: 2,
		},

		{
			name: "Test Case #3: Add a node at beginning of linked list (3)",
			args: args{
				head: CreateLinkedListFromArray([]int{3}),
				data: 7,
			},
			want:           7,
			expectedLength: 2,
		},

		{
			name: "Test Case #3: Add a node at beginning of linked list (24->56->45)",
			args: args{
				head: CreateLinkedListFromArray([]int{24, 56, 45}),
				data: 45,
			},
			want:           45,
			expectedLength: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.head = AddNodeToBeg(tt.args.head, tt.args.data)
			if tt.args.head.Data != tt.want {
				t.Errorf("Test case failed as: "+
					"want first node data to be %d but got %d", tt.want, tt.args.head.Data)
			}
			gotLength := GetLength(tt.args.head)
			if gotLength != tt.expectedLength {
				t.Errorf("Test case failed as: "+
					"want expected length be %d but got %d", tt.expectedLength, gotLength)
			}
		})
	}
}