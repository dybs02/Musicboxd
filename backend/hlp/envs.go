package hlp

import (
	"log"
	"os"
	"strings"
)

var Envs = map[string]string{}

func LoadEnvKeys() {
	b, err := os.ReadFile(".env.sample")
	if err != nil {
		log.Fatalf("could not read .env.sample file: %v", err)
	}

	envs := strings.Split(string(b), "=\n")
	if envs[len(envs)-1] == "" {
		envs = envs[:len(envs)-1]
	}

	for _, env := range envs {
		val := os.Getenv(env)
		if val == "" {
			log.Fatalf("did not set env vraiable %v", env)
		}
		Envs[env] = val
	}
}
