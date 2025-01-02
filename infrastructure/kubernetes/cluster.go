package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	talosclient "github.com/pulumiverse/pulumi-talos/sdk/go/talos/client"
	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/machine"
)

const ClusterName = pulumi.String("ninthfloor")

func Cluster(ctx *pulumi.Context) error {
	secrets, err := machine.NewSecrets(ctx, "secrets-v1", &machine.SecretsArgs{TalosVersion: pulumi.String("v1.9.1")})
	if err != nil {
		return err
	}

	node03IPAddress := pulumi.String("test")

	configuration := machine.GetConfigurationOutput(ctx, machine.GetConfigurationOutputArgs{
		ClusterName:     ClusterName,
		MachineType:     pulumi.String("controlplane"),
		ClusterEndpoint: pulumi.String("https://kubernetes.internal:6443"),
		MachineSecrets:  secrets.MachineSecrets,
	})

	node03apply, err := machine.NewConfigurationApply(ctx, "node03-configuration", &machine.ConfigurationApplyArgs{
		ClientConfiguration:       secrets.ClientConfiguration,
		MachineConfigurationInput: configuration.MachineConfiguration(),
		Node:                      node03IPAddress,
		OnDestroy: machine.ConfigurationApplyOnDestroyArgs{
			Graceful: pulumi.Bool(false),
			Reboot:   pulumi.Bool(true),
			Reset:    pulumi.Bool(true),
		},
		ConfigPatches: pulumi.StringArray{
			patch("node03-machine"),
		},
	})
	if err != nil {
		return err
	}

	cfg := talosclient.GetConfigurationOutput(ctx, talosclient.GetConfigurationOutputArgs{
		ClusterName: ClusterName,
		ClientConfiguration: talosclient.GetConfigurationClientConfigurationArgs{
			CaCertificate:     secrets.ClientConfiguration.CaCertificate(),
			ClientCertificate: secrets.ClientConfiguration.ClientCertificate(),
			ClientKey:         secrets.ClientConfiguration.ClientKey(),
		},
		Nodes:     pulumi.StringArray{node03IPAddress},
		Endpoints: pulumi.StringArray{node03IPAddress},
	}, nil)

	_, err = machine.NewBootstrap(ctx, "bootstrap", &machine.BootstrapArgs{
		Node:                node03IPAddress,
		ClientConfiguration: secrets.ClientConfiguration,
	}, pulumi.DependsOn([]pulumi.Resource{
		node03apply,
	}))
	if err != nil {
		return err
	}

	ctx.Export("talos-config", cfg.TalosConfig())

	return nil
}
