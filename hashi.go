package cyberark_hashiimpl

import (
	"log"
	"net/http"

	account "github.com/aharriscybr/cybr-api/pkg/cybr/account"
	auth "github.com/aharriscybr/cybr-api/pkg/cybr/auth"
	client "github.com/aharriscybr/cybr-api/pkg/cybr/http"
	types "github.com/aharriscybr/cybr-api/pkg/cybr/types"
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

	cred := types.CloudConfig {
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

func CreateAccount(name *string, address *string, username *string, platform *string, safe *string, secrettype *string, secret *string, authToken *string, domain *string) ([]byte, error) {

	port := "5432"
	dbn  := "dbo.test"

	dbProps := types.AccountProps {
		Port: &port,
		DBName: &dbn,
	}

	newAccount := types.Credential {
		Name: name,
		Address: address,
		UserName: username,
		Platform: platform,
		SecretType: secrettype,
		Secret: secret,
		SafeName: safe,
		Props: dbProps,
	}

	log.Printf("Processing Account Attribute: %s", *name)
	log.Printf("Processing Account Attribute: %s", *address)
	log.Printf("Processing Account Attribute: %s", *username)
	log.Printf("Processing Account Attribute: %s", *platform)
	log.Printf("Processing Account Attribute: %s", *safe)
	log.Printf("Processing Account Attribute: %s", *secrettype)

	account.Onboard(&newAccount, authToken, domain)

	return nil, nil

}