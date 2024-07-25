package matrix_test

import (
	"GoMatrix/internals"
	"log"
	"reflect"
	"testing"
)

func TestMatrixUnsigned(t *testing.T) {
	type args struct {
		Row    uint
		Column uint
		Iserr  bool
	}
	tests := []struct {
		name    string
		args    args
		want    *[][]uint
		wantErr bool
	}{
		// TODO: Add test cases.
		struct {
			name    string
			args    args
			want    *[][]uint
			wantErr bool
		}{},
		struct {
			name    string
			args    args
			want    *[][]uint
			wantErr bool
		}{},
		struct {
			name    string
			args    args
			want    *[][]uint
			wantErr bool
		}{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := matrix.MatrixUnsigned(tt.args.Row, tt.args.Column)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatrixUnsigned() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixUnsigned() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMatrixUnsigned(b *testing.B) {
	test := []struct {
		name   string
		Row    uint
		Column uint
		Iserr  bool
	}{
		{"Small Dimensions", 10, 10, false},
		{"Medium Dimensions", 50, 50, false},
		{"Large Dimensions", 2000, 2000, false},
		{"XXL Dimensions", 15000, 15000, false},
	}

	for _, tt := range test {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := matrix.MatrixUnsigned(tt.Row, tt.Column)
				if err != nil {
					if tt.Iserr {
						log.Printf("Error occured but that is the expected behaviour %v", err)
					} else {
						b.Errorf("Unexpected Error occured %v", err)
					}
				}
			}
		})
	}
}

func TestMatrixInt(t *testing.T) {
	type args struct {
		Row    uint
		Column uint
		Iserr  bool
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := matrix.MatrixInt(tt.args.Row, tt.args.Column)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatrixInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkMatrixInt(b *testing.B) {
	type args struct {
		Row    uint
		Column uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Bench vs python", args{10000, 10000}, false},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := matrix.MatrixInt(tt.args.Row, tt.args.Column)

				if err != nil {
					b.Errorf("Error occured")
				}
			}
		})
	}

}

func TestMatrixfloat(t *testing.T) {
	type args struct {
		Row    uint
		Column uint
		Iserr  bool
	}
	tests := []struct {
		name    string
		args    args
		want    *[][]float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := matrix.Matrixfloat(tt.args.Row, tt.args.Column)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrixfloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrixfloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixComplex(t *testing.T) {
	type args struct {
		Row    uint
		Column uint
		Iserr  bool
	}
	tests := []struct {
		name    string
		args    args
		want    *[][]complex128
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := matrix.MatrixComplex(tt.args.Row, tt.args.Column)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatrixComplex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixComplex() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestInsertElement(t *testing.T) {
	type bounds struct{
		Row	uint
		Column uint
	}

	tests :=  []struct{
		name string
		mat matrix.Matrix[matrix.Numeric]
		bounds
		value matrix.Numeric
		wanterr bool
	}{
		struct{name string; mat matrix.Matrix[matrix.Numeric]; bounds; wanterr bool}{},

	}

	for _,ct := range tests{
		t.Run(ct.name, func(t *testing.T){
			got, err := mat.InsertElement(ct.bounds.Row, ct.bounds.Column)
			if err != nil{
				t.Errorf("Insert Error: got %v", err )
			}
		})
	}



}
*/
