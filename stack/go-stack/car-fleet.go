package main 

import (
    "slices"
    "cmp"
)

type Car struct {
    Position    int 
    Speed       int
}

func CompareCars(a, b Car) int {
    return cmp.Compare(b.Position, a.Position); 
}


func carFleet(target int, position []int, speed []int) int {
    cars := [] Car {}; 
    fleetStack := make([] float32, 0); 
    for i := 0; i < len(position); i++ {
        cars = append(cars, Car{
            Position: position[i], 
            Speed   : speed[i], 
        });
    }

    slices.SortFunc(cars, CompareCars);

    for _, car := range cars {
        timeTaken := float32(target - car.Position) / float32(car.Speed); 
        fleetStack = append(fleetStack, timeTaken); 
        if len(fleetStack) >= 2 && fleetStack[len(fleetStack) - 1] <= fleetStack[len(fleetStack) - 2] {
            fleetStack = fleetStack[:len(fleetStack) - 1]; 
        }
    }

    return len(fleetStack); 
     
}