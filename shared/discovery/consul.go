package discovery

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type ServiceDiscovery struct {
	client *api.Client
}

func NewServiceDiscovery(address string) (*ServiceDiscovery, error) {
	config := api.DefaultConfig()
	config.Address = address
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &ServiceDiscovery{client: client}, nil
}

func (sd *ServiceDiscovery) RegisterService(name, host string, port int) error {
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", name, host, port),
		Name:    name,
		Address: host,
		Port:    port,
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/health", host, port),
			Interval: "10s",
			Timeout:  "5s",
		},
	}
	return sd.client.Agent().ServiceRegister(reg)
}

func (sd *ServiceDiscovery) DeregisterService(name, host string, port int) error {
	return sd.client.Agent().ServiceDeregister(fmt.Sprintf("%s-%s-%d", name, host, port))
}

func (sd *ServiceDiscovery) DiscoverService(name string) (*api.AgentService, error) {
	services, err := sd.client.Agent().Services()
	if err != nil {
		return nil, err
	}
	for _, service := range services {
		if service.Service == name {
			return service, nil
		}
	}
	return nil, fmt.Errorf("service %s not found", name)
}
