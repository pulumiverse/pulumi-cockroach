package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen"

	cockroach "github.com/pulumiverse/pulumi-cockroach/provider"
)

func main() {
	tfgen.Main("cockroach", cockroach.Provider())
}
