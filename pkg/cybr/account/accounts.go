package cyberarkapi_account

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	hClient "github.com/aharriscybr/cybr-api/pkg/cybr/http"
	cybrtypes "github.com/aharriscybr/cybr-api/pkg/cybr/types"
)

// Onboard Credentials into safe
func Onboard(cred *cybrtypes.Credential, token *string, domain *string) (string, error) {

	client := hClient.GetClient()

	api_uri := "https://" + *domain + ".privilegecloud.cyberark.cloud/PasswordVault/API/Accounts"
	method := "POST"

	body, e := cybrtypes.Cred(cred)
	if e != nil {
		log.Println("Unable to format Credential object.")
	}

	req, e := http.NewRequest(method, api_uri, bytes.NewBuffer(body))
	if e != nil {
		log.Println("Failure to initialize HTTP Request")
	}

	bearer := "Bearer " + *token
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Println("Unhandled error onboarding credential.")
		return "", err
	}

	if response.StatusCode == 201 {

		responseData := cybrtypes.CredentialResponse{}

		json.NewDecoder(response.Body).Decode(&responseData)

		log.Printf("Successfully Onboarded [%s] to [%s]", *cred.Name, *cred.SafeName)

		credID := *responseData.CredID

		return credID, err

	} else {
		
		log.Println("Your Credential was not onboarded, the vault rejected your request: ")
		log.Printf("Status [%d]", response.StatusCode)
		
		r, err := io.ReadAll(response.Body)
		if err != nil {
		
			log.Fatal(err)
		
		}

		log.Println(string(r))
		return "", err
	}

}

func Details(id *string, token *string, domain *string) (*cybrtypes.CredentialResponse, error) {

	client := hClient.GetClient()

	api_uri := "https://" + *domain + ".privilegecloud.cyberark.cloud/PasswordVault/API/Accounts/" + *id
	method := "GET"

	req, e := http.NewRequest(method, api_uri, nil)
	if e != nil {
		log.Fatal("Unable to construct api request.")
	}

	bearer := "Bearer " + *token
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Println("Failure to initialize HTTP Request")
		return nil, err
	}

	if response.StatusCode == 200 {

		responseData := cybrtypes.CredentialResponse{}

		json.NewDecoder(response.Body).Decode(&responseData)

		log.Printf("Response Code %d: Successfully fetched details for %s", response.StatusCode, *id)
		return &responseData, nil

	} else {

		log.Printf("Unable to fetch [%s] details: ", *id)
		log.Printf("Status [%d]", response.StatusCode)
		
		r, err := io.ReadAll(response.Body)
		if err != nil {
		
			log.Fatal(err)
		
		}

		log.Println(string(r))
		return nil, err

	}
}

func Remove(id *string, token *string, domain *string) (bool, error) {

	client := hClient.GetClient()

	api_uri := "https://" + *domain + ".privilegecloud.cyberark.cloud/PasswordVault/API/Accounts/" + *id
	method := "DELETE"

	req, e := http.NewRequest(method, api_uri, nil)
	if e != nil {
		log.Fatal("Unable to construct api request.")
	}

	bearer := "Bearer " + *token
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Println("Failure to initialize HTTP Request")
		return false, err
	}

	if response.StatusCode == 204 {

		log.Printf("Response Code [%d]: Successfully removed [%s]", response.StatusCode, *id)
		return true, nil

	} else {

		log.Println("Your Credential was not onboarded, the vault rejected your request: ")
		log.Printf("Status [%d]", response.StatusCode)
		
		r, err := io.ReadAll(response.Body)
		if err != nil {
		
			log.Fatal(err)
		
		}

		log.Println(string(r))
		return false, err

	}
}