package inmem

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Registry struct {
	sync.RWMutex
	addrs map[string]map[string]*serviceInstance
}

type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

func NewRegistry() *Registry {
	return &Registry{addrs: map[string]map[string]*serviceInstance{}}
}

func (registry *Registry) Register(ctx context.Context, instanceID, serviceName, hostPort string) error {
	registry.Lock()
	defer registry.Unlock()

	if _, ok := registry.addrs[serviceName]; !ok {
		registry.addrs[serviceName] = map[string]*serviceInstance{}
	}

	return nil
}

func (registry *Registry) Deregister(ctx context.Context, instanceID, serviceName string) error {
	registry.Lock()
	defer registry.Unlock()

	if _, ok := registry.addrs[serviceName]; !ok {
		return nil
	}

	delete(registry.addrs[serviceName], instanceID)

	return nil
}

func (registry *Registry) Discover(ctx context.Context, serviceName string) ([]string, error) {
	registry.RLock()
	defer registry.RUnlock()

	if len(registry.addrs[serviceName]) == 0 {
		return nil, errors.New("no service address found")
	}

	var res []string
	for _, i := range registry.addrs[serviceName] {
		res = append(res, i.hostPort)
	}

	return res, nil
}

func (registry *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	registry.RLock()
	defer registry.RUnlock()

	if len(registry.addrs[serviceName]) == 0 {
		return nil, errors.New("no service address found")
	}

	var res []string
	for _, i := range registry.addrs[serviceName] {
		if i.lastActive.Before(time.Now().Add(-5 * time.Second)) {
			continue
		}
		res = append(res, i.hostPort)
	}

	return res, nil
}
