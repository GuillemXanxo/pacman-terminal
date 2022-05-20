package game

import (
	"fmt"
	"pacman/input"
	"time"
)

var NumDots int
var Lives = 3
var Score int

type functionMoveCursor func(int, int, Configuration)

func DeadCheck(player input.Sprite, ghosts []*input.Ghost, cfg Configuration, mazeInner []string, moveCursor functionMoveCursor) {
  for _, ghost := range ghosts {
    if player.Row == ghost.Position.Row && player.Col == ghost.Position.Col {
      Lives =- 1
      if Lives != 0 {
        moveCursor(player.Row, player.Col, cfg)
        fmt.Print(cfg.Death)
        moveCursor(len(mazeInner)+2, 0, cfg)
        time.Sleep(1000 * time.Millisecond)
        player.Row, player.Col = player.StartRow, player.StartCol
      }
    }
     
  }
}
