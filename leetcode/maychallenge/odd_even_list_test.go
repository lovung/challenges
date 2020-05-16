package maychallenge

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test 1",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "TestCase01",
			args: args{
				head: newListNode([]int{1, 2, 3, 4, 5}),
			},
			want: newListNode([]int{1, 3, 5, 2, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
