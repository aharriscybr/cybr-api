package auth

func handleIdentityAuthn() string {

	h.httpClient();
	// authnUrl := "https://" + tenant + ".id.cyberark.cloud/oauth2/platformtoken"
	// method := "POST"

	// payload := strings.NewReader("client_id=" + url.QueryEscape(pasUser) + "&grant_type=" + gt + "&client_secret=" + pasPassword)

	// req, err := http.NewRequest(method, authnUrl, payload)
	// if err != nil {

	// 	log.Fatal(err)

	// }

	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// res, err := client.Do(req)
	// if err != nil {

	// 	log.Fatal(err)

	// }
	// defer res.Body.Close()

	// body, err := io.ReadAll(res.Body)
	// if err != nil {

	// 	log.Fatal(err)

	// }

	// if res.StatusCode == 200 {

	// 	authzToken := token{}
	// 	jsonError := json.Unmarshal(body, &authzToken)
	// 	if jsonError != nil {

	// 		log.Fatal(jsonError)

	// 	}
	// 	return string(authzToken.Access_token)

	// } else {

	// 	log.Println(string(body))
	// 	log.Fatal("Failed to authenticate:", res.StatusCode)

	// }

	// log.Fatal("Unable to authenticate to ISPSS.")
	// return "Failed."

}