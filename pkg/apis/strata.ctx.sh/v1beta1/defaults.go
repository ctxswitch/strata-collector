package v1beta1

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// DefraultCollectorEnabled is the default value for enabling the collector service.
	DefaultCollectorEnabled bool = true
	// DefaultCollectorWorkers is the default number of workers for the collector service.
	DefaultCollectorWorkers int64 = 1

	// DefaultDiscoveryPrefix is the default prefix for all resources.
	DefaultDiscoveryPrefix string = "prometheus.io"
	// DefaultDiscoveryIntervalSeconds is the default interval in seconds that the discovery
	DefaultDiscoveryIntervalSeconds int64 = 10
	// DefaultDiscoveryIncludeMetadata is the default value for including metadata.
	DefaultDiscoveryIncludeMetadata bool = false
	// DefaultDiscoveryEnabled is the default value for enabling the discovery service.
	DefaultDiscoveryEnabled bool = true
	// DefaultDiscoveryResourcePods is the default value for including pods in discovery.
	DefaultDiscoveryResourcePods bool = true
	// DefaultDiscoveryResourceServices is the default value for including services in discovery.
	DefaultDiscoveryResourceServices bool = true
	// DefaultDiscoveryResourceEndpoints is the default value for including endpoints in discovery.
	DefaultDiscoveryResourceEndpoints bool = true
)

// Defaulted sets the resource defaults.
func Defaulted(obj client.Object) {
	switch obj := obj.(type) {
	case *Collector:
		defaultedCollector(obj)
	case *Discovery:
		defaultedDiscovery(obj)
	}
}

func defaultedCollector(obj *Collector) {
	if obj.Spec.Enabled == nil {
		enabled := DefaultCollectorEnabled
		obj.Spec.Enabled = &enabled
	}

	if obj.Spec.Workers == nil {
		workers := DefaultCollectorWorkers
		obj.Spec.Workers = &workers
	}

	if obj.Spec.Output == nil {
		name := "default"
		output := CollectorOutput{
			Name:   &name,
			Stdout: &Stdout{},
		}
		obj.Spec.Output = &output
	}
}

func defaultedDiscovery(obj *Discovery) {
	if obj.Spec.Enabled == nil {
		enabled := DefaultDiscoveryEnabled
		obj.Spec.Enabled = &enabled
	}

	if obj.Spec.IntervalSeconds == nil {
		interval := DefaultDiscoveryIntervalSeconds
		obj.Spec.IntervalSeconds = &interval
	}

	if obj.Spec.Prefix == nil {
		prefix := DefaultDiscoveryPrefix
		obj.Spec.Prefix = &prefix
	}

	if obj.Spec.IncludeMetadata == nil {
		includeMetadata := DefaultDiscoveryIncludeMetadata
		obj.Spec.IncludeMetadata = &includeMetadata
	}

	obj.Spec.Resources = defaultedDiscoveryResources(obj.Spec.Resources)
	fmt.Printf("defaultedDiscoveryResources: %v", *obj.Spec.Resources)
}

func defaultedDiscoveryResources(obj *DiscoveryResources) *DiscoveryResources {
	if obj == nil {
		obj = &DiscoveryResources{}
	}

	if obj.Pods == nil {
		pods := DefaultDiscoveryResourcePods
		obj.Pods = &pods
	}

	if obj.Services == nil {
		services := DefaultDiscoveryResourceServices
		obj.Services = &services
	}

	if obj.Endpoints == nil {
		endpoints := DefaultDiscoveryResourceEndpoints
		obj.Endpoints = &endpoints
	}

	return obj
}
