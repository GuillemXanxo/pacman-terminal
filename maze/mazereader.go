package maze

import (
	"bufio"
	/* bufio permite realizar operaciones de lectura y escritura sin tener que preocuparnos
	de dimensionar correctamente los slices de bytes a leer o escribir ni de tener que
	comprobar el num de bytes leidos o escritos y controlar el completado de las operaciones */
	"os"
	"pacman/input"
)



func LoadMaze(mazefile string) ([]string, error) {
  mazeInner := []string{}
//1- Open and close the file
//open the file with the maze
  file, err := os.Open(mazefile)
//check for errors in opening
  if err != nil {
    return nil, err
  }
//defer the close of the file
  defer file.Close()

//2- Save the string in the file to our var maze
//
  scanner := bufio.NewScanner(file)
//save each line to a var and append that line to our maze slice. Scan read each line of a file with \n end of lines
//scanner.Scan() will return true while there is something to be read from the file, 
//and scanner.Text() will return the next line of input.
  for scanner.Scan(){
    line := scanner.Text()
    mazeInner = append(mazeInner, line)
  }

//3- Capture player position
// traverse each character of the maze and create a new player when it locates a `P`
  for row, line := range mazeInner {
    for col, char := range line {
      switch char {
      case 'P':
        input.Player = input.Sprite{Row: row, Col: col}
      case 'G':
        input.Ghosts = append(input.Ghosts, &input.Sprite{Row: row, Col: col})
      }
    }
  }
  
  return mazeInner, nil
//if everything goes ok return error == nil
}
