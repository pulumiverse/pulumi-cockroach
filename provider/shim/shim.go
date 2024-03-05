package shim

import (
	cockroach "github.com/cockroachdb/terraform-provider-cockroach/internal/provider"
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
)

func NewProvider(version string) tfpf.Provider {
	prov := cockroach.New(version)

	return prov()
}
