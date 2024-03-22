package cyberarkapi_hashiimpl

import (
	"log"
	"net/http"

	account "github.com/aharriscybr/cybr-api/pkg/cybr/account"
	auth "github.com/aharriscybr/cybr-api/pkg/cybr/auth"
	client "github.com/aharriscybr/cybr-api/pkg/cybr/http"
	safes "github.com/aharriscybr/cybr-api/pkg/cybr/safe"
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
		log.Println("Required information has not been set.")
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

/*
=========================================
* Account Management
=========================================
*/

// Interface to Creation of account
func CreateAccount(c *cybrtypes.Credential, authToken *string, domain *string) (string, error) {

	id, err := account.Onboard(c, authToken, domain)
	if err != nil {

		log.Println("Unable to onboard account, please check information and try again.")

		return "", nil

	}

	return id, nil
}

// Interface to Read account
func GetAccount(id *string, authToken *string, domain *string) (*cybrtypes.CredentialResponse, error) {

	details, err := account.Details(id, authToken, domain)
	if err != nil {

		log.Println("Error retrieving details.... boop.")
		return nil, err

	}

	return details, nil

}

// Interface to Delete Account
func RemoveAccount(id *string, authToken *string, domain *string) (error) {

	result, err := account.Remove(id, authToken, domain)
	if err != nil {
		
		return err
	}

	if result {

		return nil

	} else {

		return nil
	}

}

/*
=========================================
* Safe Management
=========================================
*/

// Interface to Creation of safe
func CreateSafe(s *cybrtypes.SafeData, authToken *string, domain *string) (*cybrtypes.SafeData, error) {

	safe, err := safes.Onboard(s, authToken, domain)
	if err != nil {

		log.Println("Unable to onboard safe, please check information and try again.")

		return nil, nil

	}

	log.Printf("Generating Permission %s.", *s.Level)
	log.Printf("Ownership Properties: %s, %s, %s", *s.Owner, *s.OwnerType, *s.Level)

	var block []byte

	if *s.Level == "full" {
		block, err = cybrtypes.FullAdmin(s.OwnerType, s.Owner)
		if err != nil {
			log.Println("Error generating permissions block.")
		}
	}
	if *s.Level == "read" {
		block, err = cybrtypes.ReadOnly(s.OwnerType, s.Owner)
		if err != nil {
			log.Println("Error generating permissions block.")
		}
	}
	if *s.Level == "approver" {
		block, err = cybrtypes.Approver(s.OwnerType, s.Owner)
		if err != nil {

			log.Println("Error generating permissions block.")
		}
	}
	if *s.Level == "manager" {
		block, err = cybrtypes.Manager(s.OwnerType, s.Owner)
		if err != nil {
			log.Println("Error generating permissions block.")
		}
	}

	log.Println(string(block))

	result, err := safes.AddMember(s.Name, block, authToken, domain)
	if err != nil {
		log.Println("Unable to update membership.")
	}

	if result {
		
		return safe, nil
	
	} else {

		log.Println("One or more errors occurred in provisioning the safe, please check your values and logs and try again.")
		return nil, nil

	}
	
}

// Interface to Read Safe
func GetSafe(id *string, authToken *string, domain *string) (*cybrtypes.SafeData, error) {

	details, err := safes.Details(id, authToken, domain)
	if err != nil {

		log.Println("Error retrieving details.... boop.")
		return nil, err

	}

	return details, nil

}

// Interface to Delete Safe
func RemoveSafe(id *string, authToken *string, domain *string) (error) {

	result, err := safes.Remove(id, authToken, domain)
	if err != nil {
		
		return err
	}

	if result {

		return nil

	} else {

		return nil
	}

}