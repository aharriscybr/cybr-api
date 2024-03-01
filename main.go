package main

import (
	"log"

	"github.com/aharriscybr/cybr-api/pkg/cybr/types"
)


func main() {
	
	// todo
	SHUser, err := types.FullAdmin("user", "andrew");
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", string(SHUser))

}