package may2024

import (
	"fmt"
	"testing"
)

func Test_mincostToHireWorkers(t *testing.T) {
	type args struct {
		quality []int
		wage    []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				quality: []int{10, 20, 5},
				wage:    []int{70, 50, 30},
				k:       2,
			},
			want: 105.00000,
		},
		{
			args: args{
				quality: []int{3, 1, 10, 10, 1},
				wage:    []int{4, 8, 2, 2, 7},
				k:       3,
			},
			want: 30.66667,
		},
		{
			args: args{
				quality: []int{25, 68, 35, 62, 52, 57, 35, 83, 40, 51},
				wage:    []int{147, 97, 251, 129, 438, 443, 120, 366, 362, 343},
				k:       6,
			},
			want: 1979.31429,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mincostToHireWorkers1(tt.args.quality, tt.args.wage, tt.args.k)
			gotFmt := fmt.Sprintf("%.5f", got)
			if gotFmt != fmt.Sprintf("%.5f", tt.want) {
				t.Errorf("mincostToHireWorkers1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			got := mincostToHireWorkers2(tt.args.quality, tt.args.wage, tt.args.k)
			gotFmt := fmt.Sprintf("%.5f", got)
			if gotFmt != fmt.Sprintf("%.5f", tt.want) {
				t.Errorf("mincostToHireWorkers2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			got := mincostToHireWorkers3(tt.args.quality, tt.args.wage, tt.args.k)
			gotFmt := fmt.Sprintf("%.5f", got)
			if gotFmt != fmt.Sprintf("%.5f", tt.want) {
				t.Errorf("mincostToHireWorkers3() = %v, want %v", got, tt.want)
			}
		})
	}
}
