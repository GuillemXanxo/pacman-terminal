package input

//A struct to hold players position
type Sprite struct {
  Row int
  Col int
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

func MovePlayer(direction string, maze []string, numDots, score int) (int, int) {
  Player.Row, Player.Col = makeMove(Player.Row, Player.Col, direction, maze)
  switch maze[Player.Row][Player.Col] {
  case '.':
    numDots--
    score++
    //Remove dot from maze
    maze[Player.Row] = maze[Player.Row][0:Player.Col] + " " + maze[Player.Row][Player.Col+1:]
    //string are inmutable, we cannot just change the . for a blank
    //We generate a new string from copying the original until the position, blank, after the position to the end
  }
  return numDots, score
}
