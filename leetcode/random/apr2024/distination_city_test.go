package apr2024

import "testing"

func Test_destCity(t *testing.T) {
	tests := []struct {
		paths [][]string
		want  string
	}{
		{
			paths: [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			want:  "Sao Paulo",
		},
		{
			paths: [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			want:  "A",
		},
		{
			paths: [][]string{{"A", "Z"}},
			want:  "Z",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := destCity(tt.paths); got != tt.want {
				t.Errorf("destCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
