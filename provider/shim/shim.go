package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	cockroach "github.com/cockroachdb/terraform-provider-cockroach/internal/provider"
	"github.com/lbrlabs/pulumi-cockroach/provider/pkg/version"
)

func NewProvider() tfpf.Provider {
	prov := cockroach.New(version.Version)

	return prov()
}