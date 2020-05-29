package maychallenge

import "testing"

func Test_possibleBipartition(t *testing.T) {
	type args struct {
		N        int
		dislikes [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				N:        3,
				dislikes: [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 3}},
			},
			want: false,
		},
		{
			name: "test 2",
			args: args{
				N:        10,
				dislikes: [][]int{[]int{4, 7}, []int{4, 8}, []int{2, 8}, []int{8, 9}, []int{1, 6}, []int{5, 8}, []int{1, 2}, []int{6, 7}, []int{3, 10}, []int{8, 10}, []int{1, 5}, []int{7, 10}, []int{1, 10}, []int{3, 5}, []int{3, 6}, []int{1, 4}, []int{3, 9}, []int{2, 3}, []int{1, 9}, []int{7, 9}, []int{2, 7}, []int{6, 8}, []int{5, 7}, []int{3, 4}},
			},
			want: true,
		},
		{
			name: "test 3",
			args: args{
				N:        4,
				dislikes: [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 4}},
			},
			want: true,
		},
		{
			name: "test 4",
			args: args{
				N:        5,
				dislikes: [][]int{[]int{1, 2}, []int{3, 4}, []int{4, 5}, []int{3, 5}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartition(tt.args.N, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
