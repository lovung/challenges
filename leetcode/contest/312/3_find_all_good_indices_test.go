package contest

import (
	"reflect"
	"testing"
)

func Test_goodIndices(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{2, 1, 1, 1, 3, 4, 1},
				k:    2,
			},
			want: []int{2, 3},
		},
		{
			args: args{
				nums: []int{2, 1, 2, 1, 2, 1, 2},
				k:    2,
			},
			want: []int{2, 4},
		},
		{
			args: args{
				nums: []int{2, 1, 2, 3, 1, 2, 1, 2},
				k:    2,
			},
			want: []int{5},
		},
		{
			args: args{
				nums: []int{478184, 863008, 716977, 921182, 182844, 350527, 541165, 881224},
				k:    1,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodIndices(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goodIndices() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := goodIndices2(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goodIndices2() = %v, want %v", got, tt.want)
			}
		})
	}
}
