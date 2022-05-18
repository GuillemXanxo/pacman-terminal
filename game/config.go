package game

import (
	"encoding/json"
	"os"
)

// Configuration struct keeps the data unmarshalled from config.json
type Configuration struct {
    Player   string `json:"player"`
    Ghost    string `json:"ghost"`
    Wall     string `json:"wall"`
    Dot      string `json:"dot"`
    Pill     string `json:"pill"`
    Death    string `json:"death"`
    Space    string `json:"space"`
    UseEmoji bool   `json:"use_emoji"`
}

var Config Configuration

func LoadConfig(file string) error {
  f, err := os.Open(file)
  if err != nil {
    return err
  }
  defer f.Close()

  decoder := json.NewDecoder(f)
  err = decoder.Decode(&Config)
  if err != nil {
    return err
  }
  return nil
}
