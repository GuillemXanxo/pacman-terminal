package input

import "time"

//A struct to hold players position
type Sprite struct {
  Row int
  Col int
  StartRow int
  StartCol int
}

var Player Sprite

func makeMove (oldRow, oldCol int, direction string, maze []string) (newRow, newCol int) {
  newRow, newCol = oldRow, oldCol

	switch direction {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(maze) {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}
	}
//if the movement hit a wall is cancelled
  if maze[newRow][newCol] == '#' {
    newRow = oldRow
    newCol = oldCol
  }

  return
}

func MovePlayer(direction string, maze []string, numDots, score int, pill time.Duration) (int, int) {
  Player.Row, Player.Col = makeMove(Player.Row, Player.Col, direction, maze)
  switch maze[Player.Row][Player.Col] {
  case '.':
    numDots--
    score++
    removeDot(Player.Row, Player.Col, maze)
  case 'X':
    score += 10
    removeDot(Player.Row, Player.Col, maze)
    go changeGhostStatus(pill)
  }
  return numDots, score
}

//We generate a new string from copying the original until the position, blank, after the position to the end
func removeDot(row, col int, maze []string) { 
  maze[row] = maze[row][0:col] + " " + maze[row][col+1:]
}
