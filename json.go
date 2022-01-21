package lambda_utils_go

import (
	"encoding/json"
	jsonpatch "github.com/evanphx/json-patch"
	"log"
)

func ApplyJsonPatch(original interface{}, patchBody string, result interface{}) error {
	originalJson, err := json.Marshal(original)
	if err != nil {
		return err
	}

	patch, err := jsonpatch.DecodePatch([]byte(patchBody))
	if err != nil {
		return err
	}

	modified, err := patch.Apply(originalJson)
	if err != nil {
		return err
	}

	log.Printf("original: %s", originalJson)
	log.Printf("patch: %s", patchBody)
	log.Printf("result: %s", string(modified))

	if err := json.Unmarshal(modified, result); err != nil {
		return err
	}

	return nil
}
