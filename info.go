package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "net/http"

func GetInfo() map[string]interface{} {
	resp, err := http.Get("https://bandwidth.waits.io/info.json")
	Check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	var jErr = json.Unmarshal(body, &data)
	Check(jErr)

	return data
}

func ShowInfo() {
	info := GetInfo()
	fmt.Printf("IPv4 address:\t%s\nHostname:\t%s\n", info["ip"], info["host"])
}

func ShowIp() {
	info := GetInfo()
	fmt.Printf("%s\n", info["ip"])
}
