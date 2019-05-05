package ch2

//DeleteDuplicate - delete duplicate from sorted array
func DeleteDuplicate(in []int) []int {
	writeindex := 1
	for i := 1; i < len(in); i++ {
		if in[writeindex-1] != in[i] {
			in[writeindex] = in[i]
			writeindex++
		}
	}
	return in[:writeindex]
}
