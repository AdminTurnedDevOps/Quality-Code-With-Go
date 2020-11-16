package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerinstance/mgmt/containerinstance"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {

	subscriptionID := os.Args[1]
	resourceGroupName := os.Args[2]
	name := os.Args[3]
	location := os.Args[3]

	createContainerGroup(subscriptionID, resourceGroupName, name, location)

}

func azureAuth() autorest.Authorizer {
	auth, err := auth.NewAuthorizerFromCLI()

	if err != nil {
		log.Println("There was an error authenticating with the current Azure CLI profile")
	}

	return auth
}

func createContainerGroup(subscriptionID, resourceGroupName, name, location string) containerinstance.ContainerGroupsCreateOrUpdateFuture {
	containerGroup := containerinstance.NewContainerGroupsClient(subscriptionID)

	containerGroup.Authorizer = azureAuth()
	create, err := containerGroup.CreateOrUpdate(context.Background(), resourceGroupName, name, containerinstance.ContainerGroup{
		Name:     &name,
		Location: &location, ContainerGroupProperties: &containerinstance.ContainerGroupProperties{
			IPAddress: &containerinstance.IPAddress{
				Type: containerinstance.Public,
				Ports: &[]containerinstance.Port{
					{
						Port:     to.Int32Ptr(8080),
						Protocol: containerinstance.TCP,
					},
				},
			},
			OsType: containerinstance.Linux,
			Containers: &[]containerinstance.Container{
				{
					Name: to.StringPtr("gowebapi"),
					ContainerProperties: &containerinstance.ContainerProperties{
						Ports: &[]containerinstance.ContainerPort{
							{
								Port: to.Int32Ptr(8080),
							},
						},
						Image: to.StringPtr("adminturneddevops/golangwebapi:latest"),
						Resources: &containerinstance.ResourceRequirements{
							Limits: &containerinstance.ResourceLimits{
								MemoryInGB: to.Float64Ptr(1),
								CPU:        to.Float64Ptr(1),
							},
							Requests: &containerinstance.ResourceRequests{
								MemoryInGB: to.Float64Ptr(1),
								CPU:        to.Float64Ptr(1),
							},
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Println(err)
	}

	return create
}
