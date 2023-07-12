package cockroach

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/cockroachdb/terraform-provider-cockroach/shim"
	"github.com/lbrlabs/pulumi-cockroach/provider/pkg/version"
	tfpfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// all of the tls token components used below.
const (
	cockroachPkg = "cockroach"
	cockroachMod = "index"
)

// cockroachMember manufactures a type token for the TLS package and the given module and type.
func cockroachMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(cockroachPkg + ":" + mod + ":" + mem)
}

// cockroachType manufactures a type token for the TLS package and the given module and type.
func cockroachType(mod string, typ string) tokens.Type {
	return tokens.Type(cockroachMember(mod, typ))
}

// cockroachDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the TLS package and names the file by simply lower casing the data
// source's first character.
func cockroachDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return cockroachMember(mod+"/"+fn, res)
}

// cockroachResource manufactures a standard resource token given a module and resource name.
// It automatically uses the TLS package and names the file by simply lower casing the resource's
// first character.
func cockroachResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return cockroachType(mod+"/"+fn, res)
}

//go:embed cmd/pulumi-resource-cockroach/bridge-metadata.json
var bridgeMetadata []byte

// Provider returns additional overlaid schema and metadata associated with the tls package.
func Provider() tfpfbridge.ProviderInfo {
	info := tfbridge.ProviderInfo{
		Name:              "cockroach",
		DisplayName:       "CockroachDB",
		Publisher:         "lbrlabs",
		PluginDownloadURL: "github://api.github.com/lbrlabs",
		Description:       "A Pulumi package to create and managed cockroach db resources in Pulumi programs.",
		Keywords:          []string{"pulumi", "cockroach"},
		License:           "Apache-2.0",
		LogoURL:           "https://raw.githubusercontent.com/lbrlabs/pulumi-cockroach/main/assets/logo.png",
		Homepage:          "https://lbrlabs.com",
		Repository:        "https://github.com/lbrlabs/pulumi-cockroach",
		Version:           version.Version,
		GitHubOrg:         "cockroachdb",
		MetadataInfo:      tfbridge.NewProviderMetadata(bridgeMetadata),
		Resources: map[string]*tfbridge.ResourceInfo{
			"cockroach_allow_list":                      {Tok: cockroachResource(cockroachMod, "AllowList")},
			"cockroach_client_ca_cert":                  {
				Tok: cockroachResource(cockroachMod, "CaCert"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
			},
			"cockroach_cluster":                         {
				Tok: cockroachResource(cockroachMod, "Cluster"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
			},
			"cockroach_cmek":                            {
				Tok: cockroachResource(cockroachMod, "Cmek"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
			},
			"cockroach_database":                        {Tok: cockroachResource(cockroachMod, "Database")},
			"cockroach_finalize_version_upgrade":        {
				Tok: cockroachResource(cockroachMod, "FinalizeVersionUpgrade"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
			},
			"cockroach_log_export_config":               {
				Tok: cockroachResource(cockroachMod, "LogExportConfig"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
			},
			"cockroach_maintenance_window":              {
				Tok: cockroachResource(cockroachMod, "MaintenanceWindow"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
			},
			"cockroach_metric_export_cloudwatch_config": {
				Tok: cockroachResource(cockroachMod, "MetricExportCloudwatchConfig"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
			},
			"cockroach_metric_export_datadog_config":    {
				Tok: cockroachResource(cockroachMod, "MetricExportDatadogConfig"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
			},
			"cockroach_private_endpoint_connection":     {Tok: cockroachResource(cockroachMod, "PrivateEndpointConnection")},
			"cockroach_private_endpoint_services":       {Tok: cockroachResource(cockroachMod, "PrivateEndpointServices")},
			"cockroach_sql_user":                        {Tok: cockroachResource(cockroachMod, "SqlUser")},
			"cockroach_user_role_grants":                {Tok: cockroachResource(cockroachMod, "UserRoleGrants")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"cockroach_cluster":           {Tok: cockroachDataSource(cockroachMod, "getCockroachCluster")},
			"cockroach_cluster_cert":      {Tok: cockroachDataSource(cockroachMod, "getClusterCert")},
			"cockroach_connection_string": {Tok: cockroachDataSource(cockroachMod, "getConnectionString")},
			"cockroach_organization":      {Tok: cockroachDataSource(cockroachMod, "getOrganization")},
			"cockroach_person_user":       {Tok: cockroachDataSource(cockroachMod, "getPersonUser")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@lbrlabs/pulumi-cockroach",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "lbrlabs_pulumi_cockroach",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/lbrlabs/pulumi-%[1]s/sdk/", cockroachPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				cockroachPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Lbrlabs.PulumiPackage",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"cockroach": "Cockroach",
			},
		},
	}

	return tfpfbridge.ProviderInfo{
		ProviderInfo: info,
		NewProvider:  shim.NewProvider,
	}
}
