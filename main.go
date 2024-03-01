package main

import (
	"log"

	"github.com/aharriscybr/cybr-api/pkg/cybr/auth"
	"github.com/aharriscybr/cybr-api/pkg/cybr/types"
)


func main() {
	
	// todo
	SHUser, err := types.Approver("user", "andrew");
	if err != nil {
		log.Fatal(err)
	}

	session := types.Authn {
		Tenant: "aap4847",
		ClientID: "conjur_automation@cyberark.cloud.4272",
		ClientSecret: "Apple-Banana-Computer12",
	}

	token, result := auth.GetIdentityToken(session)
	if !result {
		log.Fatal(result)
	}

	log.Printf("%s", token)
	log.Printf("%s", SHUser)

}