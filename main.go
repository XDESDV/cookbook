package main

import "fmt"

func main() {
	// var cookbook Recipes

	// var r Recipe
	// r.Title = "crepes"
	// r.Season = "toutes"
	// r.Description = "Oeufs, fraine, huile, lait, eau"
	// r.TimeToCook = 30
	// cookbook = append(cookbook, r)
	// cookbook = append(cookbook, Recipe{"Crepes", "toutes", "Oeufs, fraine, huile, lait, eau", 30})

	var cookbook = make(map[string]Recipe)
	cookbook = InitRecipes(cookbook)
	List(cookbook)
	r := Search("crEpEs", cookbook)
	fmt.Println("Search result : ", r)

	cookbook = Delete("crEpEs", cookbook)
	List(cookbook)
}

func InitRecipes(c Recipes) Recipes {
	Add(Recipe{"Crepes", "toutes", "Oeufs, fraine, huile, lait, eau", 30}, c)
	Add(Recipe{"Pates Carbo", "toutes", "pates, gruyeres, lardon, oeuf", 15}, c)
	return c
}
