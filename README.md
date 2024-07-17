#  pokedexcli

This is a Go project implementing a REPL (read-evaluate-print loop) with the inspiration of Pokémon.

## Project Description:

This project implements a simple, text-based Pokedex application in Go. It simulates a basic Pokedex experience where users can explore locations, encounter Pokemon, and build their collection. It was implemented following the 17th course on [Boot.dev](https://www.boot.dev/tracks/backend)

## Installation:

### 1. Prerequisites:
- Go version 1.22 or later installed on your system. You can verify the version by running `go version` in your terminal.
- A code editor or IDE for writing and editing Go code.

### 2. Clone the Repository:
```Bash
git clone https://github.com/<your-username>/<project-name>.git
```

### 3. Navigate to the project directory:

```Bash
cd <project-name>
```
Use code with caution.

4. Build the project:

```Bash
go build -o <project-name>
```
Use code with caution.

This will create an executable file named <project-name> in the current directory.

## Usage:

It allows users to explore the Pokemon world through various commands:

- **help**: Displays a message explaining the available commands and their functionalities.
    ![image](https://github.com/user-attachments/assets/d53a85a0-22c9-440a-bcca-0f1e1c5b3b29)
- **exit**: Terminates the application and exits the Pokedex.
- **map**: Shows the names of 20 location areas in the Pokemon world. Each subsequent call displays the next 20 locations, providing a way to navigate through the list.
    ![image](https://github.com/user-attachments/assets/4f6c18a7-0b1e-4e4c-968d-708ce3df397e)
- **mapb**: Displays the previous 20 locations visited through the "map" command, allowing users to go back in their exploration history.
- **explore [location name]**: Takes the name of a specific location area as input and lists all the Pokemon found in that area.
- **catch [pokemon name]**: Attempts to "catch" a Pokemon by adding its name to the user's Pokedex.
    ![image](https://github.com/user-attachments/assets/5e90b3bb-e4cd-459e-b59e-f836b13c10da)

- **inspect [pokemon name]**: Takes the name of a Pokemon as input and displays detailed information such as name, height, weight, stats, and type(s).
- **pokedex**: Prints a list of all the Pokemon names currently stored in the user's Pokedex.

## Upcoming Features:

- [x] **Cycling Through Commands**: Support for the "up" arrow to cycle through previous commands. Being a developer, I found myself hitting that key but it was not cycling 😆. This has been added so happy cycling instead of typing the command again
- [ ] **Progress Saving**: Persist a user's Pokedex to their computer so they can save progress between sessions
- [ ] **More Test Coverage**: Simply put, it's just for me to practice more with writing tests in Go. It's really low priority-ish for me though for now
