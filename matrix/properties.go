package matrix

import (
	"fmt"
)

//Type methods?

func (m *Matrix[N]) InsertElement(Row, Column int, value N) error {
	//early exit for improper values
	if Row <= 0 || Column <= 0 {
		return fmt.Errorf("invalid input: Index cannot be less than or equal to zero")
	} //more boundary checks
	if Row > len(*m) || Column > len(m.GetColumn(0)) {
		return fmt.Errorf("invalid Bounds. Bounds out of range of matrix")
	}

	return nil
}

// Implement bound checks in all methods or func that calls this
func (m *Matrix[N]) GetColumn(ColumnIndex uint) (Column []N) {
	//check bounds

	//to keep this function clean i should inplement bounc checks in functions that call this
	for _, i := range *m {
		Column = append(Column, i[ColumnIndex])
	}
	return Column
}

/*
func DeleteElement[m Matrix](mat *m, Row, Column int) error {
	if Row <= 0 || Column <= 0 {
		return fmt.Errorf("Invalid input: Index cannot be less than or equal to zero")
	}

	mat[Row][Column] = 0
	return nil

}

func Random[m Matrix](mat *m, maxvalue int) {
	var count int = 0
	for n, i := range mat {
		count++
		for j := range i {
			seed := rand.Intn(maxvalue)
			// seed negative number
			if math.Floor(float64(count)/4) == 0 { //
				seed = seed * -1
			}

			mat[n][j] = seed
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

/*
func Print[m Matrix](args *m) {
	switch args.(type){
	case [][]int:
		for _, i := range args{
			fmt.Println(i)
		}
	case [][]uint:
		for _, i := range args{
			fmt.Println(i)
		}
	case [][]float64:
		for _, i := range args{
			fmt.Println(i)
		}
	case [][]complex128:
		for _, i := range args{
			fmt.Println(i)
		}
	}
}
*/
