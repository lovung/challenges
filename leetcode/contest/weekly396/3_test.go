package weekly396

import "testing"

func Test_minAnagramLength(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			s:    "abba",
			want: 2,
		},
		{
			s:    "cdef",
			want: 4,
		},
		{
			s:    "aaaaaaa",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAnagramLength(tt.s); got != tt.want {
				t.Errorf("minAnagramLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
