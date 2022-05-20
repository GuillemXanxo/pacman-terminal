package game

import (
	"encoding/json"
	"os"
	"time"
)

// Configuration struct keeps the data unmarshalled from config.json
type Configuration struct {
    Player    string `json:"player"`
    Ghost     string `json:"ghost"`
    Wall      string `json:"wall"`
    Dot       string `json:"dot"`
    Pill      string `json:"pill"`
    Death     string `json:"death"`
    Space     string `json:"space"`
    UseEmoji  bool   `json:"use_emoji"`
    GhostBlue string `json:"ghost_blue"`
	  PillDurationSecs time.Duration `json:"pill_duration_secs"`
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
