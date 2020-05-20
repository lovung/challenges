package maychallenge

import (
	"reflect"
	"testing"
)

func TestStack_push(t *testing.T) {
	type fields struct {
		size  int
		count int
		stack []interface{}
	}
	type args struct {
		n interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				size:  tt.fields.size,
				count: tt.fields.count,
				stack: tt.fields.stack,
			}
			if err := s.push(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Stack.push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStack_pop(t *testing.T) {
	type fields struct {
		size  int
		count int
		stack []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "test 1",
			fields: fields{
				size:  10,
				count: 0,
				stack: make([]interface{}, 10),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				size:  tt.fields.size,
				count: tt.fields.count,
				stack: tt.fields.stack,
			}
			if got := s.pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
