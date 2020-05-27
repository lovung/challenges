package interview

import "testing"

func Test_checkTriangularArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				nums: []int{1, 3, 5, 7, 9, 8, 6, 4, 2, 0},
			},
			want: true,
		},
		{
			name: "test 4",
			args: args{
				nums: []int{1, 3, 5, 7, 9, 8, 5, 4, 2, 0},
			},
			want: false,
		},
		{
			name: "test 5",
			args: args{
				nums: []int{1, 3, 5, 7, 9, 7, 6, 4, 2, 0},
			},
			want: false,
		},
		{
			name: "test 6",
			args: args{
				nums: []int{1, 3, 5, 7, 9, 7, 6, 4, 2, 100},
			},
			want: false,
		},
		{
			name: "test 7",
			args: args{
				nums: []int{1, 3, 5, 7, 9, 9, 6, 4, 2, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkTriangularArray(tt.args.nums); got != tt.want {
				t.Errorf("checkTriangularArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
