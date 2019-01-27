package problems

/*
	https://sites.google.com/site/prologsite/prolog-problems
	Gandalf- 2019
*/

func sequence(length int) []int {

	result := make([]int, length)

	for i := 0; i < length; i++ {
		result[i] = i
	}

	return result
}

func equal(first, second []int) bool {

	if len(first) != len(second) {
		return false
	}

	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}
