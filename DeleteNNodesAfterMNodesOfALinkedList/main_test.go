package DeleteNNodesAfterMNodesOfALinkedList

import (
	"reflect"
	"testing"
)

func Test_deleteNodes(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "standard", want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 11, Next: &ListNode{Val: 12}}}}}}, args: args{n: 3, m: 2, head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10, Next: &ListNode{Val: 11, Next: &ListNode{Val: 12, Next: &ListNode{Val: 13}}}}}}}}}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNodes(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
