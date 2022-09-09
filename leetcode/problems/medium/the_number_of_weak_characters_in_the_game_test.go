package medium

import "testing"

func Test_numberOfWeakCharacters(t *testing.T) {
	type args struct {
		properties [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				properties: [][]int{
					{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5},
				},
			},
			want: 6,
		},
		{
			args: args{
				properties: [][]int{
					{5, 5}, {6, 3}, {3, 6},
				},
			},
			want: 0,
		},
		{
			args: args{
				properties: [][]int{
					{2, 2}, {3, 3},
				},
			},
			want: 1,
		},
		{
			args: args{
				properties: [][]int{
					{1, 5}, {10, 4}, {4, 3},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWeakCharacters(tt.args.properties); got != tt.want {
				t.Errorf("numberOfWeakCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
