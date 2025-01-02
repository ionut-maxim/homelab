package main

import (
	"embed"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

//go:embed patches
var patches embed.FS

func patch(name string) pulumi.String {
	// Ignore errors
	b, _ := patches.ReadFile(name + ".yaml")
	if b != nil {
		return pulumi.String(b)
	}
	return ""
}
