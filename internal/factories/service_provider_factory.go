package factories

import (
	"fmt"

	"github.com/mahsamnz/check24/internal/interfaces"
)

type ServiceProviderFactory struct {
	ServiceProviders map[string]interfaces.ServiceProvider
}

func NewServiceProviderFactory() *ServiceProviderFactory {
	return &ServiceProviderFactory{
		ServiceProviders: make(map[string]interfaces.ServiceProvider),
	}
}

func (factory *ServiceProviderFactory) RegisterService(provider interfaces.ServiceProvider) {
	factory.ServiceProviders[provider.GetIdentifier()] = provider
}

func (factory *ServiceProviderFactory) HasProvider(id string) bool {
	_, exists := factory.ServiceProviders[id]
	return exists
}

func (factory *ServiceProviderFactory) GetProvider(id string) (interfaces.ServiceProvider, error) {
	serviceProvider, exists := factory.ServiceProviders[id]
	if !exists {
		return nil, fmt.Errorf("service provider with id %s is not defined", id)
	}
	return serviceProvider, nil
}
