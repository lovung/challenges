package problems

import "testing"

func Test_singleNonDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			},
			want: 2,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{3, 3, 7, 7, 10, 11, 11},
			},
			want: 10,
		},
		{
			name: "test 3",
			args: args{
				nums: []int{1, 1, 2, 3, 3},
			},
			want: 2,
		},
		{
			name: "test 4",
			args: args{
				nums: []int{1, 2, 2, 3, 3},
			},
			want: 1,
		},
		{
			name: "test 5",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "test 6",
			args: args{
				nums: []int{1, 1, 2, 2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNonDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
