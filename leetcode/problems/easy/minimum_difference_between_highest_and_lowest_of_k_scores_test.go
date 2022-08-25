package easy

import "testing"

func Test_minimumDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{90},
				k:    1,
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				nums: []int{9, 4, 1, 7},
				k:    2,
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				nums: []int{9, 4, 1, 7, 12},
				k:    3,
			},
			want: 5,
		},
		{
			name: "test4",
			args: args{
				nums: []int{41900, 69441, 94407, 37498, 20299, 10856, 36221, 2231, 54526,
					79072, 84309, 76765, 92282, 13401, 44698, 17586, 98455, 47895, 98889,
					65298, 32271, 23801, 83153, 12186, 7453, 79460, 67209, 54576, 87785,
					47738, 40750, 31265, 77990, 93502, 50364, 75098, 11712, 80013, 24193,
					35209, 56300, 85735, 3590, 24858, 6780, 50086, 87549, 7413, 90444, 12284,
					44970, 39274, 81201, 43353, 75808, 14508, 17389, 10313, 90055, 43102,
					18659, 20802, 70315, 48843, 12273, 78876, 36638, 17051, 20478,
				},
				k: 5,
			},
			want: 1428,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
