package hard

import (
	"testing"
)

func Test_maxTwoEvents(t *testing.T) {
	type args struct {
		events [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test0",
			args: args{
				events: [][]int{},
			},
			want: 0,
		},
		{
			name: "test1",
			args: args{
				events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}},
			},
			want: 7,
		},
		{
			name: "test2",
			args: args{
				events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}},
			},
			want: 10,
		},
		{
			name: "test3",
			args: args{
				events: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			},
			want: 1,
		},
		{
			name: "test4",
			args: args{
				events: [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTwoEvents(tt.args.events); got != tt.want {
				t.Errorf("maxTwoEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildEventKey(t *testing.T) {
	type args struct {
		event []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildEventKey(tt.args.event); got != tt.want {
				t.Errorf("buildEventKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
