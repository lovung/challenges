package apr2024

import (
	"reflect"
	"testing"
)

func Test_deckRevealedIncreasing(t *testing.T) {
	tests := []struct {
		name string
		deck []int
		want []int
	}{
		// TODO: Add test cases.
		{
			deck: []int{17, 13, 11, 2, 3, 5, 7},
			want: []int{2, 13, 3, 11, 5, 17, 7},
		},
		{
			deck: []int{1, 1000},
			want: []int{1, 1000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deckRevealedIncreasing(tt.deck); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deckRevealedIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
