//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package docker

import (
	"github.com/testcontainers/testcontainers-go"
)

type EndpointName string

var (
	HTTP EndpointName = "http"
	GRPC EndpointName = "grpc"
)

type endpoint struct {
	port, uri string
}

type DockerContainer struct {
	name        string
	endpoints   map[EndpointName]endpoint
	container   testcontainers.Container
	envSettings map[string]string
}

func (d *DockerContainer) Name() string {
	return d.name
}

func (d *DockerContainer) URI() string {
	return d.GetEndpoint(HTTP)
}

func (d *DockerContainer) GetEndpoint(name EndpointName) string {
	if endpoint, ok := d.endpoints[name]; ok {
		return endpoint.uri
	}
	return ""
}
