package input

import (
	"log"
	"os"
	"os/exec"
)

func InitTerminal() {
  //To enable the cbreak mode we are going to call an external command that controls terminal behaviour, the stty command. 
  //We are also going to disable terminal echo so we don't polute the screen with the output of key presses.
  cbTerm := exec.Command("stty", "cbreak", "-echo")
  cbTerm.Stdin = os.Stdin

  err := cbTerm.Run()
  if err != nil {
    log.Fatalln("Error occurred while activating cbreak mode terminal: ", err)
  }
}

func RestoreTerminal() {
  //Same process as initTerm with flags reversed
  cookedTerm := exec.Command("stty", "-cbreak", "echo")
  cookedTerm.Stdin = os.Stdin

  err := cookedTerm.Run()
  if err != nil {
    log.Fatalln("Error occurred while activating cooked mode terminal: ", err)
  }
}

func ReadFromTerminal() (string, error) {
  buffer := make([]byte, 100)

  cnt, err := os.Stdin.Read(buffer)
  if err != nil {
    return "", err
  }
//if we are reading just one byte (cnt==1) and it is the ESC key (0x1b) then we return ESC
  if cnt == 1 && buffer[0] == 0x1b { 
    return "ESC", nil
  } else if cnt >=3 {
      if buffer[0] == 0x1b && buffer[1] == '[' {
        switch buffer[2] {
        case 'A':
          return "UP", nil
        case 'B':
          return "DOWN", nil
        case 'C':
          return "RIGHT", nil
        case 'D':
          return "LEFT", nil
        }
      }
    
  }
  return "", nil
}
