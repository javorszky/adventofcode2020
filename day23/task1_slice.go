package day23

func task1slice() {
	getLabel(getInputs(), 100)
}

func getLabel(in []int, step int) string {
	// create the slice
	return ""
}

func createCircleSlice(in []int) []int {
	crko := make([]int, len(in)+1)
	prev := 0
	first := 0
	for _, i := range in {
		if prev == 0 {
			first = i
			prev = i
			continue
		}
		crko[prev] = i
		prev = i
	}
	crko[prev] = first

	return crko
}
