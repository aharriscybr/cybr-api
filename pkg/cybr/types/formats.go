package cyberarkapi_types

import (
	"encoding/json"
	"log"
)

// Build credential object for onboarding via API
func Cred(cred *Credential) ([]byte, error) {

	d := &cred;

	jsonData, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData, nil
}


// Build safe object for onboarding via API
func Safe(safe *SafeData) ([]byte, error) {

	d := &safe;

	jsonData, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData, nil
}