package matrix

import (
	"context"
	"sync"
)

// Checks if two matrices are equal
//
// can perform a quick check where it compares the dimensions and element values
// can also try to perform a deep check where i use reflection to ensure that element types
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

			for j := 0; j <= a.getBounds().row; j++ {
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
func Add[N Numeric](a, b *Matrix[N]) *Matrix[N] {
	//check bounds
	if a.getBounds().row != b.getBounds().row || a.getBounds().column != b.getBounds().column {
		//quick exit if the dimensions of the two matrices dont match
		return nil
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
				new[k][i] = (*a)[k][i] + (*b)[k][i]
			}

		}(k)
	}

	wg.Wait()
	return &new
}
