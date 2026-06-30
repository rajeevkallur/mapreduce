package mapreduce_test

import (
	"fmt"
	"testing"
	"mrmodule/internal/mapreduce"
)

func TestIntToFloat(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	result := mapreduce.Map(input, func(x int) float64 {
		return float64(x)
	})

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], v)
		}
	}
}

func TestFloatToInt(t *testing.T) {
	input := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	expected := []int{1, 2, 3, 4, 5}

	result := mapreduce.Map(input, func(x float64) int {
		return int(x)
	})

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], v)
		}
	}
}
func TestFloatToString(t *testing.T) {
	input := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	expected := []string{"1.1", "2.2", "3.3", "4.4", "5.5"}

	result := mapreduce.Map(input, func(x float64) string {
		return fmt.Sprintf("%g", x)
	})

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], v)
		}
	}
}
