package dequeue

import (
	"reflect"
	"testing"
)

func TestDequeue_GetFront(t *testing.T) {

	tests := []struct {
		name string
		dq   *Dequeue
		want int
	}{
		{
			name: "Test case #1",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				firstNode:=NewNode(80)
				dq.Front=firstNode
				dq.Rear=firstNode
				return dq
			}(),
			want: 80,
		},

		{
			name: "Test case #2",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				firstNode:=NewNode(5)
				secondNode:=NewNode(6)

				firstNode.Next=secondNode

				secondNode.Prev=firstNode

				dq.Front=firstNode
				dq.Rear=secondNode

				return dq
			}(),
			want: 5,
		},

		{
			name: "Test case #3",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				firstNode:=NewNode(20)
				secondNode:=NewNode(35)
				thirdNode:=NewNode(11)

				firstNode.Next=secondNode
				secondNode.Next=thirdNode

				thirdNode.Prev=secondNode
				secondNode.Prev=firstNode

				dq.Front=firstNode
				dq.Rear=thirdNode
				return dq
			}(),
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dq.GetFront().Data; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dequeue.GetFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDequeue_GetRear(t *testing.T) {

	tests := []struct {
		name string
		dq   *Dequeue
		want int
	}{
		{
			name: "Test case #1",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				firstNode:=NewNode(80)
				dq.Front=firstNode
				dq.Rear=firstNode
				return dq
			}(),
			want: 80,
		},

		{
			name: "Test case #2",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				firstNode:=NewNode(5)
				secondNode:=NewNode(6)

				firstNode.Next=secondNode

				secondNode.Prev=firstNode

				dq.Front=firstNode
				dq.Rear=secondNode

				return dq
			}(),
			want: 6,
		},

		{
			name: "Test case #3",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				firstNode:=NewNode(20)
				secondNode:=NewNode(35)
				thirdNode:=NewNode(11)

				firstNode.Next=secondNode
				secondNode.Next=thirdNode

				thirdNode.Prev=secondNode
				secondNode.Prev=firstNode

				dq.Front=firstNode
				dq.Rear=thirdNode
				return dq
			}(),
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dq.GetRear().Data; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dequeue.GetFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDequeue_PushFront(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name     string
		dq       *Dequeue
		args     args
		wantData []int
	}{
		{
			name: "Test Case #1",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				return dq
			}(),
			args: args{
				data: 7,
			},
			wantData: []int{7},
		},

		{
			name: "Test Case #2",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				dq.Front = NewNode(2)
				dq.Rear = dq.Front
				return dq
			}(),
			args: args{
				data: 7,
			},
			wantData: []int{7, 2},
		},

		{
			name: "Test Case #2",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				firstNode := NewNode(1)
				secondNode := NewNode(2)
				thirdNode := NewNode(3)
				fourthNode := NewNode(4)

				firstNode.Next = secondNode
				secondNode.Next = thirdNode
				thirdNode.Next = fourthNode

				fourthNode.Prev = thirdNode
				thirdNode.Prev = secondNode
				secondNode.Prev = firstNode

				dq.Front = firstNode
				dq.Rear = fourthNode

				return dq
			}(),
			args: args{
				data: 7,
			},
			wantData: []int{7, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dq.PushFront(tt.args.data)
			index := 0
			head := tt.dq.Front
			for head.Next != nil {
				if head.Data != tt.wantData[index] {
					t.Errorf("PushFront():Test case failed as "+
						"want %d but got %d", tt.wantData[index], head.Data)
				}
				index++
				head = head.Next
			}
		})
	}
}

func TestDequeue_PushRear(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name     string
		dq       *Dequeue
		args     args
		wantData []int
	}{
		{
			name: "Test Case #1",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				return dq
			}(),
			args: args{
				data: 7,
			},
			wantData: []int{7},
		},

		{
			name: "Test Case #2",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				dq.Front = NewNode(2)
				dq.Rear = dq.Front
				return dq
			}(),
			args: args{
				data: 7,
			},
			wantData: []int{2, 7},
		},

		{
			name: "Test Case #2",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				firstNode := NewNode(1)
				secondNode := NewNode(2)
				thirdNode := NewNode(3)
				fourthNode := NewNode(4)

				firstNode.Next = secondNode
				secondNode.Next = thirdNode
				thirdNode.Next = fourthNode

				fourthNode.Prev = thirdNode
				thirdNode.Prev = secondNode
				secondNode.Prev = firstNode

				dq.Front = firstNode
				dq.Rear = fourthNode

				return dq
			}(),
			args: args{
				data: 7,
			},
			wantData: []int{1, 2, 3, 4, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dq.PushRear(tt.args.data)
			index := 0
			head := tt.dq.Front
			for head.Next != nil {
				if head.Data != tt.wantData[index] {
					t.Errorf("PushFront():Test case failed as "+
						"want %d but got %d", tt.wantData[index], head.Data)
				}
				index++
				head = head.Next
			}
		})
	}
}

func TestDequeue_PopRear(t *testing.T) {
	tests := []struct {
		name     string
		dq       *Dequeue
		wantData int
	}{

		{
			name: "Test Case #1",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				dq.Front = NewNode(2)
				dq.Rear = dq.Front
				return dq
			}(),
			wantData: 2,
		},

		{
			name: "Test Case #2",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				firstNode := NewNode(1)
				secondNode := NewNode(2)
				thirdNode := NewNode(3)
				fourthNode := NewNode(4)

				firstNode.Next = secondNode
				secondNode.Next = thirdNode
				thirdNode.Next = fourthNode

				fourthNode.Prev = thirdNode
				thirdNode.Prev = secondNode
				secondNode.Prev = firstNode

				dq.Front = firstNode
				dq.Rear = fourthNode

				return dq
			}(),

			wantData: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode := tt.dq.PopRear()
			if gotNode.Data != tt.wantData {
				t.Errorf("PopRear(): test case failed as "+
					"want %d but got %d for popped node data", tt.wantData, gotNode.Data)
			}
		})
	}
}

func TestDequeue_PopFront(t *testing.T) {
	tests := []struct {
		name     string
		dq       *Dequeue
		wantData int
	}{

		{
			name: "Test Case #1",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				dq.Front = NewNode(2)
				dq.Rear = dq.Front
				return dq
			}(),
			wantData: 2,
		},

		{
			name: "Test Case #2",
			dq: func() *Dequeue {
				dq := &Dequeue{}
				firstNode := NewNode(1)
				secondNode := NewNode(2)
				thirdNode := NewNode(3)
				fourthNode := NewNode(4)

				firstNode.Next = secondNode
				secondNode.Next = thirdNode
				thirdNode.Next = fourthNode

				fourthNode.Prev = thirdNode
				thirdNode.Prev = secondNode
				secondNode.Prev = firstNode

				dq.Front = firstNode
				dq.Rear = fourthNode

				return dq
			}(),

			wantData: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode := tt.dq.PopFront()
			if gotNode.Data != tt.wantData {
				t.Errorf("PopFront(): test case failed as "+
					"want %d but got %d for popped node data", tt.wantData, gotNode.Data)
			}
		})
	}
}

func TestDequeue_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		dq *Dequeue
		want   bool
	}{
		{
			name:"Test Case #1",
			dq:NewDequeue(),
			want:true,
		},

		{
			name:"Test Case #2",
			dq: func() *Dequeue{
				dq:=NewDequeue()
				dq.PushFront(1)
				return dq
			}(),
			want:false,
		},
		{
			name:"Test Case #3",
			dq: func() *Dequeue{
				dq:=NewDequeue()
				dq.PushRear(1)
				return dq
			}(),
			want:false,
		},

		{
			name:"Test Case #4",
			dq: func() *Dequeue{
				dq:=NewDequeue()
				dq.PushRear(5)
				dq.PushRear(7)
				return dq
			}(),
			want:false,
		},

		{
			name:"Test Case #5",
			dq: func() *Dequeue{
				dq:=NewDequeue()
				dq.PushFront(8)
				dq.PushFront(10)
				return dq
			}(),
			want:false,
		},

		{
			name:"Test Case #6",
			dq: func() *Dequeue{
				dq:=NewDequeue()
				dq.PushFront(15)
				dq.PushRear(20)
				return dq
			}(),
			want:false,
		},

		{
			name:"Test Case #7",
			dq: func() *Dequeue{
				dq:=NewDequeue()
				dq.PushRear(34)
				dq.PushFront(88)
				return dq
			}(),
			want:false,
		},

		{
			name:"Test Case #7",
			dq: func() *Dequeue{
				dq:=NewDequeue()
				dq.PushRear(34)
				dq.PushFront(88)
				dq.PushFront(88)
				dq.PushFront(88)
				dq.PopFront()
				dq.PopFront()
				dq.PopFront()
				dq.PopFront()
				return dq
			}(),
			want:true,
		},

		{
			name:"Test Case #8",
			dq: func() *Dequeue{
				dq:=NewDequeue()
				dq.PushRear(34)
				dq.PushFront(88)
				dq.PushFront(88)
				dq.PushFront(88)
				dq.PopFront()
				dq.PopRear()
				dq.PopRear()
				dq.PopFront()
				return dq
			}(),
			want:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dq.IsEmpty(); got != tt.want {
				t.Errorf("Dequeue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
