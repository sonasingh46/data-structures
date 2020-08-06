package doublylinkedlist

import (
	"testing"
)

func TestConvertArrayToDLL(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case #1",
			args: args{
				arr: []int{1},
			},
		},
		{
			name: "Test Case #2",
			args: args{
				arr: []int{1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertArrayToDLL(tt.args.arr)
			dllHead := got.Head
			index := 0
			for dllHead != nil {
				if dllHead.Data != tt.args.arr[index] {
					t.Errorf("Test case failed as "+
						"want %d in DLL but got %d", tt.args.arr[index], dllHead.Data)
				}
				dllHead = dllHead.Next
				index++
			}
		})
	}
}

func TestDoublyLinkedList_AddNodeToEnd(t *testing.T) {
	type args struct {
		data int
		DLL *DoublyLinkedList
	}
	tests := []struct {
		name   string
		args   args
		wantLength int
	}{
		{
			name:"Test Case #1",
			args:args{
				DLL:ConvertArrayToDLL([]int{1,2,3,4}),
				data:5,
			},
			wantLength:5,
		},

		{
			name:"Test Case #2",
			args:args{
				DLL:ConvertArrayToDLL([]int{}),
				data:5,
			},
			wantLength:1,
		},

		{
			name:"Test Case #3",
			args:args{
				DLL:ConvertArrayToDLL(nil),
				data:5,
			},
			wantLength:1,
		},

		{
			name:"Test Case #4",
			args:args{
				DLL:ConvertArrayToDLL([]int{34,56,78}),
				data:200,
			},
			wantLength:4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.DLL.AddNodeToEnd(tt.args.data)
			if tt.args.DLL.Tail.Data!=tt.args.data{
				t.Errorf("Test case failed as " +
					"expected %d at end but got %d",tt.args.data,tt.args.DLL.Tail.Data)
			}

			if tt.args.DLL.Len()!=tt.wantLength{
				t.Errorf("Test case failed as " +
					"want length %d but got %d",tt.args.DLL.Len(),tt.wantLength)
			}
		})
	}
}

func TestDoublyLinkedList_AddNodeToBeg(t *testing.T) {
	type args struct {
		data int
		DLL *DoublyLinkedList
	}
	tests := []struct {
		name   string
		args   args
		wantLength int
	}{
		{
			name:"Test Case #1",
			args:args{
				DLL:ConvertArrayToDLL([]int{1,2,3,4}),
				data:5,
			},
			wantLength:5,
		},

		{
			name:"Test Case #2",
			args:args{
				DLL:ConvertArrayToDLL([]int{}),
				data:5,
			},
			wantLength:1,
		},

		{
			name:"Test Case #3",
			args:args{
				DLL:ConvertArrayToDLL(nil),
				data:5,
			},
			wantLength:1,
		},

		{
			name:"Test Case #4",
			args:args{
				DLL:ConvertArrayToDLL([]int{34,56,78}),
				data:200,
			},
			wantLength:4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.DLL.AddNodeToBeg(tt.args.data)
			if tt.args.DLL.Head.Data!=tt.args.data{
				t.Errorf("Test case failed as " +
					"expected %d at end but got %d",tt.args.data,tt.args.DLL.Tail.Data)
			}

			if tt.args.DLL.Len()!=tt.wantLength{
				t.Errorf("Test case failed as " +
					"want length %d but got %d",tt.args.DLL.Len(),tt.wantLength)
			}
		})
	}
}
