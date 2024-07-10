// stringutils_test.go
package stringutils

import (
    "testing"
)

func TestReverse(t *testing.T) {
    input := "hello"
    expected := "olleh"
    result := Reverse(input)
    if result != expected {
        t.Errorf("Reverse(%q) == %q, want %q", input, result, expected)
    }
}

func TestToUpper(t *testing.T) {
    input := "hello"
    expected := "HELLO"
    result := ToUpper(input)
    if result != expected {
        t.Errorf("ToUpper(%q) == %q, want %q", input, result, expected)
    }
}

func TestToLower(t *testing.T) {
    input := "HELLO"
    expected := "hello"
    result := ToLower(input)
    if result != expected {
        t.Errorf("ToLower(%q) == %q, want %q", input, result, expected)
    }
}
