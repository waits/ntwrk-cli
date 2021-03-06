package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "net/http"

const infoUrl = "https://ntwrk.waits.io/info.json"

func fetchInfo(url string) map[string]interface{} {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	var jErr = json.Unmarshal(body, &data)
	check(jErr)

	return data
}

func geo(args []string) {
	var url string
	if len(args) > 0 {
		url = fmt.Sprintf("%s?ip=%s", infoUrl, args[0])
	} else {
		url = infoUrl
	}
	info := fetchInfo(url)
	fmt.Printf("GeoIP data for %s\n\n", info["ip"])
	if info["isp"] != nil {
		fmt.Printf("ISP:\t\t%s\n", info["isp"])
	}
	if info["city"] != nil && info["region"] != nil {
		fmt.Printf("Location:\t%s, %s, %s\n", info["city"], info["region"], info["country"])
	} else {
		fmt.Printf("Location:\t%s\n", info["country"])
	}
	if info["latitude"] != nil && info["longitude"] != nil {
		fmt.Printf("Coordinates:\t%f, %f\n", info["latitude"], info["longitude"])
	}
	if info["time_zone"] != nil {
		fmt.Printf("Time zone:\t%s\n", info["time_zone"])
	}
}

func info(args []string) {
	var url string
	if len(args) > 0 {
		url = fmt.Sprintf("%s?ip=%s", infoUrl, args[0])
	} else {
		url = infoUrl
	}
	info := fetchInfo(url)
	fmt.Printf("IPv4 address:\t%s\nHostname:\t%s\n", info["ip"], info["host"])
}

func ip() {
	resp := fetchInfo("https://ntwrk.waits.io/ip.json")
	fmt.Printf("%s\n", resp["ip"])
}
