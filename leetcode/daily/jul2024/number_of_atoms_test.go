package jul2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countOfAtoms(t *testing.T) {
	type args struct {
		formula string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				formula: "Mg(((H2O)2Na))4(F)(H2SO4)N",
			},
			want: "FH18MgNNa4O12S",
		},
		{
			args: args{
				formula: "Mg((H2O)2Na)4(F)(H2SO4)N",
			},
			want: "FH18MgNNa4O12S",
		},
		{
			args: args{
				formula: "K4(ON(SO3)2)2",
			},
			want: "K4N2O14S4",
		},
		{
			args: args{
				formula: "H2O",
			},
			want: "H2O",
		},
		{
			args: args{
				formula: "C6H12O6",
			},
			want: "C6H12O6",
		},
		{
			args: args{
				formula: "Mg(OH)2",
			},
			want: "H2MgO2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countOfAtoms(tt.args.formula)
			assert.Equal(t, tt.want, got)
		})
	}
}
