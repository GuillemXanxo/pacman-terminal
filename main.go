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

  //load configuration for emojis
  err = game.LoadConfig("config.json")
	if err != nil {
		log.Println("failed to load configuration:", err)
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
    PrintMaze(maze, game.Config)

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
    if game.Lives <= 0 {
      moveCursor(input.Player.Row, input.Player.Col, game.Config)
      fmt.Print(game.Config.Death)
      moveCursor(len(maze)+2, 0, game.Config)
      break
    }
    if game.NumDots == 0 {
      break
    }

    // repeat
    time.Sleep(200 * time.Millisecond)
  }
}


//each loop of the game the screen is cleared and then each line of the maze is printed again
func PrintMaze (str []string, cfg game.Configuration) {
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
  //Print player
  moveCursor(input.Player.Row, input.Player.Col, cfg)
  fmt.Print(cfg.Player)
  //Print ghosts
  for _, ghost := range input.Ghosts {
    moveCursor(ghost.Row, ghost.Col, cfg)
    fmt.Print(cfg.Ghost)
  }
  // Move cursor outside of maze drawing area, otherwise stays next to the P
  moveCursor(len(str)+1, 0, cfg)
  fmt.Println("Score:", game.Score, "\tLives:", game.Lives)
}

func moveCursor(row, col int, cfg game.Configuration) {
    if cfg.UseEmoji {
        simpleansi.MoveCursor(row, col*2)
    } else {
        simpleansi.MoveCursor(row, col)
    }
}
