package problems

/*
	https://sites.google.com/site/prologsite/prolog-problems
	Gandalf- 2019
*/

import (
	"errors"
)

func LastElement(list []int) (int, error) {
	// one

	length := len(list)

	if length == 0 {
		return 0, errors.New("empty list")
	} else {
		return list[length-1], nil
	}
}

func SecondToLast(list []int) (int, error) {
	// two

	length := len(list)

	if length < 2 {
		return 0, errors.New("list too small")
	} else {
		return list[length-2], nil
	}
}

func KthElement(list []int, index int) (int, error) {
	// three

	if index < 0 || index > len(list)-1 {
		return 0, errors.New("index out of bounds")
	} else {
		return list[index], nil
	}
}

func Count(list []int) int {
	// four
	return len(list)
}

func Reverse(list []int) []int {
	// five

	length := len(list)
	result := make([]int, length)

	for i := range list {
		result[length-i-1] = list[i]
	}

	return result
}

func main() {
}
