# Pokedex CLI

A fun **command-line Pokedex** written in **Go** that uses the **PokeAPI**.
Explore areas, encounter Pokémon, throw Pokéballs, and build your own Pokédex — all from your terminal.

---

# 🎮 Try It Yourself

Clone the project:

```bash
git clone https://github.com/YOUR_USERNAME/pokedexcli.git
cd pokedexcli
```

Build the program:

```bash
go build
```

Run it:

```bash
./pokedexcli
```

You will see:

```
Pokedex >
```

You're now inside the Pokédex CLI!

---

# 🕹 Example Interactive Session

```
Pokedex > help
Welcome to the Pokedex!

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
- psyduck
- golduck
- magikarp

Pokedex > catch psyduck
Throwing a Pokeball at psyduck...
psyduck was caught!

Pokedex > pokedex
Your Pokedex:
 - psyduck

Pokedex > inspect psyduck
Name: psyduck
Height: 8
Weight: 196
Stats:
 -hp: 50
 -attack: 52
 -defense: 48
 -special-attack: 65
 -special-defense: 50
 -speed: 55
Types:
 - water
```

---

# 📜 Commands

### Show Help

```
help
```

Displays all available commands.

---

### List Map Locations

```
map
```

Shows a page of location areas.

---

### Previous Map Page

```
mapb
```

Shows the previous page of locations.

---

### Explore an Area

```
explore <location-area>
```

Example:

```
explore pastoria-city-area
```

Displays Pokémon that can appear there.

---

### Catch a Pokémon

```
catch <pokemon>
```

Example:

```
catch pikachu
```

The higher the Pokémon's **base experience**, the harder it is to catch!

---

### Inspect a Pokémon

```
inspect <pokemon>
```

Shows detailed information about a Pokémon **if you have caught it**.

---

### View Your Pokédex

```
pokedex
```

Lists all Pokémon you have caught.

---

### Exit the Program

```
exit
```

Closes the CLI.

---

# 📁 Project Structure

```
pokedexcli
│
├── main.go
├── main_test.go
│
├── pokeapi
│   ├── api.go
│   ├── location_area.go
│   └── pokemon.go
│
├── pokecache
│   ├── cache.go
│   └── cache_test.go
│
├── go.mod
└── README.md
```

---

# 🌐 API

This project uses the free public API:

https://pokeapi.co

---

# ⚡ Features

✔ Interactive CLI
✔ Pokémon exploration by location
✔ Randomized catching system
✔ Inspect detailed Pokémon stats
✔ Local Pokédex storage
✔ API response caching

---

# 🧪 Built With

* Go
* PokeAPI
* Custom caching layer
* CLI REPL interface

---

# ⭐ If you like this project

Give it a ⭐ on GitHub!
