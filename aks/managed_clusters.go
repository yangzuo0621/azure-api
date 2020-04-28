package aks

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2020-03-01/containerservice"
)

// GetRestClient return autorest.Client that has authorizer
func GetRestClient(clientId string, clientSecret string, tenantId string) autorest.Client {
	config := auth.NewClientCredentialsConfig(clientId, clientSecret, tenantId)
	authorizer, err := config.Authorizer()
	if err != nil {
		panic(err)
	}
	client := autorest.Client{
		Authorizer: authorizer,
	}
	return client
}

// CreateAKSCluster creates an aks cluster
func CreateAKSCluster(client autorest.Client, urlParameters map[string]interface{}, apiVersion string, body interface{}) {

	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	preparerDecorators := []autorest.PrepareDecorator{
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.WithMethod("PUT"),
		autorest.WithBaseURL(azure.PublicCloud.ResourceManagerEndpoint),
		autorest.WithPathParameters(
			"/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.containerservice/managedclusters/{resourceName}",
			urlParameters,
		),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithJSON(body),
	}

	preparer := autorest.CreatePreparer(preparerDecorators...)
	req, err := preparer.Prepare((&http.Request{}).WithContext(context.Background()))

	if err != nil {
		panic(err)
	}

	fmt.Println(req.URL)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
	)

	content, err := ioutil.ReadAll(resp.Body)
	bodyString := string(content)

	fmt.Println(bodyString)
}

// DeleteAKSCluster deletes an aks cluster
func DeleteAKSCluster(client autorest.Client, urlParameters map[string]interface{}, apiVersion string) {

	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	preparerDecorators := []autorest.PrepareDecorator{
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.WithMethod("DELETE"),
		autorest.WithBaseURL(azure.PublicCloud.ResourceManagerEndpoint),
		autorest.WithPathParameters(
			"/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.containerservice/managedclusters/{resourceName}",
			urlParameters,
		),
		autorest.WithQueryParameters(queryParameters),
	}

	preparer := autorest.CreatePreparer(preparerDecorators...)
	req, err := preparer.Prepare((&http.Request{}).WithContext(context.Background()))

	if err != nil {
		panic(err)
	}

	fmt.Println(req.URL)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
	)

	fmt.Println(resp.Status)
}

// ListClusterUserCredentials gets aks cluster's kubeconfig
func ListClusterUserCredentials(client autorest.Client, urlParameters map[string]interface{}, apiVersion string) containerservice.CredentialResults {
	queryParameters := map[string]interface{}{
		"api-version": apiVersion,
	}
	preparerDecorators := []autorest.PrepareDecorator{
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.WithMethod("POST"),
		autorest.WithBaseURL(azure.PublicCloud.ResourceManagerEndpoint),
		autorest.WithPathParameters(
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/listClusterUserCredential",
			urlParameters,
		),
		autorest.WithQueryParameters(queryParameters),
	}

	preparer := autorest.CreatePreparer(preparerDecorators...)
	req, err := preparer.Prepare((&http.Request{}).WithContext(context.Background()))

	if err != nil {
		panic(err)
	}

	fmt.Println(req.URL)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
	)

	content, err := ioutil.ReadAll(resp.Body)

	var kubeconfigs containerservice.CredentialResults
	json.Unmarshal(content, &kubeconfigs)

	return kubeconfigs
}
