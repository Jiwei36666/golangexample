package main
import (
    "fmt"

    log "github.com/Sirupsen/logrus"
    "github.com/docker/docker/pkg/reexec"
    "github.com/docker/libnetwork"
    "github.com/docker/libnetwork/config"
    "github.com/docker/libnetwork/netlabel"
    "github.com/docker/libnetwork/options"
)

func main() {
    if reexec.Init() {
        return
    }

    // Select and configure the network driver
    networkType := "bridge"

    // Create a new controller instance
    driverOptions := options.Generic{}
    genericOption := make(map[string]interface{})
    genericOption[netlabel.GenericData] = driverOptions
    controller, err := libnetwork.New(config.OptionDriverConfig(networkType, genericOption))
    if err != nil {
        log.Fatalf("libnetwork.New: %s", err)
    }

	// Get network    
	network, err := controller.NetworkByName("network1")
    if err != nil {
        log.Fatalf("controller.NewNetwork: %s", err)
    }
	
	endpoints := network.Endpoints()
	for _, p := range endpoints {
		fmt.Printf("%v", p)
	}
}
