package matrix_test

import (
	"log"
	"os"
	"testing"
	"vectris/matrix"
)

func TestMain(m *testing.M) {
	setup()

	r := m.Run()

	teardown()
	os.Exit(r)
}

func TestMatrixUnsigned(t *testing.T) {
	dim := []uint{20, 20}

	t.Run("Normal input", func(t *testing.T) {
		_, err := matrix.MatrixUnsigned(dim[0], dim[1])
		if err != nil {
			log.Println(err)
			t.Fail()
		}
	})

	t.Run("Zero value as input", func(t *testing.T) {
		_, err := matrix.MatrixUnsigned(0, 0)
		if err != nil {
			log.Println("Error occured but this is the expected behaviour")
		} else {
			log.Println("Unexpected behaviour ")
			t.Fail()
		}
	})

}

func BenchmarkMatrixUnsigned(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matrix.MatrixUnsigned(40, 45)
	}
}

// intmatrix tests
func TestMatrixInt(t *testing.T) {
	dim := []uint{20, 20}

	t.Run("Unsigned input", func(t *testing.T) {
		v, err := matrix.MatrixInt(dim[0], dim[1])
		if err != nil {
			log.Println(err)
			t.Fail()
		}
		if v == nil {
			t.Fail()
		}

	})

	t.Run("Normal input", func(t *testing.T) {
		v, err := matrix.MatrixInt(40, 40)
		if err != nil {
			log.Println(err)
			t.Fail()
		}
		if v == nil {
			t.Fail()
		}
	})
	t.Run("Zero value as input", func(t *testing.T) {
		_, err := matrix.MatrixInt(0, 0)
		if err != nil {
			log.Println("Error occured but this is the expected behaviour")
		} else {
			log.Println("Unexpected behaviour ")
			t.Fail()
		}

	})

	t.Run("Float input", func(t *testing.T) {
		v, err := matrix.MatrixInt(0.0, 9.0)
		if v != nil {
			t.Errorf("Unexpected behaviour when test should fail")
		}
		if err != nil {
			log.Println("Error occured but this is the expected behaviour")
		} else {
			t.Errorf("Unexpected behaviour")
		}
	})

}

func BenchmarkMatrixInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matrix.MatrixInt(50, 50)
	}
}

// float matrix
func TestMatrixfloat(t *testing.T) {
	dim := []uint{20, 20}

	t.Run("Normal input", func(t *testing.T) {
		_, err := matrix.Matrixfloat(dim[0], dim[1])
		if err != nil {
			log.Println(err)
			t.Fail()
		}
	})

	t.Run("Zero value as input", func(t *testing.T) {
		_, err := matrix.Matrixfloat(0, 0)
		if err != nil {
			log.Println("Error occured but this is the expected behaviour")
		} else {
			log.Println("Unexpected behaviour ")
			t.Fail()
		}
	})

	t.Run("Float input", func(t *testing.T) {
		v, err := matrix.MatrixInt(3.00, 9.00)
		if v != nil {
			t.Errorf("Unexpected behaviour when test should fail")
		}
		if err != nil {
			log.Println("Error occured but this is the expected behaviour")
		} else {
			t.Errorf("Unexpected behaviour")
		}
	})

}

func BenchmarkMatrixfloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matrix.MatrixUnsigned(40, 45)
	}
}

// complex matrix
func TestMatrixComplex(t *testing.T) {
	dim := []uint{20, 20}

	t.Run("Normal input", func(t *testing.T) {
		_, err := matrix.MatrixComplex(dim[0], dim[1])
		if err != nil {
			log.Println(err)
			t.Fail()
		}
	})

	t.Run("Zero value as input", func(t *testing.T) {
		_, err := matrix.MatrixComplex(0, 0)
		if err != nil {
			log.Println("Error occured but this is the expected behaviour")
		} else {
			log.Println("Unexpected behaviour ")
			t.Fail()
		}
	})

	t.Run("Float input", func(t *testing.T) {
		v, err := matrix.MatrixComplex(3.00, 9.00)
		if v != nil {
			t.Errorf("Unexpected behaviour when test should fail")
		}
		if err != nil {
			log.Println("Error occured but this is the expected behaviour")
		} else {
			t.Errorf("Unexpected behaviour")
		}
	})

}

func BenchmarkMatrixComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matrix.MatrixComplex(40, 45)
	}
}
/*
func TestInsertElement(t *testing.T) {
	set := func(n interface{}) error {
		err := matrix.InsertElement(n)
		return err
	}

	t.Run("Generic input uint", func(t *testing.T) {
		list := []uint{0, 3}
		for _, i := range list {
			set(i)
		}
	})

}
	*/
func setup() {
	log.Println("Init Tests")
}

func teardown() {
	log.Println("Test Completed")
}
/*
func TestDeleteElement(t *testing.T) {
	type args struct {
		mat    *m
		Row    int
		Column int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteElement(tt.args.mat, tt.args.Row, tt.args.Column); (err != nil) != tt.wantErr {
				t.Errorf("DeleteElement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/