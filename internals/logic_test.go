package matrix_test

import (
	m "GoMatrix/internals"
	"log"
	"reflect"
	"testing"
)

func TestEqual(t *testing.T) {
	type args struct {
		a *m.Matrix[int]
		b *m.Matrix[int]
	}
	tests := []struct {
		name    string
		args    args
		wanterr bool
	}{
		// TODO: Add test cases.
		{"Small Matrix, Same type", args{Matrixhelper(10, 10), Matrixhelper(10, 6)}, false},
		{"Small Matrix , same type, zero", args{Matrixhelper(10, 1), Matrixhelper(3, 16)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := m.Equal(tt.args.a, tt.args.b); got != tt.wanterr {
				t.Errorf("Equal() = %v, want %v", got, tt.wanterr)
			}
		})
	}
}

func BenchmarkEqual(b *testing.B) {

	tests := []struct {
		name  string
		a     *m.Matrix[int]
		b     *m.Matrix[int]
		value bool
	}{
		{"Small Matrix same type", Matrixhelper(10, 10), Matrixhelper(10, 10), true},
		{"Medium Matrix same type", Matrixhelper(1000, 1000), Matrixhelper(1000, 1000), true},
		{"Large Matrix same type", Matrixhelper(10000, 2500), Matrixhelper(10000, 2500), true},
		{"Small Matrix different dim", Matrixhelper(10, 10), Matrixhelper(10, 3), false},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if m.Equal[int](tt.a, tt.b) != tt.value {
					b.Errorf("Wanted False, got true")
				}
			}
		})
	}

}

func BenchmarkDeepEqual(b *testing.B) {
	tests := []struct {
		name  string
		a     *m.Matrix[int]
		b     *m.Matrix[int]
		value bool
	}{
		{"Small Matrix same type", Matrixhelper(10, 10), Matrixhelper(10, 10), true},
		{"Medium Matrix same type", Matrixhelper(1000, 1000), Matrixhelper(1000, 1000), true},
		{"Large Matrix same type", Matrixhelper(10000, 2500), Matrixhelper(10000, 2500), true},
		{"Small Matrix different dim", Matrixhelper(10, 10), Matrixhelper(10, 3), false},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				d := reflect.DeepEqual(tt.a, tt.b)
				if d != tt.value {
					b.Errorf("wanted %v got %v", tt.value, d)

				}
			}
		})

	}

}

func TestAdd(t *testing.T) {
	type args struct {
		a *m.Matrix[int]
		b *m.Matrix[int]
	}
	/*tests := []struct {
		name string
		args args
		//want    *m.Matrix[int]
		wantErr bool
	}{
		//test boundary verification
		{"Same Dimensions, small size", args{Matrixhelper(100, 50), Matrixhelper(100, 50)}, false},
		{"Same Dimensions, Different size", args{Matrixhelper(100, 50), Matrixhelper(100, 50)}, true},
		{"Large Dimensions, small size", args{Matrixhelper(10000, 8000), Matrixhelper(10000, 8000)}, false},
	}
	log.Panicln("Testing Boundary verification")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.Add(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
	*/

	//test values
	log.Println("Testing Value")
	test2 := []struct {
		name    string
		args    args
		wantErr bool
		want    *m.Matrix[int]
	}{
		{"Small Dimensions", args{MatrixPopulate(100, 50), MatrixPopulate(100, 50)}, false, (MatrixPopulate(100, 50)).MultiplyScalar(2)},
		{"Large Dimensions", args{MatrixPopulate(1100, 750), MatrixPopulate(1100, 750)}, false, (MatrixPopulate(1100, 750)).MultiplyScalar(2)},
	}

	for _, tt := range test2 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.Add(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}

}

/*
	func TestMultiplyMatrices(t *testing.T) {
		type args struct {
			a *Matrix[N]
			b *Matrix[N]
		}
		tests := []struct {
			name    string
			args    args
			want    *Matrix[N]
			wantErr bool
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := MultiplyMatrices(tt.args.a, tt.args.b)
				if (err != nil) != tt.wantErr {
					t.Errorf("MultiplyMatrices() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("MultiplyMatrices() = %v, want %v", got, tt.want)
				}
			})
		}
	}
*/
func Matrixhelper(r, c uint) *m.Matrix[int] {
	mat, err := m.MatrixInt(r, c)
	if err != nil {
		return nil
	}
	return mat

}

func MatrixPopulate(r, c uint) *m.Matrix[int] {
	mat, err := m.MatrixInt(r, c)
	if err != nil {
		return nil
	}

	for i := 0; i < len(*mat); i++ {
		for j := 0; j < len((*mat)[0]); j++ {
			(*mat)[i][j] = int(i * 60 / 3)
		}

	}
	return mat

}
