package mapreduce

// MapInt returns a new slice containing the result of applying operation to
// each element of input.
func MapInt(input []int, operation func(int) int) []int {
	output := make([]int, len(input))
	for i, v := range input {
		output[i] = operation(v)
	}
	return output
}

// FilterInt returns a new slice containing only the elements of input for which
// condition returns true.
func FilterInt(input []int, condition func(int) bool) []int {
	output := make([]int, 0)
	for _, v := range input {
		if condition(v) {
			output = append(output, v)
		}
	}
	return output
}

// ReduceInt combines the elements of input into a single value by repeatedly
// applying operation, starting from initial.
func ReduceInt(input []int, operation func(int, int) int, initial int) int {
	result := initial
	for _, v := range input {
		result = operation(result, v)
	}
	return result
}
