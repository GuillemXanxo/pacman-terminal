package game

import (
	"pacman/input"
)
var NumDots int
var Lives = 1
var Score int

func DeadCheck(player interface{}, ghosts []*input.Sprite) {
  for _, ghost := range ghosts {
    if player == *ghost { //need to dereference ghost
      Lives = 0
    }
  }
}
