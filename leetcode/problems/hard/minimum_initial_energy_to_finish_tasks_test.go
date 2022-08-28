package hard

import "testing"

func Test_minimumEffort(t *testing.T) {
	type args struct {
		tasks [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "done-with-no-energy-remain",
			args: args{
				tasks: [][]int{{1,2},{2,4},{4,8}},
			},
			want: 8,
		},
		{
			name: "normal-case",
			args: args{
				tasks: [][]int{{1,3},{2,4},{10,11},{10,12},{8,9}},
			},
			want: 32,
		},
		{
			name: "normal-case-2",
			args: args{
				tasks: [][]int{{1,7},{2,8},{3,9},{4,10},{5,11},{6,12}},
			},
			want: 27,
		},
		{
			name: "edge-case-1",
			args: args{
				tasks: [][]int{{1,1},{1,3}},
			},
			want: 3,
		},
		{
			name: "edge-case-2",
			args: args{
				tasks: [][]int{{1,1},{1,4},{1,4}},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffort(tt.args.tasks); got != tt.want {
				t.Errorf("minimumEffort() = %v, want %v", got, tt.want)
			}
		})
	}
}
