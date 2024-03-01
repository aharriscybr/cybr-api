package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	// API Includes
	h "github.com/aharriscybr/cybr-api/pkg/cybr/http"
	"github.com/aharriscybr/cybr-api/pkg/cybr/types"
)

// Authenticate to Shared Services Identity Platform
func HandleIdentityAuthn(a types.Authn) string {

	// Get configured client
	client := h.GetClient();

	authnUrl := "https://" + a.Tenant + ".id.cyberark.cloud/oauth2/platformtoken"
	method := "POST"

	payload := strings.NewReader("client_id=" + url.QueryEscape(a.ClientID) + "&grant_type=" + a.GrantType + "&client_secret=" + a.ClientSecret)

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

			log.Fatal(jsonError)

		}
		return string(authzToken.Access_token)

	} else {

		log.Println(string(body))
		log.Fatal("Failed to authenticate:", res.StatusCode)

	}

	log.Fatal("Unable to authenticate to Shared Services.")
	return "false"

}