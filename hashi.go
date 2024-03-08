package cyberark_hashiimpl

import (
	"log"
	"net/http"

	auth "github.com/aharriscybr/cybr-api/pkg/cybr/auth"
	client "github.com/aharriscybr/cybr-api/pkg/cybr/http"
	types "github.com/aharriscybr/cybr-api/pkg/cybr/types"
)

type Client struct {
	VaultAPI *string
	AuthToken *string
	HTTPClient *http.Client
}

func NewClient(tenant *string, domain *string, clientid *string, clientsecret *string) (*Client, error) {

	hclient := client.GetClient()

	if tenant == nil || domain == nil || clientid == nil || clientsecret == nil {
		return nil, nil
	}

	c := Client {
		HTTPClient: hclient,
	}

	cred := types.CloudConfig {
		Tenant: tenant,
		Domain: domain,
		ClientID: clientid,
		ClientSecret: clientsecret,
	}
	log.Printf("Attribute: %s", *tenant)
	log.Printf("Attribute: %s", *domain)
	log.Printf("Attribute: %s", *clientid)
	log.Printf("Attribute: %s", *clientsecret)

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

// func CreateAccount(name *string, address *string, username *string, platform *string, safe *string, secrettype *string, secret *string) ([]byte, error) {

// 	newAccount := types.Credential {
// 		Name: name,
// 		Address: address,
// 		UserName: username,
// 		Platform: platform,
// 		SecretType: secrettype,
// 		Secret: secret,
// 	}

// 	account.Onboard(newAccount, )

// }