package cyberarkapi_hashiimpl

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

func CreateAccount(cred *cybrtypes.Credential, authToken *string, domain *string) ([]byte, error) {

	// dbn := cred.Props.DBName
	// p := cred.Props.Port

	//log.Printf("Processing Account Attribute: %s", *cred.Name)
	// log.Printf("Processing Account Attribute: %s", *dbn)
	// log.Printf("Processing Account Attribute: %s", *p)

	account.Onboard(cred, authToken, domain)

	return nil, nil
}
