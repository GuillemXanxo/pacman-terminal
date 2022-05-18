package maze

import (
	"fmt"
	"pacman/game"
	"pacman/input"

	"github.com/danicat/simpleansi"
)

//iterate over each entry in the Maze slice and print it
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

  simpleansi.MoveCursor(input.Player.Row, input.Player.Col)
  fmt.Print("P")
  for _, ghost := range input.Ghosts {
    simpleansi.MoveCursor(ghost.Row, ghost.Col)
    fmt.Print("G")
  }
  // Move cursor outside of maze drawing area
  simpleansi.MoveCursor(len(str)+1, 0)
  fmt.Println("Score:", game.Score, "\tLives:", game.Lives)
}
