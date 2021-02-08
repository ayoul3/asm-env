package main

import (
	"os"
	"strings"

	"github.com/ayoul3/asm-env/asm"
	"github.com/ayoul3/asm-env/decrypt"
	log "github.com/sirupsen/logrus"
)

func init() {
	if len(os.Args) < 2 {
		log.Warn("No arguments provided. Will gracefully exit")
		os.Exit(0)
	}
	if os.Getenv("DEBUG_ASM_ENV") == "true" {
		log.SetLevel(log.DebugLevel)
	}
}
func main() {
	var err error
	var asmClient *asm.Client

	if asmClient, err = asm.NewClient(asm.NewAPI()); err != nil {
		log.Fatal(err)
	}
	d := decrypt.Handler{
		AsmClient: asmClient,
		Args:      os.Args[1:],
		Envs:      prepareEnvVars(),
	}

	if err = d.Start(); err != nil {
		log.Fatal(err)
	}
}

func prepareEnvVars() map[string]string {
	log.Debug("Preparing env variables")

	envs := make(map[string]string, len(os.Environ()))
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		envs[pair[0]] = pair[1]
	}
	return envs
}
