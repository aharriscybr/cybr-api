package cyberarkapi_account

import (
	"bytes"
	"io"
	"log"
	"net/http"

	auth "github.com/aharriscybr/cybr-api/pkg/cybr/auth"
	hClient "github.com/aharriscybr/cybr-api/pkg/cybr/http"
	types "github.com/aharriscybr/cybr-api/pkg/cybr/types"
)

// Onboard Credentials into safe
func Onboard(cred *types.Credential, conf *types.CloudConfig) (bool, error) {

	client := hClient.GetClient()

	api_uri := "https://" + *conf.Domain + ".privilegecloud.cyberark.cloud/PasswordVault/API/Accounts"
	method := "POST"

	body, e := types.Cred(cred)
	if e != nil {
		log.Panicln("Unable to format Credential object.")
	}

	req, e := http.NewRequest(method, api_uri, bytes.NewBuffer(body))
	if e != nil {
		log.Fatal("Unable to construct api request.")
	}

	token, result, e := auth.GetIdentityToken(conf)
	if e != nil {
		log.Printf("Unable to Authenticate to %s", *conf.Tenant)
		return false, e
	}
	if !result {
		log.Printf("Unhandled error")
	}
	
	bearer := "Bearer " + token
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Println("Unhandled error onboarding credential.")
		return false, e
	}

	if response.StatusCode == 201 {

		log.Printf("Onboarded %s to %s", cred.Name, cred.SafeName)

	} else {
		log.Println("Your Credential was not onboarded, the vault rejected your request: ")
		log.Printf("Status %d", response.StatusCode)
		r, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(r))
	}

	return true, nil

}

