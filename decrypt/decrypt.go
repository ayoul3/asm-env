package decrypt

import (
	"os"

	"emperror.dev/errors"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) InjectDecryptedVars() (err error) {
	var decryptedValue string

	for key, value := range h.Envs {
		if !h.AsmClient.IsSecret(value) {
			continue
		}
		log.Debugf("Decrypting key ID %s", value)
		if decryptedValue, err = h.AsmClient.GetSecret(value); err != nil {
			return errors.Wrapf(err, "GetSecret %s ", value)
		}
		os.Setenv(key, decryptedValue)
	}
	return nil
}
