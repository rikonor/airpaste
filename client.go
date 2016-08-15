package airpaste

import "github.com/micro/mdns"

const (
	defaultServiceName = "default"
)

type Server struct {
	IPAddr string
	Port   int
}

func SearchForOpenServer(serviceName string) Server {
	var waitingServer Server

	// Create a new channel for incoming entries
	entriesCh := make(chan *mdns.ServiceEntry)

	go func() {
		for entry := range entriesCh {
			waitingServer = Server{
				IPAddr: entry.Addr.String(),
				Port:   entry.Port,
			}
		}
	}()

	mdns.Lookup(serviceName, entriesCh)
	close(entriesCh)

	return waitingServer
}
