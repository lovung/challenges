package easy

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		// from 1 to 10^3
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{2, 3, 4, 5, 1},
			},
			want: 12,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 5, 4, 5},
			},
			want: 16,
		},
		{
			name: "test3",
			args: args{
				nums: []int{3, 7},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
