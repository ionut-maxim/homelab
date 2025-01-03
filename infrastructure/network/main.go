package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := LocalNetwork(ctx)
		if err != nil {
			return err
		}

		nodes := NewNodes()
		if err = nodes.Run(ctx); err != nil {
			return err
		}

		return nil
	})
}
