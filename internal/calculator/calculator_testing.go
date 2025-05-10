package calculator

import (
	"testing"
)

func TestCalculateSimple(t *testing.T) {
	result, err := Calculate("2+2")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 4 {
		t.Errorf("expected 4, got %v", result)
	}
}

func TestCalculatePriority(t *testing.T) {
	result, err := Calculate("2+2*2")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 6 {
		t.Errorf("expected 6, got %v", result)
	}
}

func TestCalculateWithBrackets(t *testing.T) {
	result, err := Calculate("(2+2)*2")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 8 {
		t.Errorf("expected 8, got %v", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Calculate("10/0")
	if err == nil {
		t.Error("expected error for division by zero, got none")
	}
}
