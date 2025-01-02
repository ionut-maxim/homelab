package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		networkStack, err := pulumi.NewStackReference(ctx, "ionut-maxim/network/main", nil)
		if err != nil {
			return err
		}

		_ = networkStack.GetOutput(pulumi.String("ip-addresses"))

		//node02Address := getIPAddress("node02", ipAddresses)

		return nil
	})
}

func getIPAddress(nodeName string, addresses pulumi.AnyOutput) pulumi.StringOutput {
	return addresses.ApplyT(func(val any) string {
		if m, ok := val.(map[string]any); ok {
			if v, exists := m[nodeName]; exists {
				return v.(string)
			}
		}
		return ""
	}).(pulumi.StringOutput)
}
