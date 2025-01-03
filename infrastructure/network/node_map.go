package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/iam"
)

var clients = map[string]*iam.UserArgs{
	"node02": {
		DevIdOverride:  pulumi.Int(2847),
		FixedIp:        pulumi.String("10.20.30.12"),
		LocalDnsRecord: pulumi.String("node02.internal"),
		Mac:            pulumi.String("10:62:e5:17:6a:f0"),
		Name:           pulumi.String("node02"),
		Site:           pulumi.String("default"),
	},
	"node03": {
		FixedIp:        pulumi.String("10.20.30.13"),
		LocalDnsRecord: pulumi.String("node03.internal"),
		Mac:            pulumi.String("f4:39:09:45:fa:c5"),
		Name:           pulumi.String("node03"),
		Site:           pulumi.String("default"),
	},
	"kubernetes-vip": {
		FixedIp:        pulumi.String("10.20.30.103"),
		LocalDnsRecord: pulumi.String("kubernetes.internal"),
		Name:           pulumi.String("kubernetes-vip"),
		Site:           pulumi.String("default"),
		Mac:            pulumi.String("61:aa:2f:e5:5d:96"),
	},
}
