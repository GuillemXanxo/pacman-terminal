PACMAN with Go

This exercise is based in Pacman with Go and emojis written in [this project](https://github.com/danicat/pacgo) by @danicat.

Her project is all written in one main.go file. As my main interest was understanding modularization in Go, I have tried to create different packages for every game logic:

- maze: in this package I have left all the logic for printig the maze and loading the map.
- input: in this package there is the logic to read player input from terminal and then translating this input to the pacman movement and ghosts movement.
- game: this package contains all the meta to complete the game, which conditions should be accomplished to finish the game and the configuration.

To be done:
in package game there is the condition to check if the player should be dead. There is a not accepted dependency from input to game. This is probably due to building packages while not knowing all the project. Probably the deadCheck function should have been done in the input package.
