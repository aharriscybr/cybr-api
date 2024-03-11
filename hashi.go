package cybrapi_hashiimpl

import (
	"log"
	"net/http"

	account "github.com/aharriscybr/cybr-api/pkg/cybr/account"
	auth "github.com/aharriscybr/cybr-api/pkg/cybr/auth"
	client "github.com/aharriscybr/cybr-api/pkg/cybr/http"
	cybrtypes "github.com/aharriscybr/cybr-api/pkg/cybr/types"
)

type Client struct {
	AuthToken *string
	HTTPClient *http.Client
	Domain *string
}


// Register new Authentication session and client details
func NewClient(tenant *string, domain *string, clientid *string, clientsecret *string) (*Client, error) {

	hclient := client.GetClient()

	if tenant == nil || domain == nil || clientid == nil || clientsecret == nil {
		return nil, nil
	}

	c := Client {
		HTTPClient: hclient,
		Domain: domain,
	}

	cred := cybrtypes.CloudConfig {
		Tenant: tenant,
		Domain: domain,
		ClientID: clientid,
		ClientSecret: clientsecret,
	}

	token, result, err := auth.GetIdentityToken(&cred)
	if err != nil {
		log.Println(err)
		return &c, err
	}

	if result {

		log.Println("Successfully authenticated to shared services.")

		c.AuthToken = &token

		return &c, nil
	}

	return &c, nil

}

// Interface to Creation of account
func CreateAccount(c *cybrtypes.Credential, authToken *string, domain *string) (string, error) {


	id, err := account.Onboard(c, authToken, domain)
	if err != nil {

		log.Println("Unable to onboard account, please check information and try again.")

		return "", nil

	}

	return id, nil
}

// Interface to Delete Account
func RemoveAccount(id *string, authToken *string, domain *string) (bool, error) {

	result, err := account.Remove(id, authToken, domain)
	if err != nil {
		
		return false, nil
	}

	if result {
		return result, nil
	}

	return false, nil

}