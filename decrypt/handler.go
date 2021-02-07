package decrypt

import (
	"emperror.dev/errors"
	"github.com/ayoul3/asm-env/asm"
)

type Handler struct {
	Args      []string
	Envs      map[string]string
	AsmClient *asm.Client
}

func (h *Handler) Start() (err error) {
	if err = h.InjectDecryptedVars(); err != nil {
		return errors.Wrapf(err, "InjectDecryptedVars ")
	}

	return h.StartProcess()
}
