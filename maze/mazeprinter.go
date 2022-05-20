package maze

import (
	"fmt"
	"pacman/game"
	"pacman/input"

	"github.com/danicat/simpleansi"
)

//iterate over each entry in the Maze slice and print it
func PrintMaze (str []string, cfg game.Configuration ) {
  simpleansi.ClearScreen()
  for _, line := range str {
    for _, character := range line {
      switch character {
      case '#':
        fmt.Print(simpleansi.WithBlueBackground(cfg.Wall))
      case '.':
        fmt.Print(cfg.Dot)
      default:
        fmt.Print(cfg.Space)
      }
    }
    fmt.Println()
  }

  MoveCursor(input.Player.Row, input.Player.Col, cfg)
  fmt.Print(cfg.Player)
  for _, ghost := range input.Ghosts {
    MoveCursor(ghost.Position.Row, ghost.Position.Col, cfg)
    fmt.Print(cfg.Ghost)
  }
  // Move cursor outside of maze drawing area
  MoveCursor(len(str)+1, 0, cfg)
  fmt.Println("Score:", game.Score, "\tLives:", game.Lives)
}

func MoveCursor(row, col int, cfg game.Configuration) {
    if cfg.UseEmoji {
        simpleansi.MoveCursor(row, col*2)
    } else {
        simpleansi.MoveCursor(row, col)
    }
}
