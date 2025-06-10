package config

import "time"

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
	// TODO: Implementar guardado local
	return nil
}

var TimerConfig = newTimerConfig()
