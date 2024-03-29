package hard

import "testing"

func Test_maxValue(t *testing.T) {
	type args struct {
		events [][]int
		k      int
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
				k:      1,
			},
			want: 0,
		},
		{
			name: "test1",
			args: args{
				events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}},
				k:      2,
			},
			want: 7,
		},
		{
			name: "test2",
			args: args{
				events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}},
				k:      2,
			},
			want: 10,
		},
		{
			name: "test3",
			args: args{
				events: [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}},
				k:      3,
			},
			want: 9,
		},
		{
			name: "test4",
			args: args{
				events: [][]int{{1, 1, 4}, {2, 5, 10}, {6, 6, 1}, {1, 2, 7}, {3, 4, 4}, {5, 6, 5}},
				k:      4,
			},
			want: 16,
		},
		{
			name: "test5",
			args: args{
				events: [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}},
				k:      2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.events, tt.args.k); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
