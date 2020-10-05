package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/billing/mgmt/2020-05-01/billing"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	subID := os.Args[1]
	invoiceName := os.Args[2]

	getStatements(subID, invoiceName)
}

func azureAuth() autorest.Authorizer {
	auth, err := auth.NewAuthorizerFromCLI()

	if err != nil {
		log.Println("There was an error authenticating with the current Azure CLI profile")
	}

	return auth
}

func getStatements(subID string, invoiceName string) {
	invoice := billing.NewInvoicesClient(subID, subID)

	// IN PROGRESS
	invoice.DownloadBillingSubscriptionInvoice(context.TODO(), invoiceName, "test")
}
