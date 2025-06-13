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

type toJSON struct {
	Work       string `json:"workTimer"`
	ShortBreak string `json:"shortBreak"`
	LongBreak  string `json:"longBreak"`
}

type timerConfig struct {
	Work       time.Duration `json:"workTimer"`
	ShortBreak time.Duration `json:"shortBreak"`
	LongBreak  time.Duration `json:"longBreak"`
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

	toJS := toJSON{
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

func (t *timerConfig) LoadConfig() *timerConfig {
	fromJS := new(toJSON)
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
		err = json.NewDecoder(f).Decode(fromJS)
		if err != nil {
			log.Fatal(err)
		}
	}

	t.Work, err = time.ParseDuration(fromJS.Work)
	if err != nil {
		log.Fatal(err)
	}
	t.ShortBreak, err = time.ParseDuration(fromJS.ShortBreak)
	if err != nil {
		log.Fatal(err)
	}
	t.LongBreak, err = time.ParseDuration(fromJS.LongBreak)
	if err != nil {
		log.Fatal(err)
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
