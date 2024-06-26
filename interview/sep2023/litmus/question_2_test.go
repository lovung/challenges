package litmus

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findAllHobbyists(t *testing.T) {
	type args struct {
		hobby   string
		hobbies map[string][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				hobby: "Yoga",
				hobbies: map[string][]string{
					"Steve": {"Fashion", "Piano", "Reading"},
					"Patty": {"Drama", "Magic", "Pets"},
					"Chad":  {"Puzzles", "Pets", "Yoga"},
				},
			},
			want: []string{"Chad"},
		},
		{
			args: args{
				hobby: "Pets",
				hobbies: map[string][]string{
					"Steve": {"Fashion", "Piano", "Reading"},
					"Patty": {"Drama", "Magic", "Pets"},
					"Chad":  {"Puzzles", "Pets", "Yoga"},
				},
			},
			want: []string{"Patty", "Chad"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAllHobbyists(tt.args.hobby, tt.args.hobbies)
			sort.Strings(got)
			sort.Strings(tt.want)
			assert.Equal(t, tt.want, got)
		})
	}
}
