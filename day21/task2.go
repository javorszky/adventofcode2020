package day21

import (
	"fmt"
	"sort"
	"strings"
)

func task2() {
	canonicalList := canonicalFoodList(foodMapping(getInputs()))
	fmt.Printf("Day 21 task 2: the canonical food list is: %s\n", canonicalList)
}

// canonicalFoodList will arrange the ingredients alphabetically by their allergen and separate them by commas to
// produce the canonical dangerous ingredient list.
func canonicalFoodList(in map[string]string) string {
	sortHelper := make([]string, 0)
	ingredientListSlice := make([]string, 0)
	for allergen := range in {
		sortHelper = append(sortHelper, allergen)
	}
	sort.Strings(sortHelper)
	for _, allergen := range sortHelper {
		ingredientListSlice = append(ingredientListSlice, in[allergen])
	}
	return strings.Join(ingredientListSlice, ",")
}

func foodMapping(s []string) map[string]string {
	allergensAndFoods := createAllergenFoodMapping(s)

	reduced := make(map[string][]string, 0)
	for allergen, foods := range allergensAndFoods {
		reduced[allergen] = reduceElementsToCommon(foods...)
	}

	return sherlockFoodMap(reduced)
}

func removeElementFromSlice(s string, slice []string) []string {
	found := false
	foundAt := 0
	for idx, element := range slice {
		if s == element {
			foundAt = idx
			found = true
		}
	}

	if !found {
		return slice
	}

	return append(slice[:foundAt], slice[foundAt+1:]...)
}

func sherlockFoodMap(allergenFoodMap map[string][]string) map[string]string {
	out := make(map[string]string, 0)
	toDelete := ""
	for {
		found := false
		// let's iterate over the allergen food map to find allergens that only correspond to one food.
		for allergen, food := range allergenFoodMap {
			if len(food) == 1 {
				found = true
				out[allergen] = food[0]
				toDelete = food[0]
				break
			}
		}

		// if found, delete that one from the map
		if found {
			for allergen := range allergenFoodMap {
				allergenFoodMap[allergen] = removeElementFromSlice(toDelete, allergenFoodMap[allergen])
			}
			continue
		}
		break
	}

	return out
}
