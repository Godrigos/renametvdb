package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Define application wide variables
var authorization string

// Initialize the variables before main execution
func init() {

	//------------------------------- Login ----------------------------------
	// Prepare the login body for the request
	reqBody, err := json.Marshal(map[string]string{
		"apikey": "4205WF0K63GQTEAO",
	})
	if err != nil {
		panic(err)
	}

	// Execute the request and store the response
	respLogin, err := http.Post("https://api.thetvdb.com/login",
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	if respLogin.StatusCode == 401 {
		panic("Not Authorized! Invalid credentials and/or API token.")
	}
	defer respLogin.Body.Close()

	body, err := ioutil.ReadAll(respLogin.Body)
	if err != nil {
		panic(err)
	}

	var token authentication
	err = json.Unmarshal(body, &token)
	if err != nil {
		panic(err)
	}

	// Compose the Authorization header string
	authorization = "Bearer " + token.Token
}

type authentication struct {
	Token string `json:"token"`
}

type languages struct {
	Data []struct {
		Abbreviation string `json:"abbreviation"`
		EnglishName  string `json:"englishName"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
	} `json:"data"`
}
