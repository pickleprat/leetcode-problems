package main


func main() {
	cars := [] Car {
		{ 	Position: 10, Speed: 2 }, 
		{	Position: 8, Speed: 4  }, 
		{ 	Position: 0, Speed: 1 }, 
		{ 	Position: 5, Speed: 1 }, 
		{	Position: 3, Speed: 3 }, 
	}; 

	positions := make([] int, 0); 
	speeds := make([] int, 0); 
	for _, car := range cars {
		positions = append(positions, car.Position); 
		speeds = append(speeds, car.Speed); 
	} 

	carFleet(12, positions, speeds); 
}