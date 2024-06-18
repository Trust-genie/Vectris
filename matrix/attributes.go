package matrix

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
