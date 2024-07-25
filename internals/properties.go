package matrix

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

//Type methods

func (m *Matrix[N]) InsertElement(Row, Column int, value N) error {
	//check if the bounds are valid
	if err := m.checkBounds(Row, Column); err != nil {
		return err
	}

	//now should check if the type of value N matches the type of matrix m
	//but i cant do that

	(*m)[Row][Column] = value
	return nil
}

func (m *Matrix[N]) DeleteElement(Row, Column int) error {
	err := m.InsertElement(Row, Column, 0)
	if err != nil {
		return err
	}
	return nil

}

// insert 2 for a random unit matrix
func (m *Matrix[N]) Random(maxvalue int) error {
	var wg sync.WaitGroup
	errch := make(chan error)
	var count int = 0

	for n, i := range *m {
		for j := range i {
			count++
			wg.Add(1)
			go func(n, j int) {
				defer wg.Done()
				defer close(errch)

				r := rand.New(rand.NewSource(time.Now().UnixNano()))

				seed := r.Intn(maxvalue)
				// seed negative number
				if math.Floor(float64(count)/4) == 0 { //
					seed = seed * -1
				}
				//now since our matrix could be any type
				//just inserting seed will not work so we need a type switch

				var v N
				switch (interface{})(v).(type) {
				case uint:
					v = interface{}(uint(math.Abs(float64(seed)))).(N)
				case int:
					v = interface{}(seed).(N)
				case int16:
					v = interface{}(int16(seed)).(N)
				case int64:
					v = interface{}(int64(seed)).(N)
				case float64:
					v = interface{}(float64(seed)).(N)
				case complex128:
					v = interface{}(complex(float64(seed), float64(r.Intn(maxvalue)))).(N)
				default:
					//this is impossible. Matrix type is clearly defined
					errch <- fmt.Errorf("Unsupported N")

				}

				(*m)[n][j] = v
			}(n, j)
		}

	}
	wg.Wait()
	return nil
}

//Given go type enforcement, i cant write code that would handle all data types i want to conver so
//im stuck with duplicating code

// Creates a Natural number martrix. Positive values only
func MatrixUnsigned(Row, Column uint) (*Matrix[uint], error) {
	//early exist if the matrix dimensions are invalid
	if Row == 0 || Column == 0 {
		return nil, fmt.Errorf("cannot create a matrix with zero rows or columns")
	}

	var mat = make(Matrix[uint], Column)
	for i := range mat {
		mat[i] = make([]uint, Row)
	}
	return &mat, nil
}

// Creates a int Point matrix
func MatrixInt(Row, Column uint) (*Matrix[int], error) {
	//early exist if the matrix dimensions are invalid
	if Row == 0 || Column == 0 {
		return nil, fmt.Errorf("cannot create a matrix with zero rows or columns")
	}

	var mat = make(Matrix[int], Column)
	for i := range mat {
		mat[i] = make([]int, Row)
	}
	return &mat, nil
}

// Creates a floating Point matrix
func Matrixfloat(Row, Column uint) (*Matrix[float64], error) {
	//early exist if the matrix dimensions are invalid
	if Row == 0 || Column == 0 {
		return nil, fmt.Errorf("cannot create a matrix with zero rows or columns")
	}

	var mat = make(Matrix[float64], Column)
	for i := range mat {
		mat[i] = make([]float64, Row)
	}
	return &mat, nil
}

// creates a complex number matrix
func MatrixComplex(Row, Column uint) (*Matrix[complex128], error) {
	//early exist if the matrix dimensions are invalid
	if Row == 0 || Column == 0 {
		return nil, fmt.Errorf("cannot create a matrix with zero rows or columns")
	}

	var mat = make(Matrix[complex128], Column)
	for i := range mat {
		mat[i] = make([]complex128, Row)
	}
	return &mat, nil

}

func (m *Matrix[N]) Print() {
	if m.getBounds().column > 20 || m.getBounds().row > 20 {
		fmt.Println("Matrix dimensions greater than 20 units. \n Please consider writing to file or getting slice instead")
	}
	for _, i := range *m {
		fmt.Println(i)
	}
}
