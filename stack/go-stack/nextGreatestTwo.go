package main 

type StackElement struct {
    Number int 
    Index int 
}


func nextGreaterElements(nums []int) []int {
    ngm := make(map[string] int); 
    stack := make([] StackElement, 0); 
    idx := 0 
    stack = append(stack, StackElement{
        Number: nums[0], 
        Index: 0, 
    }); 

    for len(stack) != 0 && idx < len(nums) {
        num := nums[idx]; 
        if num > stack[len(stack) - 1].Number {
            for len(stack) != 0 && num > stack[len(stack) - 1].Number {
                top := stack[len(stack) - 1]; 
                ngm[fmt.Sprintf("%d-%d", top.Index, top.Number)] = idx;
                stack = stack[:len(stack)-1];  
            }
        }

        stack = append(stack, StackElement{
            Index: idx, 
            Number: nums[idx], 
        }); 
        idx++; 
    }

    if len(stack) != 0 {
        idx = 0; 
        for idx < len(nums) {
            num := nums[idx]; 
            if num > stack[len(stack) - 1].Number {
                for len(stack) != 0 && num > stack[len(stack) - 1].Number {
                    top := stack[len(stack) - 1]; 
                    ngm[fmt.Sprintf("%d-%d", top.Index, top.Number)] = idx;
                    stack = stack[:len(stack)-1];  
                }
            }
            stack = append(stack, StackElement{
                Index: idx, 
                Number: nums[idx], 
            }); 
            idx++; 
        }
    }

    output := make([] int, 0); 
    for numIdx, num := range nums {
        if idx, ok := ngm[fmt.Sprintf("%d-%d", numIdx, num)]; ok {
            output = append(output, nums[idx]); 
        } else {
            output = append(output, -1); 
        }
    }

    return output 

}
