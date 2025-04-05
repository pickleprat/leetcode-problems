package main 

import "fmt"; 

// a way to execute an operation between two numbers 
func Execute(operand string, left int, right int) int {
    switch operand {
        case "+": 
            return left + right; 
        case "-": 
            return left - right; 
        case "/": 
            return left / right; 
        case "*": 
            return left * right; 
        default: 
            return 0; 
    }
}

func IsOperand(operand string ) bool {
    return operand == "+" || operand == "-" || operand == "*" || operand == "/"; 
}

func evalRPN(tokens []string) int {
    // a way to convert strings to integers between -200 and 200 
    convertToNumber := make(map[string] int, 0);
    stack := make([] int, 0);  
    for i := -200; i <= 200; i++ {
        convertToNumber[fmt.Sprintf("%d", i)] = i; 
    }

    for _, token := range tokens {
        if IsOperand(token) && len(stack) >= 2 {
            right := stack[len(stack) - 1]; 
            left := stack[len(stack) - 2]; 
            executedOutput := Execute(token, left, right);
            stack = stack[:len(stack) - 1]
            stack[len(stack) - 1] = executedOutput;  
        } else {
            stack = append(stack, convertToNumber[token]); 
        }
    }

    return stack[0]; 
}
