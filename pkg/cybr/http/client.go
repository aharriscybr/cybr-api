package cyberarkapi_client

import (
	"net/http"
	"time"
)

// Return configured HTTP Client
func GetClient() *http.Client {

	c := &http.Client{

		Transport: &http.Transport{
			TLSHandshakeTimeout: 700 * time.Millisecond,
		},
	}
	
	return c

}