package main

import (
	"encoding/json"
	"log"

	cybrtypes "github.com/aharriscybr/cybr-api/pkg/cybr/types"
)

func main(){

	var name, add, un, pf, sn, st, sc string

	name = "blah"
	add = "blah1"
	un = "blah2"
	pf = "blah3"
	sn = "blah4"
	st = "blah5"
	sc = "blah6"
	// port := "5432"
	// dbn := "test"
	//dsn := "test2"

	props := cybrtypes.AccountProps {
		// Port: &port,
		// DBName: &dbn,
	}

	cred := cybrtypes.Credential {
		Name: &name,
		Address: &add,
		UserName: &un,
		Platform: &pf,
		SafeName: &sn,
		SecretType: &st,
		Secret: &sc,
		Props: &props,
	}

	thisBlock, err := json.Marshal(cred)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(thisBlock))

}