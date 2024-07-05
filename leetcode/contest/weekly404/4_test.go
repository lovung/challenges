package weekly404

import "testing"

func Test_minimumDiameterAfterMerge(t *testing.T) {
	type args struct {
		edges1 [][]int
		edges2 [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				edges1: [][]int{{0, 1}, {2, 0}, {3, 2}, {3, 6}, {8, 7}, {4, 8}, {5, 4}, {3, 5}, {3, 9}},
				edges2: [][]int{{0, 1}, {0, 2}, {0, 3}},
			},
			want: 7,
		},
		{
			args: args{
				edges1: [][]int{{1, 0}, {2, 3}, {1, 4}, {2, 1}, {2, 5}},
				edges2: [][]int{{4, 5}, {2, 6}, {3, 2}, {4, 7}, {3, 4}, {0, 3}, {1, 0}, {1, 8}},
			},
			want: 6,
		},
		{
			args: args{
				edges1: [][]int{{0, 1}, {0, 2}, {0, 3}},
				edges2: [][]int{{0, 1}},
			},
			want: 3,
		},
		{
			args: args{
				edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}},
				edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}},
			},
			want: 5,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDiameterAfterMerge(tt.args.edges1, tt.args.edges2); got != tt.want {
				t.Errorf("minimumDiameterAfterMerge() = %v, want %v", got, tt.want)
			}
		})
		// Ignore wrong testcase
		if i > 0 {
			t.Run(tt.name, func(t *testing.T) {
				if got := minimumDiameterAfterMerge_WA(tt.args.edges1, tt.args.edges2); got != tt.want {
					t.Errorf("minimumDiameterAfterMerge_WA() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
