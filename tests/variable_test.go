package gonstraint

import (
	"github.com/znacer/gonstraint/variable"
	"testing"
)

func TestNewVariable(t *testing.T) {
	// Create a new variable
	variable, err := NewVariable("x", []int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the variable's properties
	if variable.Name != "x" {
		t.Errorf("Expected variable name to be 'x', got %s", variable.Name)
	}
	if len(variable.Domain) != 5 {
		t.Errorf("Expected domain to have 5 elements, got %d", len(variable.Domain))
	}
	if variable.Value != nil {
		t.Errorf("Expected variable value to be nil, got %v", variable.Value)
	}

	// Test the String() method
	expectedString := "x ∈ [1, 2, 3, 4, 5]"
	if variable.String() != expectedString {
		t.Errorf("Expected String() to return '%s', got '%s'", expectedString, variable.String())
	}
}

func TestNewVariable_InvalidDomain(t *testing.T) {
	// Try creating a variable with an empty domain
	_, err := NewVariable("x", []int{})
	if err == nil {
		t.Errorf("Expected error for empty domain, got nil")
	}

	// Try creating a variable with a nil domain
	_, err = NewVariable("x", nil)
	if err == nil {
		t.Errorf("Expected error for nil domain, got nil")
	}
}
