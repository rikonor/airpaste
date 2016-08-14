package airpaste

import "github.com/hashicorp/mdns"

const (
	defaultServiceName = "default"
)

type Service struct {
	IPAddr string
	Port   int
}

func getWaitingServers(serviceName string) []Service {
	var waitingServers []Service

	// Create a new channel for incoming entries
	entriesCh := make(chan *mdns.ServiceEntry)

	go func() {
		for entry := range entriesCh {
			waitingServers = append(waitingServers, Service{
				IPAddr: entry.Addr.String(),
				Port:   entry.Port,
			})
		}
	}()

	mdns.Lookup(serviceName, entriesCh)
	close(entriesCh)

	return waitingServers
}

// func main() {
// 	services := getWaitingServers(defaultServiceName)
// 	for _, svc := range services {
// 		fmt.Println(svc.IPAddr, svc.Port)
// 	}
// }
