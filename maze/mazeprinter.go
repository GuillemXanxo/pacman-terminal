package maze

import (
	"fmt"

	"github.com/danicat/simpleansi"
)

//iterate over each entry in the Maze slice and print it
func PrintMaze (str []string) {
  simpleansi.ClearScreen()
  for _, line := range str {
    fmt.Println(line)
  }
}
