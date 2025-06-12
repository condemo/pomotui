package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"time"

	"github.com/condemo/pomotui/utils"
)

var (
	TimerConfig   = newTimerConfig()
	GeneralConfig = newGeneralConfig()
)

type timerConfig struct {
	Work       time.Duration
	ShortBreak time.Duration
	LongBreak  time.Duration
}

func newTimerConfig() *timerConfig {
	// TODO: Cargar los datos desde archivo local
	return &timerConfig{
		Work:       time.Minute * 1,
		ShortBreak: time.Minute * 1,
		LongBreak:  time.Minute * 1,
	}
}

func (t timerConfig) Save() error {
	f, err := os.Create(GeneralConfig.ConfigFile)
	if err != nil {
		return err
	}
	defer f.Close()

	toJS := struct {
		Work       string `json:"workTimer"`
		ShortBreak string `json:"shortBreak"`
		LongBreak  string `json:"longBreak"`
	}{
		t.Work.String(),
		t.ShortBreak.String(),
		t.LongBreak.String(),
	}

	err = json.NewEncoder(f).Encode(toJS)
	if err != nil {
		return err
	}

	return nil
}

type generalConfig struct {
	ConfigDir  string
	ConfigFile string
}

func newGeneralConfig() *generalConfig {
	hd, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configDir := path.Join(hd, ".config/pomotui")
	err = utils.CheckFolder(configDir)
	if err != nil {
		log.Fatal(err)
	}

	return &generalConfig{
		ConfigDir:  configDir,
		ConfigFile: path.Join(configDir, "config.json"),
	}
}
