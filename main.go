package main

import (
	"flag"
	"fmt"
	"log"
	"pacman/game"
	"pacman/input"
	"pacman/maze"
	"time"

	"github.com/danicat/simpleansi"
)

var (
  configFile = flag.String("config-file", "config.json", "path to custom configuration file")
  mazeFile   = flag.String("maze-file", "maze01.txt", "path to a custom maze file")
  /* The String function of the flag package accepts three parameters: 
  a flag name, a default value and a description (to be exhibited when --help is used). 
  It returns a pointer to a string which will hold the value of the flag. */
)

func main() {
  flag.Parse()

  //prepare the terminal
  input.InitTerminal()
  defer input.RestoreTerminal()

  //load maze
  maze, err := maze.LoadMaze(*mazeFile)
  if err != nil {
    log.Println("Error while loading maze")
    return
  }

  //load configuration for emojis
  err = game.LoadConfig(*configFile)
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
      game.NumDots, game.Score = input.MovePlayer(move, maze, game.NumDots, game.Score, game.Config.PillDurationSecs)
    default:  
    }
    input.MoveGhosts(maze)

    // process collisions
    game.DeadCheck(input.Player, input.Ghosts, game.Config, maze, MoveCursor)

    // check game over
    if game.Lives <= 0 {
      MoveCursor(input.Player.Row, input.Player.Col, game.Config)
      fmt.Print(game.Config.Death)
      MoveCursor(len(maze)+2, 0, game.Config)
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
  MoveCursor(input.Player.Row, input.Player.Col, cfg)
  fmt.Print(cfg.Player)
  //Print ghosts
  for _, ghost := range input.Ghosts {
    MoveCursor(ghost.Position.Row, ghost.Position.Col, cfg)
    if ghost.Status == input.GhostStatusNormal {
      fmt.Print(cfg.Ghost)
    } else if ghost.Status == input.GhostStatusBlue {
      fmt.Print(cfg.GhostBlue)
    }
  }
  // Move cursor outside of maze drawing area, otherwise stays next to the P
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
