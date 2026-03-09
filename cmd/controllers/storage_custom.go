// Copyright © 2026 OpenCHAMI a Series of LF Projects, LLC
//
// SPDX-License-Identifier: MIT

package controllers

import (
	"context"

	v1 "github.com/OpenCHAMI/smd2/apis/smd2.openchami.org/v1"
	"github.com/OpenCHAMI/smd2/internal/storage"
)

// EntStorageExtras implements the plugins.StorageExtras interface by delegating
// to the package-level functions in internal/storage.
type EntStorageExtras struct{}

func (s *EntStorageExtras) LoadComponentByID(ctx context.Context, id string) (*v1.Component, error) {
	return storage.LoadComponentByID(ctx, id)
}

func (s *EntStorageExtras) LoadComponentEndpointByID(ctx context.Context, id string) (*v1.ComponentEndpoint, error) {
	return storage.LoadComponentEndpointByID(ctx, id)
}

func (s *EntStorageExtras) LoadRedfishEndpointByID(ctx context.Context, id string) (*v1.RedfishEndpoint, error) {
	return storage.LoadRedfishEndpointByID(ctx, id)
}

func (s *EntStorageExtras) LoadEthernetInterfaceByID(ctx context.Context, id string) (*v1.EthernetInterface, error) {
	return storage.LoadEthernetInterfaceByID(ctx, id)
}

func (s *EntStorageExtras) LoadServiceEndpointByID(ctx context.Context, id string) (*v1.ServiceEndpoint, error) {
	return storage.LoadServiceEndpointByID(ctx, id)
}

func (s *EntStorageExtras) LoadGroupByLabel(ctx context.Context, label string) (*v1.Group, error) {
	return storage.LoadGroupByLabel(ctx, label)
}
