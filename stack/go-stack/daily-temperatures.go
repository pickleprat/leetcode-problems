package main 

type Temperature struct {
    Temp int 
    Day int 
}

func dailyTemperatures(temperatures []int) []int {
    index := 0; 
    stack := make([] Temperature, 0);
    answers := make([] int, len(temperatures)); 

    for index < len(temperatures) {
        if len(stack) != 0 {
            top := stack[len(stack) - 1]; 
            if top.Temp < temperatures[index] {
                stack = stack[:len(stack) - 1]
                answers[top.Day] = index - top.Day;  
            } else {
                stack = append(stack, Temperature{
                    Temp: temperatures[index], 
                    Day: index, 
                }); 
                index++; 
            }
        } else {
            stack = append(stack, Temperature{
                Temp: temperatures[index], 
                Day: index, 
            }); 
            index++; 
        }
    }

    return answers 
}

