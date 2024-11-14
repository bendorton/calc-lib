package calc

import (
	"fmt"
	"testing"
)

func TestAddition_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 7, want: 12},
		{a: 10, b: 11, want: 21},
		{a: 100, b: 100, want: 200},
		{a: -100, b: 100, want: 0},
		{a: 0, b: -3, want: -3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			this := &Addition{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtraction_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: 1, b: 1, want: 0},
		{a: 10, b: 7, want: 3},
		{a: 5, b: 7, want: -2},
		{a: 10, b: 11, want: -1},
		{a: 100, b: 100, want: 0},
		{a: -100, b: 100, want: -200},
		{a: 0, b: -3, want: 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			this := &Subtraction{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMultiplication_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: 1, b: 1, want: 1},
		{a: 5, b: 7, want: 35},
		{a: 10, b: 11, want: 110},
		{a: 100, b: 100, want: 10_000},
		{a: -100, b: 100, want: -10_000},
		{a: 0, b: -3, want: 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			this := &Multiplication{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 1, b: 1, want: 1},
		{a: 14, b: 7, want: 2},
		{a: 100, b: 10, want: 10},
		{a: 10, b: 3, want: 3},
		{a: -100, b: 100, want: -1},
		{a: 5, b: -2, want: -2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			this := &Division{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
