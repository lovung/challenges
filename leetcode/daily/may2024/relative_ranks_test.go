package may2024

import (
	"reflect"
	"testing"
)

func Test_findRelativeRanks(t *testing.T) {
	tests := []struct {
		name  string
		score []int
		want  []string
	}{
		{
			score: []int{5, 4, 3, 2, 1},
			want:  []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			score: []int{10, 3, 8, 9, 4},
			want:  []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
		{
			score: []int{2, 1},
			want:  []string{"Gold Medal", "Silver Medal"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRelativeRanks(tt.score); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRelativeRanks() = %v, want %v", got, tt.want)
			}
		})
	}
}
