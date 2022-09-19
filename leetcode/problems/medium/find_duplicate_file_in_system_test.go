package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDuplicate(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			args: args{
				paths: []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
			},
			want: [][]string{{"root/a/1.txt", "root/c/3.txt"}, {"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}},
		},
		{
			args: args{
				paths: []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"},
			},
			want: [][]string{{"root/a/1.txt", "root/c/3.txt"}, {"root/a/2.txt", "root/c/d/4.txt"}},
		},
		{
			args: args{
				paths: []string{"root/a 1.txt(abcd) 2.txt(efsfgh)", "root/c 3.txt(abdfcd)", "root/c/d 4.txt(efggdfh)"},
			},
			want: [][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicate(tt.args.paths)
			for i := range tt.want {
				assert.Contains(t, got, tt.want[i])
			}
		})
	}
}
