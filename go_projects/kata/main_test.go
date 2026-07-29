package main

import "testing"

func TestSpin(t *testing.T){
	tests:= []struct{
		name string
		str string
		expected string
	} {
		{"Normal", "I can enjoy Working", "I can yojne gnikroW"},
		{"ZeroValue", "", ""},
		{"Palindrome", "ololo", "ololo"},
		{"Max", "AAAAAAAAAAAAAAAA", "AAAAAAAAAAAAAAAA" },
		{"Min", "a", "a"},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			result:= SpinWords(tt.str)
			if result != tt.expected{
				t.Errorf("SpinWords(%s) = %s, Хотели  %s", 
			 tt.str, result, tt.expected)
			}
		})
	}
}