package main 

import "unicode"

func isPalindrome(s string) bool {
    start := 0; 
    end := len(s) - 1; 
    for start <= end {
        firstLetter := s[start]; 
        lastLetter := s[end];

        if !unicode.IsNumber(rune(firstLetter)) && !unicode.IsLetter(rune(firstLetter)) {
            start++; continue 
        } else {
            firstLetter = byte(unicode.ToLower(rune(firstLetter))); 
        }

        if !unicode.IsNumber(rune(lastLetter)) && !unicode.IsLetter(rune(lastLetter)) {
            end--; continue 
        } else {
            lastLetter = byte(unicode.ToLower(rune(lastLetter))); 
        }

        if firstLetter != lastLetter {
            return false 
        }

        start++; end--; 
    }
    return true; 
}
