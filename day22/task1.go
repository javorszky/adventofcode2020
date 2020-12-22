package day22

import "errors"

func task1() {
	_, _ = getInputs()
}

// slicePop takes in a slice, pops the first value off, and returns the value, and the slice without the first value.
// Returns error in case slice is nil or less than one element long.
func slicePop(in []int) (int, []int, error) {
	if in == nil {
		return 0, nil, errors.New("slicePop: passed in nil slice")
	}

	if len(in) < 1 {
		return 0, nil, errors.New("slicePop: passed in empty slice")
	}
	return in[0], in[1:], nil
}

// sliceAdd will take a slice, and a number of ints, and returns a slice where the additional elements are added to the
// end of the original slice.
func sliceAdd(in []int, elements ...int) []int {
	for _, el := range elements {
		in = append(in, el)
	}
	return in
}
