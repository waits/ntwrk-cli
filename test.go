package main

import "fmt"
import "net/http"
import "time"

const downloadUrl = "http://cdn.ntwrk.waits.io/128m"

func trackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}

func test() {
	fmt.Println("Testing download speed...")

	resp, err := http.Get(downloadUrl)
	check(err)
	defer resp.Body.Close()

	buf := make([]byte, 32768)
	total := 0
	start := time.Now()
	var bandwidth, elapsed float64
	for elapsed < 8 {
		n, err := resp.Body.Read(buf)
		check(err)

		total += n
		elapsed = time.Since(start).Seconds()
		progress := int8(elapsed / 8 * 100)
		bandwidth = float64(total) / elapsed / 131072
		fmt.Printf("\r%3d%% @ %5.2f Mbps", progress, bandwidth)
	}

	fmt.Printf("\n\nYour download speed is %.2f Mbps.\n", bandwidth)
}
