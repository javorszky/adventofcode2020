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
