package main

import (
	"fmt"
	"log"
	"pacman/game"
	"pacman/input"
	"pacman/maze"
	"time"

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

  //process input async
  intro := make(chan string)
  go func(ch chan<- string) { //argument only accepts a chan to write, not to read
    for {
      intro, err := input.ReadFromTerminal()
      if err != nil {
        log.Println("Error reading input: ", err)
        ch<- "ESC"
      }
    ch<- intro
    }
  }(intro)

//game loop
  for {
    //update screen
    PrintMaze(maze)

    // process movement
    select {
    case move := <-intro:
      if move == "ESC" {
				game.Lives = 0
			}
      game.NumDots, game.Score = input.MovePlayer(move, maze, game.NumDots, game.Score)
    default:  
    }
    input.MoveGhosts(maze)

    // process collisions
    game.DeadCheck(input.Player, input.Ghosts)

    // check game over
    if game.Lives <= 0 || game.NumDots == 0 {
      break
    }

    // repeat
    time.Sleep(200 * time.Millisecond)
  }
}

//each loop of the game the screen is cleared and then each line of the maze is printed again
func PrintMaze (str []string) {
  simpleansi.ClearScreen()
  for _, line := range str {
    for _, character := range line {
      switch character {
      case '#':
        fallthrough //tant si troba # com . imprimira el caracter
      case '.':
        fmt.Printf("%c", character)
      default:
        fmt.Print(" ")
      }
    }
    fmt.Println()
  }
  //Print player
  simpleansi.MoveCursor(input.Player.Row, input.Player.Col)
  fmt.Print("P")
  //Print ghosts
  for _, ghost := range input.Ghosts {
    simpleansi.MoveCursor(ghost.Row, ghost.Col)
    fmt.Print("G")
  }
  // Move cursor outside of maze drawing area, otherwise stays next to the P
  simpleansi.MoveCursor(len(str)+1, 0)
  fmt.Println("Score:", game.Score, "\tLives:", game.Lives)
}
