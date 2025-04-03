package main

import "fmt"


func main() {
	s1 := "{[{()}]}"; 
	s2 := "{{[{()}]}"; 

	if isValid(s1) && !isValid(s2)  {
		fmt.Printf("The strings %s and %s  were correctly evaluated.\n", s1, s2); 
	} 
}