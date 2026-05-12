package schemas

import _ "embed"

// ComponentsSchema holds the JSON Schema for ComponentSpec.
//
//go:embed components_schema.json
var ComponentsSchema []byte

// EthernetInterfaceSchema holds the JSON Schema for EthernetInterfaceSpec.
//
//go:embed ethernet_interface_schema.json
var EthernetInterfaceSchema []byte

// GroupSchema holds the JSON Schema for GroupSpec.
//
//go:embed group_schema.json
var GroupSchema []byte

// HardwareSchema holds the JSON Schema for HardwareSpec.
//
//go:embed hardware_schema.json
var HardwareSchema []byte

// RedfishEndpointSchema holds the JSON Schema for RedfishEndpointSpec.
//
//go:embed redfish_endpoint_schema.json
var RedfishEndpointSchema []byte

// ServiceEndpointSchema holds the JSON Schema for ServiceEndpointSpec.
//
//go:embed service_endpoint_schema.json
var ServiceEndpointSchema []byte
