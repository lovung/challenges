package medium

import "testing"

func Test_countGoodNumbers(t *testing.T) {
	tests := []struct {
		args int64
		want int
	}{
		{24, 971328007},
		{50, 564908303},
		{806166225460393, 643535977},
		{1, 5},
		{3, 100},
		{9, 800000},
		{10, 3200000},
		{11, 16000000},
		{12, 64000000},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := countGoodNumbers(tt.args); got != tt.want {
				t.Errorf("countGoodNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
