package main

type Recipe struct {
	Title       string
	Season      string
	Description string
	TimeToCook  uint8
}

// type Recipes []Recipe

type Recipes map[string]Recipe
