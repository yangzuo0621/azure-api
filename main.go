package main

import (
	"fmt"

	"github.com/yangzuo0621/azure-api/aks"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clusterName             = "zuya-aks-test-001"
	subscriptionID          = "99a2af06-30ed-4c4f-9f74-2fc72b032fec"
	tenantID                = "9c0adc1d-423f-4d79-a7c0-352bc5beadce"
	clientID                = "0615c5a5-5a0e-4a94-a0a6-490f7847da55"
	clientSecret            = ""
	groupName               = "testkeyvaultgroup"
	myDiskEncryptionSetName = "/subscriptions/99a2af06-30ed-4c4f-9f74-2fc72b032fec/resourceGroups/testkeyvaultgroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSetName1"
	location                = "west us 2"
	kubernetesVersion       = "1.17.3"
	apiVersion              = "2020-03-01"
)

func main() {

	// payload := containerservice.ManagedCluster{
	// 	Location: &location,
	// 	ManagedClusterProperties: &containerservice.ManagedClusterProperties{
	// 		KubernetesVersion:       &kubernetesVersion,
	// 		EnableRBAC:              to.BoolPtr(true),
	// 		EnablePodSecurityPolicy: to.BoolPtr(false),
	// 		// DiskEncryptionSetID:     &myDiskEncryptionSetName,
	// 		DNSPrefix:     to.StringPtr("zuyatest30503264-aaaa"),
	// 		AddonProfiles: nil,
	// 		ServicePrincipalProfile: &containerservice.ManagedClusterServicePrincipalProfile{
	// 			ClientID: to.StringPtr(clientID),
	// 			Secret:   to.StringPtr(clientSecret),
	// 		},
	// 		AgentPoolProfiles: &[]containerservice.ManagedClusterAgentPoolProfile{
	// 			{
	// 				Name:              to.StringPtr("agentpool0"),
	// 				Count:             to.Int32Ptr(1),
	// 				VMSize:            "Standard_DS2_v2",
	// 				MaxCount:          to.Int32Ptr(3),
	// 				MinCount:          to.Int32Ptr(1),
	// 				EnableAutoScaling: to.BoolPtr(true),
	// 				Type:              "VirtualMachineScaleSets",
	// 				Mode:              containerservice.System,
	// 			},
	// 		},
	// 	},
	// }

	urlParameters := map[string]interface{}{
		"subscriptionId":    subscriptionID,
		"resourceGroupName": groupName,
		"resourceName":      clusterName,
	}

	client := aks.GetRestClient(clientID, clientSecret, tenantID)
	// aks.CreateAKSCluster(client, urlParameters, apiVersion, &payload)
	// aks.DeleteAKSCluster(client, urlParameters, apiVersion)
	result := aks.ListClusterUserCredentials(client, urlParameters, apiVersion)

	if (len(*result.Kubeconfigs)) <= 0 {
		panic("")
	}

	kubeconfig := (*result.Kubeconfigs)[0]
	clientcmdConfig, err := clientcmd.Load(*kubeconfig.Value)

	if err != nil {
		panic(err)
	}

	clientcmd.NewNonInteractiveClientConfig(*clientcmdConfig, "", &clientcmd.ConfigOverrides{}, nil)
	directClientcmdConfig := clientcmd.NewNonInteractiveClientConfig(*clientcmdConfig, "", &clientcmd.ConfigOverrides{}, nil)
	clientRestConfig, err := directClientcmdConfig.ClientConfig()

	clientset, err := kubernetes.NewForConfig(clientRestConfig)

	if err != nil {
		panic(err)
	}

	v, _ := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	nodes := v.Items
	fmt.Println(len(nodes))
}
