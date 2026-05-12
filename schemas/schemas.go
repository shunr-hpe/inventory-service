package schemas

import _ "embed"

// ComponentsSchema holds the JSON Schema for Component.
//
//go:embed components_schema.json
var ComponentsSchema []byte

// EthernetInterfaceSchema holds the JSON Schema for EthernetInterface.
//
//go:embed ethernet_interface_schema.json
var EthernetInterfaceSchema []byte

// GroupSchema holds the JSON Schema for Group.
//
//go:embed group_schema.json
var GroupSchema []byte

// HardwareSchema holds the JSON Schema for Hardware.
//
//go:embed hardware_schema.json
var HardwareSchema []byte

// RedfishEndpointSchema holds the JSON Schema for RedfishEndpoint.
//
//go:embed redfish_endpoint_schema.json
var RedfishEndpointSchema []byte

// ServiceEndpointSchema holds the JSON Schema for ServiceEndpoint.
//
//go:embed service_endpoint_schema.json
var ServiceEndpointSchema []byte
