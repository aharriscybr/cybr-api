package cyberark

import (
	"log"
	"net/http"
	"time"

	"github.com/aharriscybr/cybr-api/pkg/cybr/auth"
	"github.com/aharriscybr/cybr-api/pkg/cybr/types"
)

var (
	api_uri string
)

type Client struct {
	URI string
	HTTPClient *http.Client
	Auth string
}

// Return configured HTTP Client
func GetClient(conf *types.CloudConfig) (*Client, error) {

	if conf.Tenant == "" || conf.ClientID == "" || conf.ClientSecret == "" || conf.Domain == "" || conf.Action == "" {

		return nil, nil

	} else {

		if conf.Action == "auth" {
			api_uri = "https://" + conf.Tenant + ".id.cyberark.cloud/oauth2/platformtoken"

			token, result, err := auth.GetIdentityToken(conf)
			if err != nil {
				log.Fatal(err)
			}
			if result {
				c := Client {
					HTTPClient: &http.Client{Timeout: 10 * time.Second},
					URI: api_uri,
					Auth: token,
				}
				return &c, nil
			}

		}
		// TODO Add onboarding client information
		if conf.Action == "create" {
			api_uri = "https://" + conf.Domain + ".privilegecloud.cyberark.cloud/PasswordVault/API/Accounts"
			c := Client {
				HTTPClient: &http.Client{Timeout: 10 * time.Second},
				URI: api_uri,
				Auth: "blank",
			}	
			return &c, nil
		}
	
	}

	return nil, nil
}