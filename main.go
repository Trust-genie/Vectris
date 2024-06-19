package main

import (
	"fmt"
	"sync"
	"vectris/matrix"
)

func main() {
	var wg sync.WaitGroup
	//do something amazing
	mat1, _ := matrix.MatrixUnsigned(20, 20)
	mat2, _ := matrix.MatrixUnsigned(20, 20)

	var count uint = 1000
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 20; j++ {
			for i := 0; i < 20; i++ {
				mat1.InsertElement(j, i, count)
				count++
			}
		}
	}()

	var count2 uint = 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 20; j++ {
			for i := 0; i < 20; i++ {
				mat2.InsertElement(j, i, count2)
				count2++
			}
		}
	}()

	wg.Wait()
	new, _ := matrix.Add[uint](mat1, mat2)

	mat1.Print()
	fmt.Println()
	mat2.Print()
	fmt.Println()
	new.Print()

}
