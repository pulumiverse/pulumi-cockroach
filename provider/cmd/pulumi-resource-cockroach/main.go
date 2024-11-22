// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//go:generate go run ./generate.go

package main

import (
	"context"

	_ "embed"

	tfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"

	cockroach "github.com/pulumiverse/pulumi-cockroach/provider"
)

//go:embed schema-embed.json
var pulumiSchema []byte

//go:embed bridge-metadata.json
var bridgeMetadata []byte

func main() {
	meta := tfbridge.ProviderMetadata{
		PackageSchema:  pulumiSchema,
		BridgeMetadata: bridgeMetadata,
	}
	tfbridge.Main(context.Background(), "cockroach", cockroach.Provider(), meta)
}
