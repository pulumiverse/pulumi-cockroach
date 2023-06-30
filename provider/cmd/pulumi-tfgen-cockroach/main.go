package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/pf/tfgen"

	cockroach "github.com/lbrlabs/pulumi-cockroach/provider"
)

func main() {
	tfgen.Main("cockroach", cockroach.Provider())
}