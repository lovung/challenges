package contest394

import "testing"

func Test_numberOfSpecialChars1(t *testing.T) {
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
			word: "abBCab",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSpecialChars1(tt.word); got != tt.want {
				t.Errorf("numberOfSpecialChars1() = %v, want %v", got, tt.want)
			}
		})
	}
}
