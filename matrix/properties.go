package matrix

import (
	"fmt"
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


func (mat *Matrix[N]) DeleteElement(Row, Column int) error {
	err := mat.InsertElement(Row, Column, 0)
	if err != nil {
		return err
	}
	return nil

}

/*

// insert 2 for a random unit matrix
func (mat *Matrix[N]) Random(maxvalue int) {
	var count int = 0

	for n, i := range *mat {
		count++
		for j := range i {
			seed := rand.Intn(maxvalue)
			// seed negative number
			if math.Floor(float64(count)/4) == 0 { //
				seed = seed * -1
			}
			//now since our matrix could be any type
			//just inserting seed will not work so we need a type switch

			var v N
			switch any(v).(type) {
			case uint:
				if seed >= 0 {
					seed = -1 * seed
				}
				v = N(uint(seed))
			case int:
				v = N(int(seed))
			case int16:
				v = N(int16(seed))
			case int64:
				v = N(int64(seed))
			case float64:
				v = N(float64(seed))
			case complex128:
				v = N(complex(float64(seed), 0))
			default:
				//this is impossible. Matrix type is clearly defined
				fmt.Errorf("Unsupported N")

			}

			mat.InsertElement(n, j, v)
		}
	}
}

*/
//Given go type enforcement, i cant write code that would handle all data types i want to conver so
//im stuck with duplicating code

// Creates a Natural number martrix. Positive values only
func MatrixUnsigned(Row, Column uint) (*[][]uint, error) {
	//early exist if the matrix dimensions are invalid
	if Row == 0 || Column == 0 {
		return nil, fmt.Errorf("cannot create a matrix with zero rows or columns")
	}

	var mat = make([][]uint, Column)
	for i := range mat {
		mat[i] = make([]uint, Row)
	}
	return &mat, nil
}

// Creates a int Point matrix
func MatrixInt(Row, Column uint) ([][]int, error) {
	//early exist if the matrix dimensions are invalid
	if Row == 0 || Column == 0 {
		return nil, fmt.Errorf("cannot create a matrix with zero rows or columns")
	}

	var mat = make([][]int, Column)
	for i := range mat {
		mat[i] = make([]int, Row)
	}
	return mat, nil
}

// Creates a floating Point matrix
func Matrixfloat(Row, Column uint) (*[][]float64, error) {
	//early exist if the matrix dimensions are invalid
	if Row == 0 || Column == 0 {
		return nil, fmt.Errorf("cannot create a matrix with zero rows or columns")
	}

	var mat = make([][]float64, Column)
	for i := range mat {
		mat[i] = make([]float64, Row)
	}
	return &mat, nil
}

// creates a complex number matrix
func MatrixComplex(Row, Column uint) (*[][]complex128, error) {
	//early exist if the matrix dimensions are invalid
	if Row == 0 || Column == 0 {
		return nil, fmt.Errorf("cannot create a matrix with zero rows or columns")
	}

	var mat = make([][]complex128, Column)
	for i := range mat {
		mat[i] = make([]complex128, Row)
	}
	return &mat, nil

}

func (m *Matrix[N]) Print() {
	for _, i := range *m {
		fmt.Println(i)
	}
}
