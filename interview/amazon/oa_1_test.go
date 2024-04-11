package amazon

import "testing"

func Test_reduceGifts(t *testing.T) {
	type args struct {
		prices    []int32
		k         int32
		threshold int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{
				prices:    []int32{9, 6, 7, 2, 7, 2},
				k:         2,
				threshold: 13,
			},
			want: 2,
		},
		{
			args: args{
				prices:    []int32{9, 6, 3, 2, 9, 10, 10, 11},
				k:         1,
				threshold: 1,
			},
			want: 8,
		},
		{
			args: args{
				prices:    []int32{9, 6, 3, 2, 9, 10, 10, 11},
				k:         1,
				threshold: 9,
			},
			want: 3,
		},
		{
			args: args{
				prices:    []int32{9, 6, 3, 2, 9, 10, 10, 11},
				k:         4,
				threshold: 1,
			},
			want: 5,
		},
		{
			args: args{
				prices:    []int32{3, 2, 1, 4, 6, 5},
				k:         3,
				threshold: 14,
			},
			want: 1,
		},
		{
			args: args{
				prices:    []int32{3, 2, 1, 4, 6, 5},
				k:         3,
				threshold: 15,
			},
			want: 0,
		},
		{
			args: args{
				prices:    []int32{3, 2, 1, 4, 6, 5},
				k:         1,
				threshold: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reduceGifts(tt.args.prices, tt.args.k, tt.args.threshold); got != tt.want {
				t.Errorf("reduceGifts() = %v, want %v", got, tt.want)
			}
		})
	}
}
