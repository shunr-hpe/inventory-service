package v1

import (
	"context"
	"encoding/json"

	"github.com/openchami/fabrica/pkg/fabrica"
)

type EthernetInterface struct {
	APIVersion string                  `json:"apiVersion"`
	Kind       string                  `json:"kind"`
	Metadata   fabrica.Metadata        `json:"metadata"`
	Spec       EthernetInterfaceSpec   `json:"spec" validate:"required"`
	Status     EthernetInterfaceStatus `json:"status,omitempty"`
}

type EthernetInterfaceSpec struct {
	Description string `json:"description,omitempty" validate:"max=200"`
	OContext    string `json:"@odata.context"`
	Oid         string `json:"@odata.id"`
	Otype       string `json:"@odata.type"`
	AutoNeg     *bool  `json:"AutoNeg,omitempty"`
	FQDN        string `json:"FQDN"`
	FullDuplex  *bool  `json:"FullDuplex,omitempty"`

	Hostname            string              `json:"HostName"`
	Id                  string              `json:"Id"`
	IPv4Addresses       []IPv4Address       `json:"IPv4Addresses"`
	IPv6Addresses       []IPv6Address       `json:"IPv6Addresses"`
	IPv6StaticAddresses []IPv6StaticAddress `json:"IPv6StaticAddresses"`
	IPv6DefaultGateway  string              `json:"IPv6DefaultGateway"`
	InterfaceEnabled    *bool               `json:"InterfaceEnabled,omitempty"`

	MACAddress             string      `json:"MACAddress"`
	PermanentMACAddress    string      `json:"PermanentMACAddress"`
	MTUSize                json.Number `json:"MTUSize"`
	MaxIPv6StaticAddresses json.Number `json:"MaxIPv6StaticAddresses"`
	Name                   string      `json:"Name"`
	NameServers            []string    `json:"NameServers"`
	SpeedMbps              json.Number `json:"SpeedMbps"`
	Status                 StatusRF    `json:"Status"`
	UefiDevicePath         string      `json:"UefiDevicePath"`
	VLAN                   VLAN        `json:"VLAN"`
}

type EthernetInterfaceStatus struct {
	Phase   string `json:"phase,omitempty"`
	Message string `json:"message,omitempty"`
	Ready   bool   `json:"ready"`
}

func (r *EthernetInterface) Validate(ctx context.Context) error {

	return nil
}

func (r *EthernetInterface) GetKind() string {
	return "EthernetInterface"
}

func (r *EthernetInterface) GetName() string {
	return r.Metadata.Name
}

func (r *EthernetInterface) GetUID() string {
	return r.Metadata.UID
}

func (r *EthernetInterface) IsHub() {}

type VLAN struct {
	VLANEnable *bool       `json:"VLANEnable,omitempty"`
	VLANid     json.Number `json:"VLANid"`
}
type IPv4Address struct {
	Address       string `json:"Address"`
	AddressOrigin string `json:"AddressOrigin"`
	Gateway       string `json:"Gateway"`
	SubnetMask    string `json:"SubnetMask"`
}
type IPv6Address struct {
	Address       string      `json:"Address"`
	AddressOrigin string      `json:"AddressOrigin"`
	AddressState  string      `json:"AddressState"`
	PrefixLength  json.Number `json:"PrefixLength"`
}
type IPv6StaticAddress struct {
	Address      string      `json:"Address"`
	PrefixLength json.Number `json:"PrefixLength"`
}
