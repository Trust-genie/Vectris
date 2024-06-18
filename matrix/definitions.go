package matrix

type Numeric interface {
	int | int16 | int64 | float64 | complex128 | uint
}

type Matrix[N Numeric] [][]N
