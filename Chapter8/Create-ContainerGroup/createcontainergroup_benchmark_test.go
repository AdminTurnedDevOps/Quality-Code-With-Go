package main

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerinstance/mgmt/containerinstance"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var subscriptionID string = "220284d2-6a19-4781-87f8-5c564ec4fec9"
var location string = "eastus"
var resourceGroupName string = "dev2"
var name string = "mjltestcontainergroup"

func BenchmarkContainerGroup(b *testing.B) {
	auth, err := auth.NewAuthorizerFromCLI()

	containerGroup := containerinstance.NewContainerGroupsClient(subscriptionID)

	containerGroup.Authorizer = auth
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
		b.Error(err)
	} else {
		b.Log(create)
	}
}
