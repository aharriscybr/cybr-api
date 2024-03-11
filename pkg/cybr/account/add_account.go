package cyberarkapi_account

import (
	"bytes"
	"io"
	"log"
	"net/http"

	hClient "github.com/aharriscybr/cybr-api/pkg/cybr/http"
	cybrtypes "github.com/aharriscybr/cybr-api/pkg/cybr/types"
)

// Onboard Credentials into safe
func Onboard(cred *cybrtypes.Credential, token *string, domain *string) (bool, error) {

	client := hClient.GetClient()

	api_uri := "https://" + *domain + ".privilegecloud.cyberark.cloud/PasswordVault/API/Accounts"
	method := "POST"

	body, e := cybrtypes.Cred(cred)
	if e != nil {
		log.Panicln("Unable to format Credential object.")
	}

	req, e := http.NewRequest(method, api_uri, bytes.NewBuffer(body))
	if e != nil {
		log.Fatal("Unable to construct api request.")
	}

	bearer := "Bearer " + *token
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Println("Unhandled error onboarding credential.")
		return false, err
	}

	if response.StatusCode == 201 {

		text, err := io.ReadAll(response.Request.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Onboarded %s to %s", *cred.Name, *cred.SafeName)
		log.Println(string(text))

	} else {
		log.Println("Your Credential was not onboarded, the vault rejected your request: ")
		log.Printf("Status %d", response.StatusCode)
		r, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(r))
		return false, err
	}

	return true, nil

}

func Get() (bool, error) {

	return false, nil
}

func Update() (bool, error) {

	return false, nil
}

func Delete(cred *cybrtypes.Credential, token *string, domain *string) (bool, error) {

	return false, nil
}