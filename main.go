package main

import (
	"fmt"
	"github.com/aniketgawade/contrail-go-api"
	"github.com/aniketgawade/contrail-go-api/config"
	"github.com/aniketgawade/contrail-go-api/types"
	"github.com/pborman/uuid"
	"os"
)

var oc_client *contrail.Client


func connectToContrailApiServer() {
	if oc_client == nil {
	        fmt.Printf("Connecting to API Server ")
		oc_client = contrail.NewClient("127.0.0.1", 8082)
	}
}

func main() {
	connectToContrailApiServer()
	var parent_id string
	var err error
	parent_id, err = config.GetProjectId(
		oc_client, "default-project", "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	        fmt.Printf("Network : Cannot Find Tenant %s", "default-project")
		os.Exit(1)
	}
	uuid := uuid.New()
	net_uuid, err := config.CreateNetworkWithSubnet(oc_client, parent_id, "quick-test-net", uuid, "10.0.0.0/24", "10.0.0.254")
	fmt.Printf("Network created with name :%s", net_uuid)
	netobj, err := types.VirtualNetworkByUuid(oc_client, net_uuid)
	config.AddSubnet(oc_client, netobj, "4.2.2.0/24")	
	config.AddSubnet(oc_client, netobj, "5.2.2.0/24")	
	config.AddSubnet(oc_client, netobj, "6.2.2.0/24")	
	config.RemoveSubnet(oc_client, netobj, "5.2.2.0/24")
	config.RemoveSubnet(oc_client, netobj, "10.0.0.0/24")
}
