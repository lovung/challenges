package contest394

import "testing"

func Test_numberOfSpecialChars2(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{
		{
			word: "aaAbcBC",
			want: 3,
		},
		{
			word: "abc",
			want: 0,
		},
		{
			word: "AbBCab",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSpecialChars2(tt.word); got != tt.want {
				t.Errorf("numberOfSpecialChars2() = %v, want %v", got, tt.want)
			}
		})
	}
}
