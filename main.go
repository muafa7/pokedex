package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"math/rand"
	"github.com/rasqi7/pokedexcli/pokeapi"
)

const baseLocationURL = "https://pokeapi.co/api/v2/location-area"

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	Next     string
	Previous string
	Pokedex map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandMap(cfg *config, args []string) error {
	var url string
	if cfg.Next == ""{
		url = baseLocationURL
	} else {
		url = cfg.Next
	}

	resp, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		fmt.Println("Error fetching location areas:", err)
		return nil
	}

	for _, res := range resp.Results {
		fmt.Println(res.Name)
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil
}

func commandMapb(cfg *config, args []string) error {
	var url string
	if cfg.Previous == ""{
		fmt.Println("you're on the first page")
		return nil
	}

	url = cfg.Previous

	resp, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		fmt.Println("Error fetching location areas:", err)
		return nil
	}

	for _, res := range resp.Results {
		fmt.Println(res.Name)
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("No area provided")
		return nil
	}

	areaName := args[0]
	resp, err := pokeapi.GetLocationArea(areaName)
	if err != nil {
		fmt.Println("Error fetching location area:", err)
		return nil
	}

	fmt.Println("Exploring", areaName+"...")
	fmt.Println("Found Pokemon:")

	for _, res := range resp.PokemonEncounters {
		fmt.Println("- ",res.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("No pokemon provided")
		return nil
	}
	
	pokemonName := args[0]

	resp, err := pokeapi.GetPokemon(pokemonName)
	if err != nil {
		fmt.Println("Error fetching pokemon:", err)
		return nil
	}

	fmt.Println("Throwing a Pokeball at", pokemonName+"...")
	
	baseExperience := resp.BaseExperience
	chance := rand.Intn(baseExperience)
	
	if chance < 40 {
		cfg.Pokedex[resp.Name] = resp
		fmt.Println(pokemonName, " was caught!")
		} else {
		fmt.Println(pokemonName, " escaped!")
	}

	return nil
}

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("No pokemon provided")
		return nil
	}
	
	pokemonName := args[0]

	pokemon, exists := cfg.Pokedex[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Println(" -"+stat.Stat.Name+":", stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println(" -", t.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("you have not caught pokemon")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.Pokedex {
		fmt.Println("- ", pokemon.Name)
	}

	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Location area list")
	fmt.Println("mapb: Location area list")
	fmt.Println("explore: Explore location area")
	fmt.Println("catch: Catch pokemon")
	fmt.Println("inspect: Inspect caught pokemon stats")
	fmt.Println("pokedex: List of caught pokemon")
	return nil
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Pokedex guide",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Location area list",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Location area list",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect caught pokemon stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List of caught pokemon",
			callback:    commandPokedex,
		},
	}
	cfg := config{
		Pokedex: make(map[string]pokeapi.Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input:= scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		command, exist := commands[cleaned[0]]
		if exist {
			command.callback(&cfg, cleaned[1:])
		} else {
			fmt.Println("Unknown command")
		}
	}
}