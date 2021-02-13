package decrypt

import (
	"fmt"
	"os/exec"
	"syscall"

	"emperror.dev/errors"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) StartProcess() (err error) {
	envv := make([]string, 0)
	binary, err := exec.LookPath(h.Args[0])
	if err != nil {
		return errors.Wrapf(err, "LookPath %s", h.Args[0])
	}
	for key, val := range h.Envs {
		envv = append(envv, fmt.Sprintf("%s=%s", key, val))
	}
	log.Debugf("Found absolute path %s", binary)
	return syscall.Exec(binary, h.Args, envv)
}
