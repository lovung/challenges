package biweekly130

import "testing"

func Test_minimumSubstringsInPartition(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			s:    "fabccddg",
			want: 3,
		},
		{
			s:    "abababaccddb",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSubstringsInPartition(tt.s); got != tt.want {
				t.Errorf("minimumSubstringsInPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
