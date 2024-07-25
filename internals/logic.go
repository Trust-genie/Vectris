package matrix

import (
	"context"
	"fmt"
	"math/cmplx"
	"sync"
)

// Checks if two matrices are equal
//
// can perform a quick check where it compares the dimensions and element values
// Much Much faster that reflect.DeepEqual for large matrices
// are equal e.g(int and uint or uint and uint64)
func Equal[N Numeric](a, b *Matrix[N]) bool {

	if a.getBounds().row != b.getBounds().row || a.getBounds().column != b.getBounds().column {
		//quick exit if the dimensions of the two matrices dont match
		return false
	}

	//this is a good place to do things concurrently
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var holder bool = true

	//next check if the types match
	//luckyly go helps us check this at compile time and raises a type mismatch
	//i dont have to worry, right?

	//check values between two matrices

	//undoubtedly this will be the bottle neck. i could try a fan in fan out
	//lets get technical. What is the best case scenerio?
	//	1. we create a goroutine for each row  that performs the check for that row and do that all at once
	//	2. immediately we find a  false value we halt the search regardless of context. for{}?
	//
	//this starts a goroutine for every row and performs a search
	var once sync.Once
	for k := 0; k < len(*a); k++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()

			for j := 0; j < a.getBounds().row; j++ {
				select {
				case <-ctx.Done():
					return
				default:
					if (*a)[k][j] != (*b)[k][j] {

						once.Do(func() {
							holder = false
							cancel()
						})
					}
				}

			}

		}(k)
	}

	wg.Wait()
	return holder
}

// This only works for matrices of the same type
func Add[N Numeric](a, b *Matrix[N]) (*Matrix[N], error) {
	//check bounds
	if a.getBounds().row != b.getBounds().row || a.getBounds().column != b.getBounds().column {
		//quick exit if the dimensions of the two matrices dont match
		return nil, fmt.Errorf("bound mismatch: \n \t Bounds for matrices do not match. Try Resize()")
	}

	// make an empty copy of a the matrix
	new := make(Matrix[N], len((*a)))
	for i := range new {
		new[i] = make([]N, len((*a)[i]))
	}

	//again this is a really really good place to do things concurrently
	//create a go routine for every row and perform the additions as necessary
	var wg sync.WaitGroup
	for k := 0; k < len(*a); k++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			for i := 0; i < len(*a); i++ {
				new[k][i] = (*a)[k][i] + (*b)[k][i] // now this is fine but what is there are a million rows
			}

		}(k)
	}

	wg.Wait()
	return &new, nil
}

// sub is easily a duplication of Function add with special caveat for uint Matrices

func (m *Matrix[N]) MultiplyScalar(args N) *Matrix[N] {
	//use full for scalar
	//no need for bound checks
	//no need for type checks as well? go enforces type checking?

	//i could hard code values for 0 and 1. zero returns a null matrix and i just return m

	//doing things concurrently
	var wg sync.WaitGroup

	for i := 0; i < len(*m); i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			for j := 0; j < len((*m)[0]); j++ {
				(*m)[k][j] = args * (*m)[k][j] // again what if there are a million rows
			}
		}(i)
	}

	wg.Wait()
	return m
}

/*

//multiply vector?

// Now the daddy that daddyed their daddy
func MultiplyMatrices[N Numeric](a, b *Matrix[N]) (*Matrix[N], error) {
	//check the bounds

	//1. get bounds
	DimA := a.getBounds()
	DimB := b.getBounds()
	//perform check on conformability
	if DimB.row != DimA.column {
		if DimB.column == DimA.row {
			return nil, fmt.Errorf("bound mismatch: \n \t matrices boundries forbid multiplication. Try switching arguments in function call")
		}

		return nil, fmt.Errorf("bound mismatch: \n \t matrices boundries forbid multiplication. Try Resize() ")
	}

	//2. verify type match
	//again go helps us enforce types between matrices so i dont have to worry about that

	//3. create new matrix
	new := make([][]N, DimA.row)
	for i := range new {
		new[i] = make([]N, DimB.column)
	}

	//4. begin multiplication
	var wg sync.WaitGroup
	for k := 0; k < len(*a); k++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			//magic

		}(k)
	}

	wg.Wait()
	return &new, nil
}

*/

func (m *Matrix[N]) ConvertToFloat() *[][]float64 {
	var wg sync.WaitGroup
	//get bounds
	Dim := m.getBounds()
	//create identical matrix
	new := make([][]float64, Dim.row)
	for i := range new {
		new[i] = make([]float64, Dim.column)
	}

	//copy values in old to new
	for i := 0; i < len(*m); i++ {
		wg.Add(1)
		go func(k int) {
			for j := 0; j < len((*m)[0]); j++ {
				new[k][j] = float64((*m)[k][j])
			}
		}(i)
	}

	wg.Wait()
	return &new

}

func (m *Matrix[N]) ConvertToComplex() *[][]complex128 {
	var wg sync.WaitGroup
	//get bounds
	Dim := m.getBounds()
	//create identical matrix
	new := make([][]complex128, Dim.row)
	for i := range new {
		new[i] = make([]complex128, Dim.column)
	}

	//copy values in old to new
	for i := 0; i < len(*m); i++ {
		wg.Add(1)
		go func(k int) {
			for j := 0; j < len((*m)[0]); j++ {
				switch v := (interface{})((*m)[k][j]).(type) {
				case float64:
					new[k][j] = complex(v, 0)
				}
			}
		}(i)
	}

	wg.Wait()
	return &new

}

func (m *Matrix[N]) Resize(New_Row, New_column uint) (*Matrix[N], error) {
	//1. check bounds
	if uint(m.getBounds().row) < New_Row || uint(m.getBounds().column) < New_column {
		//quick exit if the dimensions provided would shrink the matrix
		return nil, fmt.Errorf("Resize error: \n \t Cannot Shrink matrix, check values for new row and column provided")
	}

	new := make(Matrix[N], New_Row)
	for i := 0; i < len(new); i++ {
		new[i] = make([]N, New_column)
	}

	//add old values into new matrix
	var wg sync.WaitGroup
	for k := 0; k < len(*m); k++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			for i := 0; i < len(*m); i++ {
				new[k][i] = (*m)[k][i] // now this is fine but what is there are a million rows
			}

		}(k)
	}

	wg.Wait()
	return &new, nil

}

func (m *Matrix[complex128]) Conjugate() {
	//do some magic
	var wg sync.WaitGroup
	for i := 0; i < len(*m); i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			for j := 0; j < len((*m)[0]); j++ {
				(*m)[k][j] = interface{}(cmplx.Conj((*m)[k][j])).(complex128)
			}
		}(i)

	}
	wg.Wait()
}

func (m *Matrix[N]) Transpose() error {
	// we would the fliping the matrix on its axis

	return nil
}

func (m *Matrix[N]) Determinant() N {
	//this one is a bit tricky because it is really compulationally expensive

	return 0
}
