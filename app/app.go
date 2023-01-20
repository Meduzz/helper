package app

import (
	"github.com/tkanos/gonfig"
)

type (
	// App the interface we rely on
	App interface {
		Start() error
	}
)

// Initiate read the configFile (and ENVS), call Start then return any errors
func Initiate[T App](configFile string, app T) error {
	err := gonfig.GetConf(configFile, app)

	if err != nil {
		return err
	}

	err = app.Start()

	if err != nil {
		return err
	}

	return nil
}
