package main

import (
	"fmt"
	"strings"
)

func Add(r Recipe, c Recipes) Recipes {
	c[strings.ToUpper(r.Title)] = r
	return c
}

func List(c Recipes) {
	fmt.Println("CookBook :")
	for key, value := range c {
		fmt.Printf("%s : %v", key, value)
		fmt.Println()
	}
}

func Search(keyWord string, c Recipes) Recipe {
	return c[strings.ToUpper(keyWord)]
}

func Delete(keyWord string, c Recipes) Recipes {
	delete(c, strings.ToUpper(keyWord))
	return c
}
