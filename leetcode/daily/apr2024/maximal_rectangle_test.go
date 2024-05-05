package apr2024

import "testing"

func Test_maximalRectangle(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]byte
		want   int
	}{
		{
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 6,
		},
		{
			matrix: [][]byte{
				{'1', '1', '1', '1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '1', '0', '0', '0'},
				{'0', '1', '1', '1', '1', '0', '0', '0'},
			},
			want: 21,
		},
		{
			matrix: [][]byte{
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
			},
			want: 20,
		},
		{
			matrix: [][]byte{
				{'0', '0', '0', '1', '0', '1', '1', '1'},
				{'0', '1', '1', '0', '0', '1', '0', '1'},
				{'1', '0', '1', '1', '1', '1', '0', '1'},
				{'0', '0', '0', '1', '0', '0', '0', '0'},
				{'0', '0', '1', '0', '0', '0', '1', '0'},
				{'1', '1', '1', '0', '0', '1', '1', '1'},
				{'1', '0', '0', '1', '1', '0', '0', '1'},
				{'0', '1', '0', '0', '1', '1', '0', '0'},
				{'1', '0', '0', '1', '0', '0', '0', '0'},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle1(tt.matrix); got != tt.want {
				t.Errorf("maximalRectangle1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle2(tt.matrix); got != tt.want {
				t.Errorf("maximalRectangle2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle3(tt.matrix); got != tt.want {
				t.Errorf("maximalRectangle3() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle4(tt.matrix); got != tt.want {
				t.Errorf("maximalRectangle4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMaximalRectangle(b *testing.B) {
	type mat [][]byte
	matrix := []mat{
		[][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		},
		[][]byte{
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
		},
		[][]byte{
			{'0', '0', '0', '1', '0', '1', '1', '1'},
			{'0', '1', '1', '0', '0', '1', '0', '1'},
			{'1', '0', '1', '1', '1', '1', '0', '1'},
			{'0', '0', '0', '1', '0', '0', '0', '0'},
			{'0', '0', '1', '0', '0', '0', '1', '0'},
			{'1', '1', '1', '0', '0', '1', '1', '1'},
			{'1', '0', '0', '1', '1', '0', '0', '1'},
			{'0', '1', '0', '0', '1', '1', '0', '0'},
			{'1', '0', '0', '1', '0', '0', '0', '0'},
		},
	}
	const l = 1000
	var bigMat mat = make([][]byte, l)
	for i := range l {
		bigMat[i] = make([]byte, l)
		for j := range l {
			bigMat[i][j] = '1'
		}
	}
	matrix = append(matrix, bigMat)
	// b.ResetTimer()
	// for range b.N {
	// 	b.Run("sol_1", func(b *testing.B) {
	// 		for i := range matrix {
	// 			got := maximalRectangle1(matrix[i])
	// 			_ = got
	// 		}
	// 	})
	// }
	b.ResetTimer()
	for range b.N {
		b.Run("sol_2", func(b *testing.B) {
			for i := range matrix {
				got := maximalRectangle2(matrix[i])
				_ = got
			}
		})
	}
	b.ResetTimer()
	for range b.N {
		b.Run("sol_3", func(b *testing.B) {
			for i := range matrix {
				got := maximalRectangle3(matrix[i])
				_ = got
			}
		})
	}
	b.ResetTimer()
	for range b.N {
		b.Run("sol_4", func(b *testing.B) {
			for i := range matrix {
				got := maximalRectangle4(matrix[i])
				_ = got
			}
		})
	}
}
