package main

import (
	"fmt"

	"github.com/Parvis123/pokedex-apis/pokedex-apis/pokemon"
)

func main() {
	pokemon, err := pokemon.GetPokemon("croagunk")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Pokemon:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Front Default Sprite:", pokemon.Sprites.Other.Home.FrontDefault)
	fmt.Println("Front Shiny Sprite:", pokemon.Sprites.Other.Home.FrontShiny)
	fmt.Println("Cry:", pokemon.Cry.Latest)


	for _, ability := range pokemon.Abilities {
		fmt.Println("Ability:", ability.Ability.Name)
	}

	for _, ability := range pokemon.Types {
		fmt.Println("Types:", ability.Type.Name)
	}
}