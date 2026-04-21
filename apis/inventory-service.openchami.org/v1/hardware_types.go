// Copyright © 2026 OpenCHAMI a Series of LF Projects, LLC
//
// SPDX-License-Identifier: MIT

package v1

import (
	"context"
	"github.com/openchami/fabrica/pkg/fabrica"
)

// Hardware represents a hardware resource
type Hardware struct {
	APIVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   fabrica.Metadata `json:"metadata"`
	Spec       HardwareSpec   `json:"spec" validate:"required"`
	Status     HardwareStatus `json:"status,omitempty"`
}

// HardwareSpec defines the desired state of Hardware
type HardwareSpec struct {
	Description string `json:"description,omitempty" validate:"max=200"`
	// Add your spec fields here
}

// HardwareStatus defines the observed state of Hardware
type HardwareStatus struct {
	Phase      string `json:"phase,omitempty"`
	Message    string `json:"message,omitempty"`
	Ready      bool   `json:"ready"`
		// Add your status fields here
}

// Validate implements custom validation logic for Hardware
func (r *Hardware) Validate(ctx context.Context) error {
	// Add custom validation logic here
	// Example:
	// if r.Spec.Description == "forbidden" {
	//     return errors.New("description 'forbidden' is not allowed")
	// }

	return nil
}
// GetKind returns the kind of the resource
func (r *Hardware) GetKind() string {
	return "Hardware"
}

// GetName returns the name of the resource
func (r *Hardware) GetName() string {
	return r.Metadata.Name
}

// GetUID returns the UID of the resource
func (r *Hardware) GetUID() string {
	return r.Metadata.UID
}

// IsHub marks this as the hub/storage version
func (r *Hardware) IsHub() {}
