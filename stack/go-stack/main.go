package main

import (
	"fmt"
	"slices"
)


func main() {
	cars := [] Car {
		{ 	Position: 10, Speed: 2 }, 
		{	Position: 8, Speed: 4  }, 
		{ 	Position: 0, Speed: 1 }, 
		{ 	Position: 5, Speed: 1 }, 
		{	Position: 3, Speed: 3 }, 
	}; 

	slices.SortStableFunc(cars, CompareCar); 
	fmt.Printf("Cars sorted: %+v", cars); 
}