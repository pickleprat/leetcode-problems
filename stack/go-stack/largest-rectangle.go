package main 

type Heights struct {
    Height      int 
    Index       int 
}

func Max(a, b int) int {
    if a>b {
        return a
    } else {
        return b
    }
}

func largestRectangleArea(heights []int) int {
    stack := make([] Heights, 0); 
    maxArea := 0; 
    var index int; 
    for i := 0; i < len(heights); i++ {
        if len(stack) != 0 {
            index = i;  
            for len(stack) != 0 && stack[len(stack) - 1].Height >= heights[i] {
                lastHeights := stack[len(stack) - 1]; 
                maxArea = Max(maxArea, lastHeights.Height * (i - lastHeights.Index)); 
                stack = stack[:len(stack) - 1]; 
                index = lastHeights.Index; 
            }

            stack = append(stack, Heights{
                Height:     heights[i], 
                Index:      index, 
            })
        } else {
            stack = append(stack, Heights{
                Height:     heights[i], 
                Index :     i, 
            })
        }
    } 

    if len(stack) != 0 {
        i := len(heights)
        for len(stack) != 0 {
            lastHeights := stack[len(stack) - 1]; 
            maxArea = Max(maxArea, lastHeights.Height * (i - lastHeights.Index)); 
            stack = stack[:len(stack) - 1]; 
        }
    }

    return maxArea
}

