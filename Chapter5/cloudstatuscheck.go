package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	azurePtr := flag.Bool("azure", false, "Azure Status Check")

	flag.Parse()

	if *azurePtr {
		azureStatus()
	}
}

func azureStatus() {
	response, err := http.Get("https://status.azure.com")
	defer response.Body.Close()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	output := string(body)

	if !strings.Contains(output, "data-label=\"Good\"") {
		log.Println("A service is down in Azure")
		log.Println("Please visit https://status.azure.com for more info")
	} else {
		log.Println("All Azure Services Are Operational")
	}
}
