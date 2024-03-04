package main

import (
	"log"

	"github.com/aharriscybr/cybr-api/pkg/cybr/auth"
	"github.com/aharriscybr/cybr-api/pkg/cybr/types"
)


func main() {
	
	// todo

	userobj := types.Authn {
		ClientID: "conjur_automation@cyberark.cloud.4272",
		ClientSecret: "Apple-Banana-Computer12",
		Tenant: "aap4847",
	}

	SyncPerms, e := types.SecretsHub()
	if e != nil {
		log.Fatal(e)
	}

	log.Println(string(SyncPerms))

	token, result, err := auth.GetIdentityToken(&userobj)
	if !result {
		log.Print(err)
	}

	if result {
		log.Print(token)
	}


}