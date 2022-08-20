package medium

import "testing"

func Test_maxEvents(t *testing.T) {
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
				events: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}},
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				events: [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}},
			},
			want: 7,
		},
		{
			name: "test3",
			args: args{
				events: [][]int{{1, 2}, {2, 3}, {3, 4}},
			},
			want: 3,
		},
		{
			name: "test4",
			args: args{
				events: [][]int{{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5}},
			},
			want: 5,
		},
		{
			name: "test5",
			args: args{
				events: [][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}},
			},
			want: 5,
		},
		{
			name: "test6",
			args: args{
				events: [][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}},
			},
			want: 4,
		},
		{
			name: "test7",
			args: args{
				events: [][]int{{1, 2}, {1, 2}, {1, 6}, {1, 2}, {1, 2}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEvents(tt.args.events); got != tt.want {
				t.Errorf("maxEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
