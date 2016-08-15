package airpaste

import (
	"os"

	"github.com/micro/mdns"
)

func PublishService(serviceName string, servicePort int) error {
	host, err := os.Hostname()
	if err != nil {
		return err
	}

	service, err := mdns.NewMDNSService(host, serviceName, "", "", servicePort, nil, []string{})
	if err != nil {
		return err
	}

	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return err
	}

	defer server.Shutdown()
	select {}
}
