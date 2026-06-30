package mapreduce

// MapInt returns a new slice containing the result of applying operation to
// each element of input.
func Map[T int|float64|string, U int|float64|string](input []T, operation func(T) U) []U {
	output := make([]U, len(input))
	for i, v := range input {
		output[i] = operation(v)
	}
	return output
}

// FilterInt returns a new slice containing only the elements of input for which
// condition returns true.
func Filter[T int|float64|string](input []T, condition func(T) bool) []T {
	output := make([]T, 0)
	for _, v := range input {
		if condition(v) {
			output = append(output, v)
		}
	}
	return output
}

// ReduceInt combines the elements of input into a single value by repeatedly
// applying operation, starting from initial.
func Reduce[T int|float64, U int|float64](input []T, operation func(T, U) T, initial T) U {
	result := initial
	for _, v := range input {
		result = operation(result, v)
	}
	return U(result)
}

