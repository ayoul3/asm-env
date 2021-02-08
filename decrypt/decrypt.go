package decrypt

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"emperror.dev/errors"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) InjectDecryptedVars() (err error) {
	var decryptedValue, nestedValue string

	for key, asmKeyID := range h.Envs {
		if !h.AsmClient.IsSecret(asmKeyID) {
			continue
		}
		log.Debugf("Decrypting key ID %s", asmKeyID)
		if decryptedValue, err = h.AsmClient.GetSecret(asmKeyID); err != nil {
			return errors.Wrapf(err, "GetSecret %s ", asmKeyID)
		}
		log.Debugf("Setting secret value in env var %s", key)

		if nestedValue, err = ExtractKeyWhenJson(asmKeyID, decryptedValue); err != nil {
			return errors.Wrapf(err, "ExtractKeyWhenJson %s ", key)
		}
		os.Setenv(key, nestedValue)
	}
	return nil
}

func ExtractKeyWhenJson(key, value string) (out string, err error) {
	var parsed map[string]string

	if !strings.Contains(key, "#") {
		return value, nil
	}

	keyParts := strings.Split(key, "#")
	desiredKey := keyParts[1]
	log.Debugf("Looking for nested key %s in secret %s", desiredKey, key)

	if err = json.Unmarshal([]byte(value), &parsed); err != nil {
		return "", errors.Wrap(err, "Unmarshal: Only simple Json structured secrets are accepted ")
	}
	for k, v := range parsed {
		if k == desiredKey {
			log.Debugf("Found nested key %s in secret %s", desiredKey, key)
			return v, nil
		}
	}
	return "", fmt.Errorf("key %s not found in Json value", key)
}
