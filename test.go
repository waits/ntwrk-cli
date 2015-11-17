package main

import "bytes"
import "errors"
import "fmt"
import "io"
import "net/http"
import "strings"
import "time"

const downloadUrl = "http://cdn.ntwrk.waits.io/128m"
const uploadUrl = "https://ntwrk.waits.io/upload"

type ProgressReader struct {
	io.Reader
	Reporter func(r int64) bool
	Final bool
}

func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	if pr.Final {
		return 0, errors.New("EOF")
	} else {
		n, err = pr.Reader.Read(p)
		pr.Final = pr.Reporter(int64(n))
		return n, err
	}
}

func runTests() {
	db, ub := download(), upload()
	fmt.Printf("\n\nYour download speed is %.2f Mbps.\n", db)
	fmt.Printf("Your upload speed is %.2f Mbps.\n", ub)
}

func download() float64 {
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

	return bandwidth
}

func upload() float64 {
	fmt.Println("\nTesting upload speed...")
	buf := bytes.NewBufferString(strings.Repeat("0123456789012345", 1048576))
	total := int64(0)
	start := time.Now()
	var bandwidth float64
	pr := &ProgressReader{buf, func(r int64) bool {
		elapsed := time.Since(start).Seconds()
		total += r
		bandwidth =  float64(total) / elapsed / 131072
		fmt.Printf("\r%3d%% @ %5.2f Mbps", int8(elapsed / 8 * 100), bandwidth)
		if elapsed >= 8 {
			return true
		} else {
			return false
		}
	}, false}

	req, err := http.NewRequest("POST", uploadUrl, pr)
	check(err)

	http.DefaultClient.Do(req)

	return bandwidth
}
