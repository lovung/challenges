package weekly395

import "testing"

func Test_medianOfUniquenessArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{3, 4, 3, 4, 5},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{4, 3, 5, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianOfUniquenessArray(tt.args.nums); got != tt.want {
				t.Errorf("medianOfUniquenessArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
