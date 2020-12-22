package day21

import (
	"fmt"
	"strings"
)

func task1() {
	num := findHowManyTimesNoAllergenFoodsAppear(getInputs())

	fmt.Printf("\nDay 21 task 1: the foods that can't have allergens in them appear %d times in total\n", num)
}

func findHowManyTimesNoAllergenFoodsAppear(s []string) int {
	allergensAndFoods := createAllergenFoodMapping(s)
	foodsAppear := createFoodFrequencyMap(s)

	allTheFoodsHelper := make([][]string, 0)
	reduced := make(map[string][]string, 0)
	for allergen, foods := range allergensAndFoods {
		allTheFoodsHelper = append(allTheFoodsHelper, foods...)
		reduced[allergen] = reduceElementsToCommon(foods...)
	}

	haveAllergenshelper := make([][]string, 0)
	for _, foods := range reduced {
		haveAllergenshelper = append(haveAllergenshelper, foods)
	}

	// these are the foods that can have allergens in them.
	haveAllergens := expandElementsToCover(haveAllergenshelper...)

	// all the foods
	allTheFoods := expandElementsToCover(allTheFoodsHelper...)

	// all the foods that do not have allergens in them.
	noAllergens := differenceElements(allTheFoods, haveAllergens)

	sum := 0
	for _, food := range noAllergens {
		sum = sum + foodsAppear[food]
	}

	return sum

}

func createAllergenFoodMapping(list []string) map[string][][]string {
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

// unionElements returns a string slice which contains elements from both of the slices once.
func unionElements(a, b []string) []string {
	n := make(map[string]struct{}, 0)
	out := make([]string, 0)
	for _, v := range a {
		n[v] = struct{}{}
	}

	for _, v := range b {
		n[v] = struct{}{}
	}

	for v := range n {
		out = append(out, v)
	}
	return out
}

// expandElementsToCover will take arbitrary number of slices, and returns a slice that contains all elements from all
// slices.
func expandElementsToCover(slices ...[]string) []string {
	union := slices[0]
	for _, list := range slices {
		union = unionElements(union, list)
	}
	return union
}

// differenceElements returns a slice of elements that are present in the first argument, but are not present in the
// second argument.
func differenceElements(presentInThis, butNotThis []string) []string {
	n := make(map[string]struct{}, 0)
	out := make([]string, 0)
	for _, v := range presentInThis {
		n[v] = struct{}{}
	}

	for _, v := range butNotThis {
		if _, ok := n[v]; ok {
			delete(n, v)
		}
	}

	for v := range n {
		out = append(out, v)
	}

	return out
}

func createFoodFrequencyMap(s []string) map[string]int {
	out := make(map[string]int, 0)
	for _, line := range s {
		line = strings.TrimRight(line, ")")
		parts := strings.Split(line, " (contains ")
		foods := strings.Split(parts[0], " ")
		for _, food := range foods {
			out[food]++
		}
	}
	return out
}
