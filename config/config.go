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
	TimerConfig   = newTimerConfig().LoadConfig()
	GeneralConfig = newGeneralConfig()
)

type timerConfig struct {
	Work       time.Duration
	ShortBreak time.Duration
	LongBreak  time.Duration
}

func newTimerConfig() *timerConfig {
	return &timerConfig{
		Work:       time.Minute * 25,
		ShortBreak: time.Minute * 5,
		LongBreak:  time.Minute * 15,
	}
}

func (t timerConfig) Save() error {
	f, err := os.Create(GeneralConfig.ConfigFile)
	if err != nil {
		return err
	}
	defer f.Close()

	toJSON := struct {
		Work       string `json:"workTimer"`
		ShortBreak string `json:"shortBreak"`
		LongBreak  string `json:"longBreak"`
	}{
		t.Work.String(),
		t.ShortBreak.String(),
		t.LongBreak.String(),
	}

	err = json.NewEncoder(f).Encode(toJSON)
	if err != nil {
		return err
	}

	return nil
}

func (t *timerConfig) LoadConfig() *timerConfig {
	f, err := utils.GetConfigFile(GeneralConfig.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fs, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if fs.Size() == 0 {
		return t
	} else {
		err = json.NewDecoder(f).Decode(t)
		if err != nil {
			log.Fatal(err)
		}
	}
	return t
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
