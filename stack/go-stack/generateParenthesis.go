package main 

func GeneratePermutationsRecursive(str string, left, total int, output [] string) [] string {
    open := total - left; 
    close := len(str) - open; 

    if left == 0 && open == close {
        output = append(output, str); 
        return output

    } else if left == 0 {
        for i := 0; i < open - close; i++ {
            str += ")"; 
        }
        output = append(output, str); 
        return output 
    } else if open == close {
        output = GeneratePermutationsRecursive(str + "(", left - 1, total, output); 
    } else {
        output = GeneratePermutationsRecursive(str + "(", left - 1, total, output ); 
        output = GeneratePermutationsRecursive(str + ")", left, total, output); 
    }

    return output; 
}


func generateParenthesisRecursive(n int) []string {
    output := make([] string, 0)
    output = GeneratePermutationsRecursive("(", n - 1, n, output); 
    return output
}

type ParenthesisValidator struct {
    Parenthesis  string 
    Left int 
}

func generateParenthesis(n int) []string {
    stack := make([] ParenthesisValidator, 0);
    output := make([] string, 0); 
    stack = append(stack, ParenthesisValidator{
        Parenthesis: "(", 
        Left: n - 1, 
    }); 

    for len(stack) != 0 {
        pv := stack[len(stack) - 1]; 
        str := pv.Parenthesis; 
        left := pv.Left; 
        open := n - left; close := len(str) - open; 
        stack = stack[:len(stack) - 1]; 
        
        if left == 0 && open == close {
            output = append(output, str); 
        } else if left == 0 {
            for i := 0; i < open - close; i++ {
                str += ")"; 
            }
            output = append(output, str); 
        } else if open == close {
            stack = append(stack, ParenthesisValidator{
                Parenthesis: str + "(", 
                Left : left - 1, 
            })
        } else {
            stack = append(stack, ParenthesisValidator{
                Parenthesis: str + "(", 
                Left: left - 1, 
            })

            stack = append(stack, ParenthesisValidator{
                Parenthesis: str + ")", 
                Left: left, 
            })
        }
    }

    return output 

}