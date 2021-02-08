package decrypt

import (
	"os"
	"os/exec"
	"syscall"

	"emperror.dev/errors"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) StartProcess() (err error) {
	var process *os.Process

	attr := os.ProcAttr{
		Dir: ".",
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
		Sys: &syscall.SysProcAttr{Noctty: true},
	}
	binary, err := exec.LookPath(h.Args[0])
	if err != nil {
		return errors.Wrapf(err, "LookPath %s", h.Args[0])
	}
	log.Debugf("Found absolute path %s", binary)
	if process, err = os.StartProcess(binary, h.Args, &attr); err != nil {
		return errors.Wrapf(err, "StartProcess")
	}

	log.Debugf("Releasing process %d", process.Pid)
	return process.Release()
}
