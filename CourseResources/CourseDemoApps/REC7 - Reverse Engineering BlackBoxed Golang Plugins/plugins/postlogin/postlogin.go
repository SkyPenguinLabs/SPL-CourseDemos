package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CE(x error) {
	if x != nil {
		log.Fatal(x)
	}
}

func PostLogin(username, password string) bool {
	var RequestResult map[string]interface{}

	PD := map[string]string{
		"username": username,
		"password": password,
	}

	jd, x := json.Marshal(PD)
	CE(x)

	///// Some BS login functionality just to make it an app.
	req, x := http.NewRequest(
		"POST",
		"https://httpbin.org/post", bytes.NewBuffer(jd),
	)
	CE(x)

	///// Headers & Opts
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, x := client.Do(req)
	CE(x)
	defer resp.Body.Close()

	resp_body, x := io.ReadAll(resp.Body)
	CE(x)

	if x := json.Unmarshal(resp_body, &RequestResult); x != nil {
		CE(x)
	}
	fmt.Println(RequestResult)
	return true
}
