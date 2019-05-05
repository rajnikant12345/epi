package ch2

func exchange(i, j int, in []int) {
	if i != j {
		p := in[i]
		in[i] = in[j]
		in[j] = p
	}
}

/*DutchNationalFlag solves dnf problem.
question:
Write a program which takes array with thee types of element as input
and sort them and print the output.
*/
func DutchNationalFlag(in []int, pivot int) []int {
	small := 0
	large := len(in) - 1
	i := 0
	for i <= large {
		if in[i] < pivot {
			exchange(small, i, in)
			i++
			small++
		} else if in[i] == pivot {
			i++
		} else {
			exchange(large, i, in)
			large--
		}
	}
	return in
}
