package input

import (
	"math/rand"
	"time"
)

//create a slice of sprite structs to hold all ghosts
var Ghosts []*Ghost

//define status of ghosts according to player having swallowed the pill
type GhostStatus string

const (
  GhostStatusNormal GhostStatus = "Normal"
  GhostStatusBlue   GhostStatus = "Blue"
)

type Ghost struct {
  Position Sprite
  Status GhostStatus
}

//define function to define ghosts movements
func ghostDirection() string {
  direction := rand.Intn(4) // dona un int entre 0 inclos i n no inclos. en aquest cas entre 0 i 3
  movement := map[int]string {
    0: "UP",
    1: "DOWN",
    2: "LEFT",
    3: "RIGHT",
  }
  return movement[direction]
}

//define a function to move ghosts
func MoveGhosts(maze []string) {
  for _, ghost := range Ghosts {
    direction := ghostDirection()
    ghost.Position.Row, ghost.Position.Col = makeMove(ghost.Position.Row, ghost.Position.Col, direction, maze)
  }
}

//a function that changes ghost status from normal to blue while pillDurationSecs in comfiguration
func changeGhostStatus(pill time.Duration) {
  var pillTimer *time.Timer
  for _, ghost := range Ghosts {
    ghost.Status = GhostStatusBlue
  }
  pillTimer = time.NewTimer(time.Second * pill)
  <-pillTimer.C
  for _, ghost := range Ghosts {
    ghost.Status = GhostStatusNormal
  }
}
