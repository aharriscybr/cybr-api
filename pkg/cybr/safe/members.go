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

// Onboard Safe
func AddMember(urlID *string, permblock []byte, token *string, domain *string) (bool, error) {

	client := hClient.GetClient()

	api_uri := "https://" + *domain + ".privilegecloud.cyberark.cloud/PasswordVault/API/Safes/" + *urlID + "/Members"
	method := "POST"

	req, e := http.NewRequest(method, api_uri, bytes.NewBuffer(permblock))
	if e != nil {
		log.Println("Failure to initialize HTTP Request")
	}

	bearer := "Bearer " + *token
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Printf("Unhandled error updating membership for [%s].", *urlID)
		return false, err
	}

	if response.StatusCode == 201 {

		responseData := cybrtypes.SafeData{}

		json.NewDecoder(response.Body).Decode(&responseData)

		log.Printf("Successfully added [%s] to [%s]",  *responseData.Owner, *responseData.Name)

		return true, err

	} else {
		
		log.Println("Unable to process membership update, the vault rejected your request: ")
		log.Printf("Status [%d]", response.StatusCode)
		
		r, err := io.ReadAll(response.Body)
		if err != nil {
		
			log.Fatal(err)
		
		}

		log.Println(string(r))
		return false, err
	}

}