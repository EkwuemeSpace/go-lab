package main

import "testing"

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        input    string
        reversed string
        expected bool
    }{
        {"racecar", "racecar", true},
        {"hello", "olleh", false},
    }
    
    for _, tt := range tests {
        if isPalindrome(tt.input, tt.reversed) != tt.expected {
            t.Errorf("isPalindrome(%s, %s) = %v, want %v", 
                tt.input, tt.reversed, !tt.expected, tt.expected)
        }
    }
}