package compare

import (
	"day01/jsonDB"
	"day01/xmlDB"
	"fmt"
)

var (
	oldCakes = make([]string, 0)
	newCakes = make([]string, 0)
)

func arrContain(arr []string, s string) bool {
	for _, val := range arr {
		if val == s {
			return true
		}
	}
	return false
}

func CakesCompare() {
	for _, val := range newCakes {
		if !arrContain(oldCakes, val) {
			fmt.Printf("ADDED cake \"%s\"\n", val)
		}
	}
	for _, val := range oldCakes {
		if !arrContain(newCakes, val) {
			fmt.Printf("REMOVED cake \"%s\"\n", val)
		}
	}
}

func TimeCompare(oldRecipes *xmlDB.XMLRecipes, newRecipes *jsonDB.JSONRecipes) {
	for _, val1 := range oldRecipes.Cakes {
		for _, val2 := range newRecipes.Cakes {
			if val1.Name == val2.Name && val1.Time != val2.Time {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n",
					val1.Name, val2.Time, val1.Time)
			}
		}
	}
}

func IngredientCompare(oldRecipes *xmlDB.XMLRecipes, newRecipes *jsonDB.JSONRecipes) {
	oldF := func(oldIng []xmlDB.XMLIngredient) []string {
		res := make([]string, 0)
		for _, val := range oldIng {
			res = append(res, val.Name)
		}
		return res
	}

	newF := func(newIng []jsonDB.JSONIngredient) []string {
		res := make([]string, 0)
		for _, val := range newIng {
			res = append(res, val.Name)
		}
		return res
	}

	for _, val1 := range oldRecipes.Cakes {
		for _, val2 := range newRecipes.Cakes {
			if val1.Name == val2.Name {
				oldIng := oldF(val1.Ingredients.Ingredients)
				newIng := newF(val2.Ingredients)

				for _, val := range newIng {
					if !arrContain(oldIng, val) {
						fmt.Printf("ADDED ingredient \"%s\"  for cake \"%s\"\n", val, val1.Name)
					}
				}
				for _, val := range oldIng {
					if !arrContain(newIng, val) {
						fmt.Printf("REMOVED ingredient \"%s\"  for cake \"%s\"\n", val, val1.Name)
					}
				}
			}
		}
	}
}

func IngredientUnitAndCountCompare(oldRecipes *xmlDB.XMLRecipes, newRecipes *jsonDB.JSONRecipes) {
	for _, val1 := range oldRecipes.Cakes {
		for _, val2 := range newRecipes.Cakes {
			if val1.Name == val2.Name {
				for _, val3 := range val1.Ingredients.Ingredients {
					for _, val4 := range val2.Ingredients {
						if val3.Name == val4.Name && val3.Unit != val4.Unit && val3.Unit != "" && val4.Unit != "" {
							fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
								val3.Name, val1.Name, val4.Unit, val3.Unit)
						}
						if val3.Name == val4.Name && val3.Count != val4.Count {
							fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
								val3.Name, val1.Name, val4.Count, val3.Count)
						}
						if val3.Name == val4.Name && val3.Unit == "" && val4.Unit != "" {
							fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
								val4.Unit, val4.Name, val1.Name)
						}
						if val3.Name == val4.Name && val3.Unit != "" && val4.Unit == "" {
							fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
								val3.Unit, val3.Name, val1.Name)
						}
					}
				}
			}
		}
	}
}

func Compare(oldRecipes *xmlDB.XMLRecipes, newRecipes *jsonDB.JSONRecipes) {
	for _, cake := range oldRecipes.Cakes {
		oldCakes = append(oldCakes, cake.Name)
	}
	for _, cake := range newRecipes.Cakes {
		newCakes = append(newCakes, cake.Name)
	}

	CakesCompare()
	TimeCompare(oldRecipes, newRecipes)
	IngredientCompare(oldRecipes, newRecipes)
	IngredientUnitAndCountCompare(oldRecipes, newRecipes)
}
