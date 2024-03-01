package hClient

import (
	"net/http"
	"time"
)


func GetClient() *http.Client {

	c := &http.Client{

		Transport: &http.Transport{
			TLSHandshakeTimeout: 700 * time.Millisecond,
		},
	}
	
	return c

}