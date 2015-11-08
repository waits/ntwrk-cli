package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "net/http"

func ip() {
	resp, err := http.Get("https://bandwidth.waits.io/ip.json")
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	var jErr = json.Unmarshal(body, &data)
	check(jErr)

	fmt.Printf("%s\n", data["ip"])
}
