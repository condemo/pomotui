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
		Work:       time.Minute * 30,
		ShortBreak: time.Minute * 5,
		LongBreak:  time.Minute * 15,
	}
}

func (t timerConfig) Save() {
	// TODO: Implementar guardado local
}

var TimerConfig = newTimerConfig()
