package grind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_accountsMerge(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			args: args{
				accounts: [][]string{
					{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
					{"John", "johnsmith@mail.com", "john00@mail.com"},
					{"Mary", "mary@mail.com"},
					{"John", "johnnybravo@mail.com"},
				},
			},
			want: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
		},
		{
			args: args{
				accounts: [][]string{
					{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
					{"John", "johnsmith1@mail.com", "john00@mail.com"},
					{"Mary", "mary@mail.com"},
					{"John", "johnnybravo@mail.com"},
					{"John", "johnsmith1@mail.com", "john_newyork@mail.com"},
				},
			},
			want: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith1@mail.com", "johnsmith@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
		},
		{
			args: args{
				accounts: [][]string{
					{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
				},
			},
			want: [][]string{
				{"Ethan", "Ethan0@m.co", "Ethan3@m.co"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, accountsMerge(tt.args.accounts))
		})
	}
}
