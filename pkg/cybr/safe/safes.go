package cyberarkapi_safe

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
func Onboard(safe *cybrtypes.SafeData, token *string, domain *string) (*cybrtypes.SafeData, error) {

	client := hClient.GetClient()

	api_uri := "https://" + *domain + ".privilegecloud.cyberark.cloud/PasswordVault/API/Safes"
	method := "POST"

	body, e := cybrtypes.Safe(safe)
	if e != nil {
		log.Println("Unable to format Safe object.")
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
		log.Println("Unhandled error onboarding safe.")
		return nil, err
	}

	if response.StatusCode == 201 {

		responseData := cybrtypes.SafeData{}

		err := json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			log.Println("Unable to decode JSON response from vault API")
			return nil, err
		}

		log.Printf("Successfully Onboarded [%s]: Name [%s] - ID [%s]", *responseData.Name, *responseData.URLID, *responseData.NUMBER,)

		return &responseData, err

	} else {
		
		log.Println("Your Safe was not onboarded, the vault rejected your request: ")
		log.Printf("Status [%d]", response.StatusCode)
		
		r, err := io.ReadAll(response.Body)
		if err != nil {
		
			log.Fatal(err)
		
		}

		log.Println(string(r))
		return nil, err
	}

}

// func Details(id *string, token *string, domain *string) (*cybrtypes.CredentialResponse, error) {

// 	client := hClient.GetClient()

// 	api_uri := "https://" + *domain + ".privilegecloud.cyberark.cloud/PasswordVault/API/Accounts/" + *id
// 	method := "GET"

// 	req, e := http.NewRequest(method, api_uri, nil)
// 	if e != nil {
// 		log.Fatal("Unable to construct api request.")
// 	}

// 	bearer := "Bearer " + *token
// 	req.Header.Add("Authorization", bearer)
// 	req.Header.Add("Content-Type", "application/json")

// 	response, err := client.Do(req)
// 	if err != nil {
// 		log.Println("Failure to initialize HTTP Request")
// 		return nil, err
// 	}

// 	if response.StatusCode == 200 {

// 		responseData := cybrtypes.CredentialResponse{}

// 		err := json.NewDecoder(response.Body).Decode(&responseData)
		// if err != nil {
		// 	log.Println("Unable to decode JSON response from vault API")
		// 	return nil, err
		// }

// 		log.Printf("Response Code %d: Successfully fetched details for %s", response.StatusCode, *id)
// 		return &responseData, nil

// 	} else {

// 		log.Printf("Unable to fetch [%s] details: ", *id)
// 		log.Printf("Status [%d]", response.StatusCode)
		
// 		r, err := io.ReadAll(response.Body)
// 		if err != nil {
		
// 			log.Fatal(err)
		
// 		}

// 		log.Println(string(r))
// 		return nil, err

// 	}
// }

func Remove(id *string, token *string, domain *string) (bool, error) {

	client := hClient.GetClient()

	api_uri := "https://" + *domain + ".privilegecloud.cyberark.cloud/PasswordVault/API/Safes/" + *id
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

		log.Println("Your Safe was not destroyed, the vault rejected your request: ")
		log.Printf("Status [%d]", response.StatusCode)
		
		r, err := io.ReadAll(response.Body)
		if err != nil {
		
			log.Fatal(err)
		
		}

		log.Println(string(r))
		return false, err

	}
}