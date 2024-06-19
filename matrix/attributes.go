package matrix

import "fmt"

func (mat *Matrix[N]) Rotate() [][]N {
	//find mat dimensions
	row := len(*mat)
	column := len(mat.GetColumn(0))
	//make duplicate
	v := make([][]N, row)
	for i := range v {
		v[i] = make([]N, column)
	}

	//begin copying
	for n, i := range *mat {
		for m, value := range i {
			//here we are
			v[m][row-n-1] = value
		}
	}

	return v

}

func Copy[N Numeric](mat *Matrix[N]) *Matrix[N] {
	new := make(Matrix[N], len((*mat)))
	for i := range new {
		new[i] = make([]N, len((*mat)[i]))
	}

	//copy values in mat to new
	for i := 0; i <= len(*mat); i++ {
		for j := 0; j <= len((*mat)[0]); j++ {
			new[i][j] = (*mat)[i][j]
		}

	}

	return &new
}

func (m *Matrix[N]) checkBounds(Row, Column int) error {
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

func (m *Matrix[N]) getBounds() struct{ row, column int } {

	return struct {
		row    int
		column int
	}{len(*m), len((*m)[0])}
}
