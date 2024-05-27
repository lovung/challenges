package weekly398

import "testing"

func Test_waysToReachStair(t *testing.T) {
	tests := []struct {
		name string
		k    int
		want int
	}{
		{
			k:    10,
			want: 0,
		},
		{
			k:    31,
			want: 6,
		},
		{
			k:    2,
			want: 4,
		},
		{
			k:    4194285,
			want: 8855,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToReachStair2(tt.k); got != tt.want {
				t.Errorf("waysToReachStair2() = %v, want %v", got, tt.want)
			}
		})
		// t.Run(tt.name, func(t *testing.T) {
		// 	if got := waysToReachStair(tt.k); got != tt.want {
		// 		t.Errorf("waysToReachStair() = %v, want %v", got, tt.want)
		// 	}
		// })
	}
}
