package main

import (
	"errors"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/iam"
)

type Nodes struct {
	input map[string]*iam.UserArgs
}

func NewNodes() *Nodes {
	return &Nodes{input: clients}
}

func (n *Nodes) Run(ctx *pulumi.Context) error {
	if n.input == nil {
		return errors.New("no nodes configured")
	}

	nodes := make(map[string]any, len(n.input))

	for name, args := range n.input {
		u, err := iam.NewUser(ctx, name, args)
		if err != nil {
			return err
		}

		nodes[name] = u.FixedIp.Elem()
	}

	ctx.Export("ip-addresses", pulumi.ToMap(nodes))

	return nil
}
