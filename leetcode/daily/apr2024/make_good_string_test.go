package apr2024

import "testing"

func Test_makeGood(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			s:    "leEeetcode",
			want: "leetcode",
		},
		{
			s:    "abBAcC",
			want: "",
		},
		{
			s:    "s",
			want: "s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGood(tt.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
