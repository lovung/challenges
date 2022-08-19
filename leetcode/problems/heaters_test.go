package problems

import "testing"

func Test_findRadius(t *testing.T) {
	type args struct {
		houses  []int
		heaters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				houses:  []int{1, 2, 3},
				heaters: []int{2},
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				houses:  []int{1, 2, 3, 4},
				heaters: []int{1, 4},
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				houses:  []int{1, 5},
				heaters: []int{2},
			},
			want: 3,
		},
		// [282475249,622650073,984943658,144108930,470211272,101027544,457850878,458777923]
		// [823564440,115438165,784484492,74243042,114807987,137522503,441282327,16531729,823378840,143542612]
		{
			name: "example 4",
			args: args{
				houses:  []int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923},
				heaters: []int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612},
			},
			want: 161834419,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRadius(tt.args.houses, tt.args.heaters); got != tt.want {
				t.Errorf("findRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}
