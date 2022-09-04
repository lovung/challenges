package contest

import "testing"

func Test_numberOfWays(t *testing.T) {
	type args struct {
		startPos int
		endPos   int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				startPos: 1,
				endPos:   2,
				k:        3,
			},
			want: 3,
		},
		{
			args: args{
				startPos: 1,
				endPos:   2,
				k:        5,
			},
			want: 10,
		},
		{
			args: args{
				startPos: 1,
				endPos:   5,
				k:        3,
			},
			want: 0,
		},
		{
			args: args{
				startPos: 1,
				endPos:   4,
				k:        3,
			},
			want: 1,
		},
		{
			args: args{
				startPos: 5,
				endPos:   1,
				k:        3,
			},
			want: 0,
		},
		{
			args: args{
				startPos: 4,
				endPos:   1,
				k:        3,
			},
			want: 1,
		},
		{
			args: args{
				startPos: 2,
				endPos:   5,
				k:        10,
			},
			want: 0,
		},
		{
			args: args{
				startPos: 264,
				endPos:   198,
				k:        68,
			},
			want: 68,
		},
		{
			args: args{
				startPos: 989,
				endPos:   1000,
				k:        99,
			},
			want: 934081896,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.args.startPos, tt.args.endPos, tt.args.k); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
