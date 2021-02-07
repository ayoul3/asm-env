package main

import (
	"os"
	"strings"

	"github.com/ayoul3/asm-env/asm"
	"github.com/ayoul3/asm-env/decrypt"
	log "github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 2 {
		log.Info("No arguments provided. Will gracefully exit")
		os.Exit(0)
	}

	d := decrypt.Handler{
		AsmClient: asm.NewClient(asm.NewAPI()),
		Args:      os.Args[1:],
		Envs:      prepareEnvVars(),
	}
	err := d.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func prepareEnvVars() map[string]string {
	envs := make(map[string]string, len(os.Environ()))
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		envs[pair[0]] = pair[1]
	}
	return envs
}
