package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex...")
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("It's empty!, Go Catch Some!")
		return nil
	}
	for _, p := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
