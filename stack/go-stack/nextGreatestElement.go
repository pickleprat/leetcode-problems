package main 

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    nextGreatestMap := make(map[int] int); 
    stack := make([] int, 0); 
    idx := 0; 

    stack = append(stack, nums2[0]); 

    for len(stack) != 0 && idx < len(nums2) {
        if nums2[idx] > stack[len(stack) - 1] {
            for len(stack) != 0 && nums2[idx] > stack[len(stack) - 1] {
                top := stack[len(stack) - 1];
                nextGreatestMap[top] = idx; 
                stack = stack[:len(stack) - 1];  
            } 
        } 
        stack = append(stack, nums2[idx]);
        idx++; 
    }

    output := make([] int, 0)
    for _, num := range nums1 {
        if ngn, ok := nextGreatestMap[num]; ok {
            output = append(output, nums2[ngn]); 
        } else {
            output = append(output, -1); 
        }
    }

    return output
}
