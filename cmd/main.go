package main

import (
	matrix "GoMatrix/internals"
	"log"
)

func main() {

	//do something amazing
	mat, err := matrix.MatrixUnsigned(200, 300)
	if err != nil {
		log.Println(err)
	}

	mat.Print()

}
