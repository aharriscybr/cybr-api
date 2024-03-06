package cyberarkapi_auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	// API Includes

	client "github.com/aharriscybr/cybr-api/pkg/cybr/http"
	types "github.com/aharriscybr/cybr-api/pkg/cybr/types"
)

// Authenticate to Shared Services Identity Platform
// This retrieves an authentication token for use within Privilege Cloud
func GetIdentityToken(a *types.CloudConfig) (token string, result bool, err error) {

	// Get configured client
	client := client.GetClient()

	authnUrl := "https://" + *a.Tenant + ".id.cyberark.cloud/oauth2/platformtoken"
	method := "POST"

	payload := strings.NewReader("client_id=" + url.QueryEscape(*a.ClientID) + "&grant_type=client_credentials&client_secret=" + *a.ClientSecret)

	log.Printf("Attempting to authenticate %s to %s", *a.ClientID, authnUrl)

	req, err := http.NewRequest(method, authnUrl, payload)
	if err != nil {

		log.Fatal(err)

	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {

	log.Fatal(err)

	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {

		log.Fatal(err)

	}

	if res.StatusCode == 200 {

		authzToken := types.Token{}
		jsonError := json.Unmarshal(body, &authzToken)
		if jsonError != nil {

			log.Println(jsonError)

		}
		return string(authzToken.Access_token), true, nil

	} else {

		log.Println(string(body))
		log.Printf("Failed to authenticate: %d", res.StatusCode)

		return "Unable to authenticate to Shared Services.", false, err

	}

}