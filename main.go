package main

import (
	"fmt"
	"log"
	"pacman/input"
	"pacman/maze"

	"github.com/danicat/simpleansi"
)



func main() {
  //prepare the terminal
  input.InitTerminal()
  defer input.RestoreTerminal()

  //load maze
  maze, err := maze.LoadMaze("maze01.txt")
  if err != nil {
    log.Println("Error while loading maze")
    return
  }

//game loop
  for {
    //update screen
    PrintMaze(maze)

    //process input
    input, err := input.ReadFromTerminal()
    if err != nil {
      log.Println("Error reading input: ", err)
      break
    }

    // process movement

    // process collisions

    // check game over
    if input == "ESC" {
      break
    }

    // repeat
  }
}

//each loop of the game the screen is cleared and then each line of the maze is printed again
func PrintMaze (str []string) {
  simpleansi.ClearScreen()
  for _, line := range str {
    fmt.Println(line)
  }
}
