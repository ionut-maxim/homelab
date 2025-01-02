package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi"
)

func LocalNetwork(ctx *pulumi.Context) (*unifi.Network, error) {
	return unifi.NewNetwork(ctx, "local", &unifi.NetworkArgs{
		DhcpEnabled:         pulumi.Bool(true),
		DhcpStart:           pulumi.String("10.20.30.200"),
		DhcpStop:            pulumi.String("10.20.31.254"),
		DomainName:          pulumi.String("internal"),
		Ipv6InterfaceType:   pulumi.String("pd"),
		Ipv6PdInterface:     pulumi.String("wan"),
		Ipv6PdStart:         pulumi.String("::2"),
		Ipv6PdStop:          pulumi.String("::7d1"),
		Ipv6RaEnable:        pulumi.Bool(true),
		Ipv6RaPriority:      pulumi.String("high"),
		Ipv6RaValidLifetime: pulumi.Int(0),
		MulticastDns:        pulumi.Bool(true),
		Name:                pulumi.String("local"),
		Purpose:             pulumi.String("corporate"),
		Site:                pulumi.String("default"),
		Subnet:              pulumi.String("10.20.30.0/23"),
	}, pulumi.Protect(true))
}
