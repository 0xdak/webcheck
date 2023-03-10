package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var client = http.Client{
	Timeout: 5 * time.Second,
}
var counter int = 0

func check_host(host string, wg *sync.WaitGroup) {
	defer wg.Done()
	prefixs := []string{"http://", "https://"}
	var alive_hosts []string
	for _, prefix := range prefixs {
		resp, err := client.Get(prefix + host)
		if err != nil {
			continue
		} else {
			alive_hosts = append(alive_hosts, strconv.Itoa(resp.StatusCode)+" : "+prefix+host)
		}
	}
	if len(alive_hosts) != 0 {
		fmt.Print(strings.Join(alive_hosts, "\n") + "\n")
		counter += 1

	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Proper usage: ./webcheck domains.txt")
		return
	}
	filename := os.Args[1]
	domains, err := os.ReadFile(filename)
	if err != nil {
		panic("")
	}
	domains_list := strings.Split(string(domains), "\n")
	// fmt.Print(strings.Join(dat_sep, "\n") + "\n")

	start := time.Now()
	var wg sync.WaitGroup
	for _, host := range domains_list {
		wg.Add(1)
		go check_host(host, &wg)
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Println("It took " + elapsed.String() + " time to check all domains!")

	fmt.Print(counter)
	fmt.Print("/")
	fmt.Print(len(domains_list) - 1)
	fmt.Print(" domains are alive")

}
