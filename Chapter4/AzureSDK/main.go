package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/billing/mgmt/2020-05-01/billing"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	subID := os.Args[1]
	invoiceName := os.Args[2]

	getInvoiceDate(subID, invoiceName)
}

func azureAuth() autorest.Authorizer {
	auth, err := auth.NewAuthorizerFromCLI()

	if err != nil {
		log.Println("There was an error authenticating with the current Azure CLI profile")
	}

	return auth
}

func getInvoiceDate(subID string, invoiceName string) {
	if subID == "" {
		log.Println("Please add in a subscription ID")
	}

	if invoiceName == "" {
		log.Println("Please add in an invoice name")
	}

	invoiceClient := billing.NewInvoicesClient(subID, subID)

	if azureAuth() == nil {
		log.Panicln("No Azure CLI auth detected!")
	}

	invoiceClient.Authorizer = azureAuth()

	getByID, err := invoiceClient.GetByID(context.Background(), invoiceName)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(*getByID.InvoiceDate)
	}
}
