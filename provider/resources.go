package cockroach

import (
	"fmt"
	"path/filepath"
	"unicode"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"

	"github.com/cockroachdb/terraform-provider-cockroach/shim"

	tfpfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumiverse/pulumi-cockroach/provider/pkg/version"
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

func computeIDField(field resource.PropertyKey) tfbridge.ComputeID {
	return tfbridge.DelegateIDField(field, "cockroach", "https://github.com/pulumiverse/pulumi-cockroach")
}

// Provider returns additional overlaid schema and metadata associated with the tls package.
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                 tfpfbridge.ShimProvider(shim.NewProvider(version.Version)),
		Name:              "cockroach",
		DisplayName:       "CockroachDB",
		Publisher:         "pulumiverse",
		PluginDownloadURL: "github://api.github.com/pulumiverse",
		Description:       "A Pulumi package to create and managed Cockroach DB resources in Pulumi programs.",
		Keywords:          []string{"pulumi", "cockroach", "pulumiverse"},
		License:           "Apache-2.0",
		LogoURL:           "https://raw.githubusercontent.com/pulumiverse/pulumi-cockroach/main/assets/logo.png",
		Homepage:          "https://www.cockroachlabs.com/",
		Repository:        "https://github.com/pulumiverse/pulumi-cockroach",
		Version:           version.Version,
		GitHubOrg:         "cockroachdb",
		MetadataInfo:      tfbridge.NewProviderMetadata(bridgeMetadata),
		Resources: map[string]*tfbridge.ResourceInfo{
			"cockroach_allow_list": {Tok: cockroachResource(cockroachMod, "AllowList")},
			"cockroach_api_key":    {Tok: cockroachResource(cockroachMod, "ApiKey")},
			"cockroach_client_ca_cert": {
				Tok: cockroachResource(cockroachMod, "CaCert"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
				ComputeID: computeIDField("id"),
			},
			"cockroach_cluster": {Tok: cockroachResource(cockroachMod, "Cluster")},
			"cockroach_cmek": {
				Tok: cockroachResource(cockroachMod, "Cmek"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
				ComputeID: computeIDField("id"),
			},
			"cockroach_database": {Tok: cockroachResource(cockroachMod, "Database")},
			"cockroach_folder":   {Tok: cockroachResource(cockroachMod, "Folder")},
			"cockroach_finalize_version_upgrade": {
				Tok: cockroachResource(cockroachMod, "FinalizeVersionUpgrade"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
				ComputeID: computeIDField("id"),
			},
			"cockroach_log_export_config": {
				Tok: cockroachResource(cockroachMod, "LogExportConfig"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
				ComputeID: computeIDField("id"),
			},
			"cockroach_maintenance_window": {
				Tok: cockroachResource(cockroachMod, "MaintenanceWindow"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
				ComputeID: computeIDField("id"),
			},
			"cockroach_metric_export_cloudwatch_config": {
				Tok: cockroachResource(cockroachMod, "MetricExportCloudwatchConfig"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
				ComputeID: computeIDField("id"),
			},
			"cockroach_metric_export_datadog_config": {
				Tok: cockroachResource(cockroachMod, "MetricExportDatadogConfig"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
				ComputeID: computeIDField("id"),
			},
			"cockroach_metric_export_prometheus_config": {
				Tok: cockroachResource(cockroachMod, "MetricExportPrometheusConfig"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
				ComputeID: computeIDField("id"),
			},

			"cockroach_private_endpoint_connection":    {Tok: cockroachResource(cockroachMod, "PrivateEndpointConnection")},
			"cockroach_private_endpoint_services":      {Tok: cockroachResource(cockroachMod, "PrivateEndpointServices")},
			"cockroach_private_endpoint_trusted_owner": {Tok: cockroachResource(cockroachMod, "PrivateEndpointTrustedOwner")},
			"cockroach_service_account":                {Tok: cockroachResource(cockroachMod, "ServiceAccount")},
			"cockroach_sql_user":                       {Tok: cockroachResource(cockroachMod, "SqlUser")},
			"cockroach_user_role_grant": {
				Tok:       cockroachResource(cockroachMod, "UserRoleGrant"),
				ComputeID: computeIDField("id"),
			},
			"cockroach_user_role_grants": {Tok: cockroachResource(cockroachMod, "UserRoleGrants")},
			"cockroach_version_deferral": {
				Tok: cockroachResource(cockroachMod, "VersionDeferral"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"id": {
						Name: "clusterId",
						Type: "string",
					},
				},
				ComputeID: computeIDField("id"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"cockroach_cluster":           {Tok: cockroachDataSource(cockroachMod, "getCockroachCluster")},
			"cockroach_cluster_cert":      {Tok: cockroachDataSource(cockroachMod, "getClusterCert")},
			"cockroach_connection_string": {Tok: cockroachDataSource(cockroachMod, "getConnectionString")},
			"cockroach_organization":      {Tok: cockroachDataSource(cockroachMod, "getOrganization")},
			"cockroach_person_user":       {Tok: cockroachDataSource(cockroachMod, "getPersonUser")},
			"cockroach_folder":            {Tok: cockroachDataSource(cockroachMod, "getFolder")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/cockroach",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
			},
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumiverse_cockroach",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PyProject:            struct{ Enabled bool }{true},
			RespectSchemaVersion: true,
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", cockroachPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				cockroachPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"cockroach": "Cockroach",
			},
			RespectSchemaVersion: true,
		},
	}

	prov.MustComputeTokens(tks.SingleModule("cockroach_", cockroachMod, tks.MakeStandard(cockroachPkg)))
	prov.SetAutonaming(255, "-")
	prov.MustApplyAutoAliases()

	return prov
}
