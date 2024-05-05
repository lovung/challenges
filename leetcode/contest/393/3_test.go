package contest393

import "testing"

func Test_findKthSmallest(t *testing.T) {
	type args struct {
		coins []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				coins: []int{3, 6, 5, 13, 9},
				k:     1000,
			},
			want: 1970,
		},
		{
			args: args{
				coins: []int{5, 2},
				k:     7,
			},
			want: 12,
		},
		{
			args: args{
				coins: []int{3, 6, 9},
				k:     3,
			},
			want: 9,
		},
		{
			args: args{
				coins: []int{5},
				k:     7,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthSmallest_1(tt.args.coins, tt.args.k); got != tt.want {
				t.Errorf("findKthSmallest_1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthSmallest_2(tt.args.coins, tt.args.k); got != tt.want {
				t.Errorf("findKthSmallest_2a() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthSmallest_3(tt.args.coins, tt.args.k); got != tt.want {
				t.Errorf("findKthSmallest_3() = %v, want %v", got, tt.want)
			}
		})
	}
}
