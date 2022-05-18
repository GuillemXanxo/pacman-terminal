package input

import "math/rand"

//create a slice of sprite structs to hold all ghosts
var Ghosts []*Sprite

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
    ghost.Row, ghost.Col = makeMove(ghost.Row, ghost.Col, direction, maze)
  }
}
