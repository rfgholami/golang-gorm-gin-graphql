package service

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

var client, err = api.NewClient(api.DefaultConfig())

func Register() {
	// Create a new Consul client

	if err != nil {
		log.Fatal(err)
	}

	// Register a service with Consul
	registration := new(api.AgentServiceRegistration)
	registration.ID = "gi_id"
	registration.Name = "go"
	registration.Address = "localhost"
	registration.Port = 9898
	registration.Check = &api.AgentServiceCheck{
		HTTP:     fmt.Sprintf("http://172.30.80.28:%d/health-check", 9898),
		Interval: "10s", // Check every 10 seconds
		Timeout:  "5s",
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal(err)
	}

}

func Find(name string) (string, int) {

	services, _, err := client.Health().Service(name, "", true, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, service := range services {
		return service.Service.Address, service.Service.Port
	}
	return "", 0
}
