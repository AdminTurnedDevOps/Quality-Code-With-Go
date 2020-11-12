package main

import (
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerinstance/mgmt/containerinstance"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {

	subscriptionID := os.Args[1]

	createContainerGroup(subscriptionID)

}

func azureAuth() autorest.Authorizer {
	auth, err := auth.NewAuthorizerFromCLI()

	if err != nil {
		log.Println("There was an error authenticating with the current Azure CLI profile")
	}

	return auth
}

func createContainerGroup(subscriptionID string) {
	containerinstance.NewContainerGroupsClient(subscriptionID)
}
