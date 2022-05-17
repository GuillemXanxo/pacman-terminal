package maze

import (
	"bufio"
	/* bufio permite realizar operaciones de lectura y escritura sin tener que preocuparnos
	de dimensionar correctamente los slices de bytes a leer o escribir ni de tener que
	comprobar el num de bytes leidos o escritos y controlar el completado de las operaciones */
	"os"
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
  for scanner.Scan(){
    line := scanner.Text()
    mazeInner = append(mazeInner, line)
  }
//scanner.Scan() will return true while there is something to be read from the file, 
//and scanner.Text() will return the next line of input.
  
  return mazeInner, nil
//if everything goes ok return error == nil
}
