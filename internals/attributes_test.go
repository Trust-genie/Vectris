package matrix_test

import (
	m "GoMatrix/internals"
	"reflect"
	"testing"
)

func TestCopy(t *testing.T) {
	type args struct {
		mat *m.Matrix[int64]
	}
	tests := []struct {
		name string
		args args
		want *m.Matrix[int64]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := m.Copy(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
