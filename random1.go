package main

import(
             "fmt"
	"math/rand"
 	"time"
)

func main() {

	// Since it doesn’t have a see initialization, it always prints same random numbers
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println("---")
	
	
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println("---")	
	
	fmt.Println("Random Int:", rand.Int())
	fmt.Println("Random Float:", rand.Float32())
	fmt.Println("Random Permutation of N ints:", rand.Perm(15))

	// A not so random function in the math/rand package
	// The NormFloat64 method returns a number based on Normal distrib
	// mean=0, stddev=1 -- if called enough and plotted will give bell curve
	fmt.Println("Normalized:", rand.NormFloat64())

}
