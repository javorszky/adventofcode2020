package day21

import "strings"

func task1() {
	_ = getInputs()
}

func parseInputs(list []string) map[string][][]string {
	orderedList := make(map[string][][]string, 0)
	for _, line := range list {
		line = strings.TrimRight(line, ")")
		parts := strings.Split(line, " (contains ")
		foods := strings.Split(parts[0], " ")
		allergens := strings.Split(parts[1], ", ")
		for _, allergen := range allergens {
			orderedList[allergen] = append(orderedList[allergen], foods)
		}
	}
	return orderedList
}

// commonElements takes two string slices, and returns a third slice that has the common elements from both of them.
func commonElements(a, b []string) []string {
	n := make([]string, 0)
	helper := make(map[string]struct{}, 0)
	for _, v := range a {
		helper[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := helper[v]; ok {
			n = append(n, v)
		}
	}

	return n
}

// reduceElementsToCommon will take a slice of slices, and only return a slice of elements that are present in all of
// the passed slices.
func reduceElementsToCommon(slices ...[]string) []string {
	toCheck := slices[0]
	for _, sl := range slices {
		toCheck = commonElements(toCheck, sl)
	}
	return toCheck
}
