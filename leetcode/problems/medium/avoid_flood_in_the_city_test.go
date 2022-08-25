package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_avoidFlood(t *testing.T) {
	type args struct {
		rains []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// Test cases:
		// [24526,23938,93949,0,61370,0,0,80971,0,0,0,55547,0,80482,74942,77698,54795,0,0,0,0,77698,0,0,54410,0,0,23938,0,0,0,0,0,0,75018,0,0,47564,0,24526,0,0,0,48639,0,68554,39045,75357,48639,6992,0,47564,54410,0,0,0,74942,6992,40783,82220,0,75018,0,61370,0,0,68554,0,0,82220,0,40783,55547,80971,0,93949,87360,0,87360,39045,0,80482,75357,0,0,54795]

		// Expected:
		// [-1,-1,-1,23938,-1,24526,61370,-1,80971,93949,1,-1,55547,-1,-1,-1,-1,77698,74942,80482,54795,-1,1,1,-1,54410,1,-1,1,1,1,1,1,1,-1,75018,1,-1,47564,-1,1,1,1,-1,48639,-1,-1,-1,-1,-1,6992,-1,-1,68554,39045,75357,-1,-1,-1,-1,82220,-1,40783,-1,1,1,-1,1,1,-1,1,-1,-1,-1,1,-1,-1,87360,-1,-1,1,-1,-1,1,1,-1]
		{
			name: "test1",
			args: args{
				rains: []int{1, 2, 3, 4},
			},
			want: []int{-1, -1, -1, -1},
		},
		{
			name: "test2",
			args: args{
				rains: []int{1, 2, 0, 0, 2, 1},
			},
			want: []int{-1, -1, 2, 1, -1, -1},
		},
		{
			name: "test3",
			args: args{
				rains: []int{1, 2, 0, 1, 2},
			},
			want: []int{},
		},
		{
			name: "test4",
			args: args{
				rains: []int{1, 2, 0, 2, 0, 1},
			},
			want: []int{-1, -1, 2, -1, 1, -1},
		},
		{
			name: "test5",
			args: args{
				rains: []int{1, 2, 0, 2, 0, 0, 1},
			},
			want: []int{-1, -1, 2, -1, 1, 1, -1},
		},
		{
			name: "test6",
			args: args{
				rains: []int{0, 1, 1},
			},
			want: []int{},
		},
		{
			name: "test7",
			args: args{
				rains: []int{1, 0, 2, 3, 0, 1, 2},
			},
			want: []int{-1, 1, -1, -1, 2, -1, -1},
		},
		{
			name: "test8",
			args: args{
				rains: []int{1, 0, 2, 0, 2, 1},
			},
			want: []int{-1, 1, -1, 2, -1, -1},
		},
		{
			name: "test9",
			args: args{
				rains: []int{1, 0, 2, 0, 3, 0, 2, 0, 0, 0, 1, 2, 3},
			},
			want: []int{-1, 1, -1, 2, -1, 3, -1, 2, 1, 1, -1, -1, -1},
		},
		{
			name: "test10",
			args: args{
				rains: []int{2, 3, 0, 0, 3, 1, 0, 1, 0, 2, 2},
			},
			want: []int{},
		},
		{
			name: "test11",
			args: args{
				rains: []int{24526, 23938, 93949, 0, 61370, 0, 0, 80971, 0, 0, 0, 55547, 0, 80482, 74942, 77698, 54795, 0, 0, 0, 0, 77698, 0, 0, 54410, 0, 0, 23938, 0, 0, 0, 0, 0, 0, 75018, 0, 0, 47564, 0, 24526, 0, 0, 0, 48639, 0, 68554, 39045, 75357, 48639, 6992, 0, 47564, 54410, 0, 0, 0, 74942, 6992, 40783, 82220, 0, 75018, 0, 61370, 0, 0, 68554, 0, 0, 82220, 0, 40783, 55547, 80971, 0, 93949, 87360, 0, 87360, 39045, 0, 80482, 75357, 0, 0, 54795},
			},
			want: []int{-1, -1, -1, 1, -1, 1, 1, -1, 1, 1, 1, -1, 1, -1, -1, -1, -1, 77698, 23938, 24526, 1, -1, 1, 1, -1, 1, 1, -1, 1, 1, 1, 1, 1, 1, -1, 75018, 61370, -1, 47564, -1, 54410, 74942, 1, -1, 48639, -1, -1, -1, -1, -1, 6992, -1, -1, 68554, 80482, 75357, -1, -1, -1, -1, 82220, -1, 40783, -1, 55547, 80971, -1, 93949, 39045, -1, 1, -1, -1, -1, 1, -1, -1, 87360, -1, -1, 1, -1, -1, 54795, 1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, avoidFlood(tt.args.rains))
		})
	}
}
