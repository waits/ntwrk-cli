package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "net/http"

func GetInfo(url string) map[string]interface{} {
	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	var jErr = json.Unmarshal(body, &data)
	Check(jErr)

	return data
}

func ShowInfo(args []string) {
	var url string
	if len(args) > 0 {
		url = fmt.Sprintf("https://ntwrk.waits.io/info.json?ip=%s", args[0])
	} else {
		url = "https://ntwrk.waits.io/info.json"
	}
	info := GetInfo(url)
	var location string
	if info["city"] != nil {
		location = fmt.Sprintf("%s, %s", info["city"], info["country"])
	} else {
		location = fmt.Sprintf("%s", info["country"])
	}
	fmt.Printf("IPv4 address:\t%s\nHostname:\t%s\nLocation:\t%s\n", info["ip"], info["host"], location)
}

func ShowIp() {
	info := GetInfo("https://ntwrk.waits.io/info.json")
	fmt.Printf("%s\n", info["ip"])
}
