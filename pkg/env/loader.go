package env

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var (
	once sync.Once
	env  = Environment{}
)

const envPrefix = "intern"

func Load() (Environment, error) {
	var err error
	once.Do(func() {
		if err = envconfig.Process(envPrefix, &env); err != nil {
			err = fmt.Errorf("failed to process env: %w", err)
			return
		}
	})

	return env, err
}
