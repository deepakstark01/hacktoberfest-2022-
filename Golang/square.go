package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	slices := getRandom2DArray(10, 100000)
	getSquareSequentially(slices...)
	getSquareConcurrently(slices...)
}

func getSquareSequentially(slices ...[]int) {
	startTime := time.Now()
	getSquareOfMultipleSlices(slices...)

	timeDiff := time.Since(startTime)

	fmt.Printf("time took by sequential program: %s \n", timeDiff.String())

}

func getSquareConcurrently(slices ...[]int) {
	wg := &sync.WaitGroup{}

	var result []int
	wg.Add(len(slices))

	startTime := time.Now()
	for _, slice := range slices {
		slice := slice
		go func() {
			defer wg.Done()
			for _, element := range slice {
				result = append(result, element*element)
			}
		}()
	}
	wg.Wait()
	timeDiff := time.Since(startTime)
	fmt.Printf("time took by concurrent program: %s \n", timeDiff.String())
}

// getSquareOfMultipleSlices returns square of multiple slices
func getSquareOfMultipleSlices(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		for _, element := range slice {
			result = append(result, element*element)
		}
	}
	return result
}

// getRandom2DArray returns d array of n*k
func getRandom2DArray(n, k int) [][]int {
	res := make([][]int, 0)
	for n > 0 {
		n = n - 1
		res = append(res, rand.Perm(k))
	}
	return res
}
